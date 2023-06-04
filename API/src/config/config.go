package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//é a string de conexão com o mysql
	StringConexaoBanco = ""
	//porta onde a API vai estar rodando
	Porta = 0
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	//jogar valor nas variáveis
	//foi criado um arquivo .env no diretorio raiz, ele contém as variáveis de ambiente da API
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT")) //Atoi pega uma string e converte para int
	if erro != nil {
		Porta = 9000
	}
	//                                usuario:senha@/bancodedados
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?chatset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}
