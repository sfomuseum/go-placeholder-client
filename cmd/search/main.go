package main

import (
	"flag"
	"fmt"
	"github.com/sfomuseum/go-placeholder-client"
	"github.com/sfomuseum/go-placeholder-client/filters"
	"log"
)

func main() {

	endpoint := flag.String("placeholder-endpoint", client.DEFAULT_ENDPOINT, "...")

	var search_filters filters.SearchFilters
	flag.Var(&search_filters, "filter", "...")

	flag.Parse()

	cl, err := client.NewPlaceholderClient(*endpoint)

	if err != nil {
		log.Fatal(err)
	}

	for _, term := range flag.Args() {

		results, err := cl.Search(term, search_filters...)

		if err != nil {
			log.Fatal(err)
		}

		for _, r := range results.Results() {
			fmt.Println(r)
		}
	}
}
