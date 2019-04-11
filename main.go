package main

import (
	"log"
	db "go_postgresql/db"//good
	//db "./db" //video
)

func main() {
	log.Printf("Hello World!\n")
	db.Connect()
}

