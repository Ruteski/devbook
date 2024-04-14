package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnStr = ""
	Porta   = 0
	// chave usada pra assiunar o token jwt
	Secretkey []byte
)

// inicializa as variaveis de ambiente
func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	//converte uma string para int
	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Porta = 9000
	}

	ConnStr = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		//ConnStr = fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_ENDERECO"),
		os.Getenv("DB_NOME"),
	)

	Secretkey = []byte(os.Getenv("SECRET_KEY"))
}
