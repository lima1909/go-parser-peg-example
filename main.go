package main

import (
	"fmt"
	"log"

	"github.com/lima1909/go-parser-peg-example/tx"
)

func main() {
	txt := `2009-01-21 hi gophers
 - gopher sticker
2019-09-19 buy books
 - I Know Why The Caged Bird Sings
 - Go in Action`

	fmt.Println("--> Input:")
	fmt.Println(txt)

	r, err := tx.Parse("example", []byte(txt))
	if err != nil {
		log.Printf("err by parsing: %v", err)
		return
	}

	fmt.Println("")
	fmt.Println("--> Result:")
	txs := r.(tx.Transactions)
	for _, t := range txs {
		fmt.Printf(" %v\n", t)
		for _, i := range t.Items {
			fmt.Printf("\t%v\n", i)
		}
	}
}
