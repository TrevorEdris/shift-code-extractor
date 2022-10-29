package main

import (
	"github.com/TrevorEdris/shift-code-extractor/app/services"
)

func main() {
	a, err := services.NewApp()
	if err != nil {
		panic(err)
	}

	err = a.Run()
	if err != nil {
		panic(err)
	}
}
