package Controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetWalletTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT wallet.id, wallet.currency, wallet.username, wallet.password, wallet.disableUser, transaction.id, transaction.dateTime, transaction.amount, transaction.description FROM wallet wallet JOIN transaction transaction ON wallet.id = transaction.id"

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	var wallet Wallet
	var wallets []Wallet
	for rows.Next() {
		if err := rows.Scan(&wallet.WalletId, &wallet.Currency, &wallet.Username, &wallet.Password,
			&wallet.DisableUser, &wallet.Transaction.TransactionId, &wallet.Transaction.DateTime,
			&wallet.Transaction.Amount, &wallet.Transaction.Description); err != nil {
			log.Fatal(err.Error())
		} else {
			wallets = append(wallets, wallet)
		}
	}

	var response WalletResponse
	if len(wallets) > 0 {
		response.Status = 200
		response.Message = "Get Wallet Transaction Success"
		response.Data = wallets
	} else {
		response.Status = 400
		response.Message = "Get Wallet Transaction Failed!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateWallet(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	Username := r.Form.Get("username")
	Password := r.Form.Get("password")
	Currency := r.Form.Get("currency")

	vars := mux.Vars(r)
	WalletID := vars["walletId"]

	_, errQuery := db.Exec("UPDATE wallet SET username=?, password=?, currency=? WHERE id=?",
		Username,
		Password,
		Currency,
		WalletID,
	)

	var response WalletResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Update Wallet Success"
	} else {
		response.Status = 400
		response.Message = "Update Wallet Failed!\n" + errQuery.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	WalletID := vars["walletId"]
	fmt.Println(WalletID)

	_, errQuery := db.Exec("UPDATE wallet set disableUser = ? where id = ?",
		1,
		WalletID,
	)

	var response WalletResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Delete Wallet Success"
	} else {
		response.Status = 400
		response.Message = "Delete Walet Failed!\n" + errQuery.Error()
		fmt.Println(errQuery)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	IdWallet, _ := strconv.Atoi(r.Form.Get("idWallet"))
	Amount, _ := strconv.Atoi(r.Form.Get("amount"))
	Description := r.Form.Get("description")

	t := time.Now()
	timestamp := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	_, errQuery := db.Exec("INSERT INTO transaction(idWallet, dateTime, amount, description) VALUES (?, ?, ?, ?)",
		IdWallet,
		timestamp,
		Amount,
		Description,
	)

	var response TransactionResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Insert Transaction Success"
	} else {
		response.Status = 400
		fmt.Println(errQuery)
		response.Message = "Insert Transaction Failed!\n" + errQuery.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
