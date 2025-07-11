package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type ExchangeRate struct {
	USDToBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func main() {
	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Cria a tabela se não existir
	createTable := `CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY, bid TEXT, created_at DATETIME);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		ctxAPI, cancelAPI := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancelAPI()

		exchangeRate, err := fetchExchangeRate(ctxAPI)
		if err != nil {
			log.Printf("Erro ao buscar cotação: %v", err)
			http.Error(w, "Erro ao buscar cotação", http.StatusInternalServerError)
			return
		}

		ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancelDB()

		if err := saveExchangeRate(ctxDB, db, exchangeRate.USDToBRL.Bid); err != nil {
			log.Printf("Erro ao salvar cotação no banco: %v", err)
			http.Error(w, "Erro ao salvar cotação", http.StatusInternalServerError)
			return
		}

		// Retorna apenas o campo bid
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"bid": exchangeRate.USDToBRL.Bid,
		})
	})

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fetchExchangeRate(ctx context.Context) (*ExchangeRate, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var exchangeRate ExchangeRate
	if err := json.NewDecoder(resp.Body).Decode(&exchangeRate); err != nil {
		return nil, err
	}

	return &exchangeRate, nil
}

func saveExchangeRate(ctx context.Context, db *sql.DB, bid string) error {
	stmt, err := db.PrepareContext(ctx, "INSERT INTO cotacoes(bid, created_at) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, bid, time.Now())
	return err
}
