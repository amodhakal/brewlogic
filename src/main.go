package main

import (
	"amodhakal/brewlogic/src/config"
	"log"
)

func main() {
	result, err := config.Query.GetUsers(config.DBContext)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(result)
}
