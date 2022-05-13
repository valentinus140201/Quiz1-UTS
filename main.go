package main

import (
	"fmt"
	"log"
	"net/http"

	controller "Quis_PBP/Controllers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/wallet", controller.GetWalletTransaction).Methods("GET")
	router.HandleFunc("/transaction", controller.InsertTransaction).Methods("POST")
	router.HandleFunc("/wallet/{walletId}", controller.UpdateWallet).Methods("PUT")
	router.HandleFunc("/wallet/{walletId}", controller.DeleteWallet).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
