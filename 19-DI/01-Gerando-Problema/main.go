package main

import (
	"database/sql"
	"fmt"

	"github.com/batistondeoliveira/fullcycle_curso_golang/19-DI/01-Gerando-Problema/product"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	repository := product.NewProductRepository(db)

	usecase := product.NewProductUseCase(repository)

	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
