package http

import (
	"net/http"

	"github.com/Paulo202005/planejamento-financeiro/adapter/http/actuator"
	"github.com/Paulo202005/planejamento-financeiro/adapter/http/transaction"
)

// Init Pacote inicial
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}
