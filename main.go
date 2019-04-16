package main

import (
	"log"
	"time"
	db "go_postgresql/db"
	pg "github.com/go-pg/pg"
)

func main() {
	log.Printf("Hello World!\n")
	pg_db := db.Connect()
	SaveProduct(pg_db)
}

func SaveProduct(dbRef *pg.DB) {

	newPI1 := &db.ProductItem {
		Name: "Product 4",
		Desc: "Product 4 desc",
		Image: "This is image path",
		Price: 4.44,
		Features: struct {
			Name string	
			Desc string	
			Imp int	
		}{
			Name: "F4",	
			Desc: "F4 Desc",
			Imp : 4,
		},
		CreatedAt: time.Now(),	
		UpdatedAt: time.Now(),	
		IsActive: true,
	}

	newPI2 := &db.ProductItem {
		Name: "Product 5",
		Desc: "Product 5 desc",
		Image: "This is image path",
		Price: 5.55,
		Features: struct {
			Name string	
			Desc string	
			Imp int	
		}{
			Name: "F5",	
			Desc: "F5 Desc",
			Imp : 5,
		},
		CreatedAt: time.Now(),	
		UpdatedAt: time.Now(),	
		IsActive: true,
	}

	newPI3 := &db.ProductItem {
		Name: "Product 6",
		Desc: "Product 6 desc",
		Image: "This is image path",
		Price: 6.33,
		Features: struct {
			Name string	
			Desc string	
			Imp int	
		}{
			Name: "F6",	
			Desc: "F6 Desc",
			Imp : 6,
		},
		CreatedAt: time.Now(),	
		UpdatedAt: time.Now(),	
		IsActive: true,
	}

	_, err = newPI1.SaveMultiple(dbRef, newPI2, newPI3)

	if err != nil {
		log.Printf("Error from save.\n")		
	}
}

