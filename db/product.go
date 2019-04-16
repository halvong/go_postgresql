package db

import (
	"log"
	"time"
	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

type ProductItem struct {
	RefPointer int `sql:"-"`
	tableName struct{} `sql:"product_items_collection"`
	ID int `sql:"id,pk"`
	Name string `sql:"name,unique"`
	Desc string `sql:"desc"`
	Image string `sql:"image"`
	Price float64 `sql:"price,type:real"`
	Features struct {
		Name string
		Desc string	
		Imp int
	} `sql:"features,type:json"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive bool `sql:"is_active"`
}

//----
func (pi *ProductItem) Save(db *pg.DB) error {

	insertErr := db.Insert(pi)

	if insertErr != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", insertErr)
		return insertErr
	} 
	log.Printf("ProductItem %s inserted successfully.\n", pi.Name)	
	return nil
}

//----
func (pi *ProductItem) SaveAndReturn(db *pg.DB) (*ProductItem, error) {

	InsertResult, insertErr := db.Model(pi).Returning("*").Insert()

	if insertErr != nil {
		log.Printf("Error while inserting new item, Reason: %v\n", insertErr)
		return nil, insertErr 
	}	

	log.Printf("ProductIte Inserted successfully.\n")
	log.Printf("Rows affected: %v\n", InsertResult.RowsAffected())
	return pi, nil
}

//----
func CreateProdItemsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := db.CreateTable(&ProductItem{}, opts)
	if createErr != nil {
		log.Printf("Error while creating table productItems, Reason: %v\n", createErr)	
		return createErr
	}
	log.Printf("Table ProductItems created successfully.\n")
	return nil
}

//----
func(pi *ProductItem) SaveMultiple(db *pg.DB, items ...*ProductItem) error {

	_, insertErr := db.Model(items).Insert()

	if insertErr != nil {
		log.Printf("Error while inserting bulk items, Reason: %v\n", insertErr)	
		return insertErr
	}
	log.Printf("Buil Insert successful\n")
	return nil	
}