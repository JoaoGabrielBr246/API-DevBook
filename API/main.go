package main

//O pacote main vai chamar os pacotes

import (
	"api/API/src/config"
	"api/API/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%d", config.Porta), r)) //ligando o servidor, ouvindo a porta do .env
}
