package main

//пароль пользователя постргы misha205 пароль: misha2005
import (
	"db/router"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	defer fmt.Println("Server ending")
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
