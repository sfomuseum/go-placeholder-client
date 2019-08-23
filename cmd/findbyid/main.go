package main

import (
	"flag"
	"fmt"
	"github.com/sfomuseum/go-placeholder-client"
	"log"
)

func main() {

	endpoint := flag.String("placeholder-endpoint", client.DEFAULT_ENDPOINT, "...")

	flag.Parse()

	cl, err := client.NewPlaceholderClient(*endpoint)

	if err != nil {
		log.Fatal(err)
	}

	ids := flag.Args()	

	results, err := cl.FindById(ids...)

	if err != nil {
		log.Fatal(err)
	}

	for _, r := range results.Results() {
		fmt.Println(r)
	}
}
