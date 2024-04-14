package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	//slice de byte de 64 posicoes
// 	chave := make([]byte, 64)
// 	fmt.Println(chave)

// 	if _, erro := rand.Read(chave); erro != nil {
// 		log.Fatal(erro)
// 	}
// 	fmt.Println(chave)

// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println(stringBase64)
// }

func main() {
	config.Carregar()

	r := router.Gerar()

	fmt.Printf("Server listen port: %d🚀\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
