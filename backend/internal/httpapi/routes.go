package httpapi

import (
	"net/http"
)

func (a *API) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /transactions/", a.getTransactions)
	mux.HandleFunc("GET /transactions/{id}/", a.getTransactionByID)
	// mux.HandleFunc("POST /transactions", postTransaction)
	return mux
}
