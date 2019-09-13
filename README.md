# this is a example for using a peg parser

For the example I use the peg parser implementation: [pigeon](https://github.com/mna/pigeon)

For the first steps a youtube video: [GopherCon 2019 Lightning Talk: Tim Raymond - Parsing Expression Grammar](https://www.youtube.com/watch?v=rwzhu5YegpE)

## example -transactions:

### structs:

```go
type Transactions []Transaction

type Transaction struct {
    ID    string // 0000-00-00 (like date: year-month-day)
    Text  string
    Items []Item
}

type Item struct {
    Name      string
}
```

### example:

the structure from this example is:
```
[TX-ID] [TX-TEXT]
 [ITEM (NAME)]
```

concrete:
```
2019-09-12 gophers
 - gopher sticker
2019-09-19 buy books
 - I Know Why The Caged Bird Sings
 - Go in Action
```

