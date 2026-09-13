package ledger

import (
	"errors"
	"sync"
	"time"
)

type Transaction struct {
	ID     int       `json:id"`
	Date   time.Time `json:"time"`
	Amount float64   `json:"price"`
	Note   string    `json:"note,omitempty"`
}

type Ledger struct {
	mu           sync.Mutex
	transactions []Transaction `json:"transactions"`
}

func New() *Ledger {
	return &Ledger{}
}

func (l *Ledger) Add(t Transaction) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.transactions = append(l.transactions, t)
}

func (l *Ledger) List() []Transaction {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.transactions
}

func (l *Ledger) Get(id int) (Transaction, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, t := range l.transactions {
		if t.ID == id {
			return t, nil
		}
	}
	return Transaction{}, errors.New("Not found")
}
