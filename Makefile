.PHONY: peg
peg:
	pigeon -o tx/parser.go tx/transactions.peg

.PHONY: gen
gen:
	go generate tx/transactions.go