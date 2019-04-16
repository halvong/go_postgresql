package db

import (
	"log"
	"os"
	pg "github.com/go-pg/pg"
)

func Connect() *pg.DB {

	opts := &pg.Options {
		User: "tom",
		Password: "tom",
		Addr: "localhost:5432",
		Database: "tuts",
	}

	var db *pg.DB = pg.Connect(opts)

	//log.Println(db)
	if db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}
	
	log.Printf("Connection to database successful.\n")
	CreateProdItemsTable(db)

	return db

	//closeErr := db.Close()
	//if closeErr != nil {
	//	log.Printf("Failed to close database.\n")
	//}
}

