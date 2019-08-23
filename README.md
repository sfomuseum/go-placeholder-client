# go-placeholder-client

Go package for talking to a Placeholder endpoint.

## Install

You will need to have both `Go` (specifically version [1.12](https://golang.org/dl/) or higher) and the `make` programs installed on your computer. Assuming you do just type:

```
make tools
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Important

This is work in progress.

## Example

```
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

	cl, _ := client.NewPlaceholderClient(*endpoint)


	for _, term := range flag.Args() {

		results, _ := cl.Search(term, search_filters...)

		for _, r := range results.Results() {
			fmt.Println(r)
		}
	}
}

```

_Error handling has been left out for the sake of brevity._

## Tools

### findbyid

_I am still working through some parsing issues with this one._

### query

```
$> go run cmd/query/main.go gowanus
85865587
```

### search

```
$> go run cmd/search/main.go gowanus
Gowanus, neighbourhood (85865587)

$> go run cmd/search/main.go -filter lang=fra germany
Allemagne, country (85633111)
Germany, neighbourhood (85821415)
Germany Township, localadmin (404484827)
Camp Dennison, locality (1293109707)
Armonai, locality (1276690151)
Germany, locality (1277080253)
Pearl, locality (1209234101)
Germany, locality (1293498807)
Buena Vista, locality (1327344629)
Germany, locality (1343296991)
Germany, locality (1343665787)
Germany, locality (1344073249)

$> go run cmd/search/main.go -filter lang=jpn germany
ドイツ, country (85633111)
Germany, neighbourhood (85821415)
Germany Township, localadmin (404484827)
Camp Dennison, locality (1293109707)
Armonai, locality (1276690151)
Germany, locality (1277080253)
Pearl, locality (1209234101)
Germany, locality (1293498807)
Buena Vista, locality (1327344629)
Germany, locality (1343296991)
Germany, locality (1343665787)
Germany, locality (1344073249)

$> go run cmd/search/main.go -filter placetype=country -filter lang=zho germany
德国, country (85633111)
```

### tokenize

```
go run cmd/tokenize/main.go 'sydney new south wales'
sydney
new south wales
```

## See also

* https://github.com/pelias/placeholder