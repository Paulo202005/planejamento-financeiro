package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Paulo202005/planejamento-financeiro/model/transaction"
	"github.com/Paulo202005/planejamento-financeiro/util"
)

// GetTransactions XXXXX
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salario",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.StringToTime("2021-12-31T17:15:00"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

// CreateATransaction XXXXXX
func CreateATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
