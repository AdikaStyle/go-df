package main

import (
	"fmt"
	"github.com/AdikaStyle/go-df/adapters"
	. "github.com/AdikaStyle/go-df/aggs"
	"github.com/AdikaStyle/go-df/backend"
	. "github.com/AdikaStyle/go-df/conds"
	. "github.com/AdikaStyle/go-df/dataframe"
	. "github.com/AdikaStyle/go-df/types"
	"os"
)

func main() {
	orders := adapters.LoadCSV("examples/order_details.csv", '|')

	products := adapters.LoadCSV("examples/products.csv", '|').
		Select("ProductID", "ProductName")

	view := orders.
		Group(By{"ProductID"}, Aggs{"UnitPrice": Sum(), "Quantity": Sum()}).
		OrderBy("ProductID", backend.Asc).
		InnerJoin(products, On("ProductID")).
		Filter(And(
			Gte("UnitPrice", Decimal(600)),
			Gt("Quantity", Decimal(1000))),
		).
		Print(os.Stdout)

	left, right := view.Split(Gte("ProductID", Decimal(30)))

	fmt.Printf("\nLeft side of split:\n")
	left.Print(os.Stdout)

	fmt.Printf("\nRight side of split:\n")
	right.Print(os.Stdout)
}
