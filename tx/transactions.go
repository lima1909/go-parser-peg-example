package tx

import "fmt"

//go:generate pigeon -o parser.go transactions.peg

// Transactions list of transaction
type Transactions []Transaction

// Transaction bracket for all items with a ID and description
type Transaction struct {
	ID    string // 0000-00-00 (like date: year-month-day)
	Text  string
	Items []Item
}

// Item description for one thing to buy
type Item struct {
	Name string
}

func (t Transaction) String() string {
	return fmt.Sprintf("%s - %s", t.ID, t.Text)
}

func (i Item) String() string {
	return fmt.Sprintf("%s", i.Name)
}
