package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// User struct to hold data about a user
type User struct {
	ID        int    `json:"id"`
	CashBack  int    `json:"cashback"`
	BankName  string `json:"bankname"`
	Category  string `json:"category"`
	Condition string `json:"condition"`
	CardName  string `json:"cardname"`
}

// BankCard struct to hold data about a bank card
type BankCard struct {
	CardNumber     string `json:"card_number"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ExpirationDate string `json:"expiration_date"`
	CardType       string `json:"card_type"`
}

// users is a slice that will act as our database for this example
var users = []User{
	{ID: 1, CashBack: 5, BankName: "Jusan", Category: "Groceries", Condition: "Minimum spend of $1000 monthly", CardName: "Credit"},
	{ID: 2, CashBack: 3, BankName: "Kaspi", Category: "Baby products", Condition: "Minimum spend of $500 monthly", CardName: "Debit"},
	{ID: 3, CashBack: 4, BankName: "Halyk", Category: "Home goods", Condition: "Minimum spend of $1500 monthly", CardName: "Credit"},
	{ID: 4, CashBack: 7, BankName: "Alfa", Category: "Travel Tickets", Condition: "Minimum spend of $1500 monthly", CardName: "Credit"},
	{ID: 5, CashBack: 2, BankName: "BCC", Category: "Coffee", Condition: "Minimum spend of $1500 monthly", CardName: "Credit"},
}

// bankCards is a slice to store bank card details
var bankCards []BankCard
var mu sync.Mutex // Mutex to synchronize access to the data

// handleTable handles the GET requests at /table and sends a JSON response of users
func handleTable(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// handleAddBankCard handles the POST requests to add a new bank card.
func handleAddBankCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newCard BankCard
	err := json.NewDecoder(r.Body).Decode(&newCard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	bankCards = append(bankCards, newCard) // Append the new bank card to the slice
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCard) // Send back the added bank card data
}

func main() {
	http.HandleFunc("/table", handleTable)               // Route for getting user data
	http.HandleFunc("/add-bank-card", handleAddBankCard) // Route for adding bank card details
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
