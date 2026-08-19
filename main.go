package main

import (
	"github.com/PhateValleyman/ccat/v2/internal/app"
	"log"
	"os"
)

func main() {
	if err := app.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
