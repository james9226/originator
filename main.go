package main

import (
	"lendotopia.com/originator/config"
	"lendotopia.com/originator/router"
	"lendotopia.com/originator/services/firestore"
)

func main() {

	//load .env
	config.Load()

	//Initialize firestore client
	firestore.Init()
	defer firestore.Client.Close()

	//Initialize the API router
	router.Init()
}
