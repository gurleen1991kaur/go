package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Accounts struct {
	Account_ID      int `json:"account_id"`
	Document_Number int `json:"document_number"`
}

type Transactions struct {
	Transaction_ID int       `json:"transaction_id"`
	Account_ID     int       `json:"account_id"`
	Operation_Type int       `json:"operation_type"`
	Amount         float32   `json:"amount"`
	EventDate      time.Time `json:"date"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "gurleen:Jaibir2021123#@tcp(127.0.0.1:3306)/cards?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/accounts", getAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}", getAccountById).Methods("GET")
	router.HandleFunc("/accounts", createAccount).Methods("POST")
	router.HandleFunc("/transactions", getTran).Methods("GET")
	router.HandleFunc("/transactions", createTran).Methods("POST")
	http.ListenAndServe(":8000", router)
}
func getAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var accounts []Accounts
	result, err := db.Query("SELECT * from accounts")
	fmt.Println(result)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var account Accounts
		err := result.Scan(&account.Account_ID, &account.Document_Number)
		if err != nil {
			panic(err.Error())
		}
		accounts = append(accounts, account)
	}
	json.NewEncoder(w).Encode(accounts)
}

func getAccountById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT account_id, document_number FROM accounts WHERE account_id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var account Accounts
	for result.Next() {
		err := result.Scan(&account.Account_ID, &account.Document_Number)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(account)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		account_id := r.FormValue("account_id")
		document_number := r.FormValue("document_number")
		insForm, err := db.Prepare("INSERT INTO accounts(Account_ID, Document_Number) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		intacc, err := strconv.Atoi(account_id)
		if err != nil {
			panic(err.Error())
		}
		intdoc, err := strconv.Atoi(document_number)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(intacc, intdoc)

		insForm.Exec(intacc, intdoc)
		fmt.Println("INSERT: account_id: " + account_id + " | Document_Number: " + document_number)

	}
	fmt.Fprintf(w, "New account was created")
}

func getTran(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transactions []Transactions
	result, err := db.Query("SELECT * from transactions")
	fmt.Println(result)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var transaction Transactions
		err := result.Scan(&transaction.Transaction_ID, &transaction.Account_ID, &transaction.Operation_Type, &transaction.Amount, &transaction.EventDate)
		if err != nil {
			panic(err.Error())
		}
		transactions = append(transactions, transaction)
	}
	json.NewEncoder(w).Encode(transactions)
}

func createTran(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		transaction_id := r.FormValue("transaction_id")
		account_id := r.FormValue("account_id")
		operationtype_id := r.FormValue("operationtype_id")
		amount := r.FormValue("amount")
		// eventdate := r.FormValue("eventdate")
		insForm, err := db.Prepare("INSERT INTO transactions(Transaction_ID, Account_ID, OperationType_ID, Amount, EventDate) VALUES(?,?,?,?, NOW())")
		if err != nil {
			panic(err.Error())
		}
		inttran, err := strconv.Atoi(transaction_id)
		if err != nil {
			panic(err.Error())
		}
		intaccid, err := strconv.Atoi(account_id)
		if err != nil {
			panic(err.Error())
		}
		intopt, err := strconv.Atoi(operationtype_id)
		if err != nil {
			panic(err.Error())
		}
		intamt, err := strconv.ParseFloat(amount, 32)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(inttran, intaccid, intopt, intamt)

		insForm.Exec(inttran, intaccid, intopt, intamt)
		fmt.Println("INSERT: transaction_id: " + transaction_id + " | account_id: " + account_id)

	}
	fmt.Fprintf(w, "New transaction was created")
}
