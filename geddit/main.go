package main

import (
	"fmt"
	"log"

	"github.com/liqlvnvn/go-reddit"
)

func main() {
	items, err := reddit.Get("goland")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
