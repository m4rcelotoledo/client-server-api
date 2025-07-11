# Makefile

.PHONY: run-server run-client build-server build-client clean

run-server:
	go run server.go

run-client:
	go run client.go

build-server:
	go build -o server server.go

build-client:
	go build -o client client.go

clean:
	rm -f server client cotacoes.db cotacao.txt
