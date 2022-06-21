package transaction

import "time"

// Transaction XXXXX
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// Transactions XXXXXX
type Transactions []Transaction
