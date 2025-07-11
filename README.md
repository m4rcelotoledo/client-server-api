# Desafio FullCycle - Pós Go Expert

## 📋 Descrição

Este projeto contém dois programas escritos em Go:

- `client.go`: um cliente HTTP que requisita a cotação do servidor e salva o resultado em um arquivo local.
- `server.go`: um servidor HTTP que consulta a cotação do dólar (USD-BRL), salva no banco SQLite e responde ao cliente.

---

## 🚀 Requisitos

- Go 1.24+ instalado
- Git
- SQLite3

---

## ⚙️ Instalação

Clone o repositório:

```bash
git clone https://github.com/m4rcelotoledo/client-server-api.git
cd client-server-api
```

Instale as dependências do projeto:

```bash
go mod tidy
```

---

## ▶️ Executando o servidor

Abra um terminal e rode o servidor:

```bash
make run-server
```

O servidor estará disponível na porta **8080**, no endpoint:

```
http://localhost:8080/cotacao
```

---

## ▶️ Executando o cliente

Em outro terminal, rode o cliente:

```bash
make run-client
```

O cliente fará uma requisição para o servidor e irá:

- Exibir o valor da cotação no terminal
- Criar um arquivo `cotacao.txt` com o seguinte conteúdo:

```
Dólar: <valor>
```

## ✔️ Comandos que você poderá usar:

| Comando             | O que faz                                  |
| ------------------- | ------------------------------------------ |
| `make run-server`   | Executa o servidor                         |
| `make run-client`   | Executa o client                           |
| `make build-server` | Compila o server (gera binário `./server`) |
| `make build-client` | Compila o client (gera binário `./client`) |
| `make clean`        | Remove binários e arquivos gerados         |


---

## 🕐 Timeouts implementados

| Componente                                | Timeout |
| ----------------------------------------- | ------- |
| Requisição do Server para a API de câmbio | 200ms   |
| Gravação da cotação no banco SQLite       | 10ms    |
| Requisição do Client para o Server        | 300ms   |

Todos os contextos disparam erros nos logs caso o tempo se esgote.

---

## 🗃️ Estrutura do banco

Ao rodar o servidor, será criado um banco SQLite (`cotacoes.db`) com a tabela:

```sql
CREATE TABLE IF NOT EXISTS cotacoes (
    id INTEGER PRIMARY KEY,
    bid TEXT,
    created_at DATETIME
);
```

---

## ✔️ Exemplo de saída no terminal do cliente

```
Cotação salva com sucesso: 5.4231
```

Exemplo de conteúdo do arquivo:

```
Dólar: 5.4231
```

---

## ✅ Testes e verificação do banco (opcional)

Verifique os registros salvos:

```bash
sqlite3 cotacoes.db
sqlite> SELECT * FROM cotacoes;
```

---

## 🔗 API pública utilizada

- [AwesomeAPI - USD-BRL](https://economia.awesomeapi.com.br/json/last/USD-BRL)

---

## 👨‍💻 Autor

Desenvolvido por **Marcelo Toledo** durante o curso Pós Go Expert - FullCycle.
