// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package products

type Product struct {
	Upc    string  `json:"upc"`
	Name   *string `json:"name"`
	Price  *int    `json:"price"`
	Weight *int    `json:"weight"`
}

func (Product) IsEntity() {}
