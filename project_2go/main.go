package main

import (
	"project_2go/db"
	"project_2go/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}