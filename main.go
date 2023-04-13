package main

import (
	"backend/src/config/command"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Your aplication run on port 8081")

	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}