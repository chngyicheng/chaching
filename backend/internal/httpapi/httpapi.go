package httpapi

import (
	"cmd/chachingd/internal/ledger"
	"fmt"
	"net/http"
	"strconv"
)

type API struct {
	ledger *ledger.Ledger
}

func New(l *ledger.Ledger) http.Handler {
	a := &API{ledger: l}
	return a.routes()
}

func (a *API) getTransactions(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s\n", a.ledger.List())
}

func (a *API) getTransactionByID(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	t, err := a.ledger.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintf(w, "%s\n", t)
}
