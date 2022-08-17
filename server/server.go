package main

import (
	"MAS/server/database"
	"MAS/server/routes"
	"flag"
	"fmt"
)

func main() {

	var port string
	var migrate bool

	//Get command line arguments
	flag.StringVar(&port, "port", "8080", "Define the starting port. Default is 8080.")
	flag.BoolVar(&migrate, "migrate", false, "Migrate the database with all the schemas. Default is false.")
	flag.Parse()

	if err := database.Connect(migrate); err != nil {
		fmt.Println(err.Error())
		return
	}

	router := routes.CreateRouter()
	router.Run(":" + port)

}
