package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	initializeEnvironment()
	app := bootstrapApp()
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(app.Run(fmt.Sprintf(":%d", port)))
}
