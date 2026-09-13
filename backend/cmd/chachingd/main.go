package main

import (
	"cmd/chachingd/internal/httpapi"
	"cmd/chachingd/internal/ledger"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	l := ledger.New()
	l.Add(ledger.Transaction{ID: 1, Date: time.Now(), Amount: 1.02, Note: "Drink"})
	l.Add(ledger.Transaction{ID: 2, Date: time.Now(), Amount: 100, Note: "Gym"})
	l.Add(ledger.Transaction{ID: 3, Date: time.Now(), Amount: 1200})

	fmt.Println("Starting server on port 10101")
	h := httpapi.New(l)

	srv := &http.Server{
		Addr:              "localhost:10101",
		Handler:           h,
		ReadHeaderTimeout: 5 * time.Second,
	}
	fmt.Println("Server started!")
	log.Fatal(srv.ListenAndServe())
}
