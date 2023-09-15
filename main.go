package main

import (
	"log"
	"os"
	"zerofp/models"
)

func main() {
	models.InitDB()
	err := os.Setenv("mysqlDsn", "go:admin666@tcp(43.139.40.48:3306)")
	if err != nil {
		log.Println("failed")
		log.Fatalln(err)
	}
	log.Println("success")
}
