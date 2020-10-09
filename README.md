# Data Frames for Go 
[![Go Report Card](https://goreportcard.com/badge/github.com/AdikaStyle/go-df)](https://goreportcard.com/report/github.com/AdikaStyle/go-df)


An effort to make a fully featured data frame/wrangling library  

## Example:
```go
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
```

Will print the following:
```
+--------------------+----------+-----------+------------------------+
|     UnitPrice      | Quantity | ProductID |      ProductName       |
+--------------------+----------+-----------+------------------------+
|  786.5999999999999 |     1057 |         2 | Chang                  |
|  704.2000000000004 |     1158 |        16 | Pavlova                |
|  706.2999999999995 |     1103 |        40 | Boston Crab Meat       |
| 1770.7999999999997 |     1263 |        56 | Gnocchi di nonna Alice |
|               2761 |     1496 |        59 | Raclette Courdavault   |
|             1638.8 |     1577 |        60 | Camembert Pierrot      |
| 2227.7999999999993 |     1083 |        62 | Tarte au sucre         |
|  829.8999999999999 |     1057 |        71 | Flotemysost            |
+--------------------+----------+-----------+------------------------+
Dataframe has 4 columns and 8 rows in total.

Left side of split:
+--------------------+----------+-----------+------------------------+
|     UnitPrice      | Quantity | ProductID |      ProductName       |
+--------------------+----------+-----------+------------------------+
|  706.2999999999995 |     1103 |        40 | Boston Crab Meat       |
| 1770.7999999999997 |     1263 |        56 | Gnocchi di nonna Alice |
|               2761 |     1496 |        59 | Raclette Courdavault   |
|             1638.8 |     1577 |        60 | Camembert Pierrot      |
| 2227.7999999999993 |     1083 |        62 | Tarte au sucre         |
|  829.8999999999999 |     1057 |        71 | Flotemysost            |
+--------------------+----------+-----------+------------------------+
Dataframe has 4 columns and 6 rows in total.

Right side of split:
+-------------------+----------+-----------+-------------+
|     UnitPrice     | Quantity | ProductID | ProductName |
+-------------------+----------+-----------+-------------+
| 786.5999999999999 |     1057 |         2 | Chang       |
| 704.2000000000004 |     1158 |        16 | Pavlova     |
+-------------------+----------+-----------+-------------+
Dataframe has 4 columns and 2 rows in total.
```

## Features
- Predefined type system: Boolean, Integer, Decimal, String, DateTime, Missing.
- Select columns.
- Filter by custom condition.
- Concatenate two data frames.
- Split by custom condition.
- Group by multiple columns with field aggregations.
- Order by single field (ASC/DESC).
- Inner join with another data frame.
- Left join with another data frame.
- Right join with another data frame.
- [TODO] Outer join with another data frame.
- Pretty Print a data frame.
- Iterate through rows.
- Iterate through columns.
- Easy to extend design.

## Supported Adapters
- From CSV.
- To CSV.
- [TODO] From SQL.
- [TODO] To SQL.
- [TODO] From Structs.
- [TODO] To Structs.
