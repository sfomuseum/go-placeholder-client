package main

import (
	"encoding/json"
	"flag"
	"github.com/sfomuseum/go-csvdict"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {

	placeholder_endpoint := flag.String("placeholder-endpoint", "", "...")

	flag.Parse()

	seen := make(map[string]int64)

	var csv_wr *csvdict.Writer

	for _, path := range flag.Args() {

		csv_r, err := csvdict.NewReaderFromPath(path)

		if err != nil {
			log.Fatalf("Failed to create CSV reader for %s, %v", err)
		}

		for {
			row, err := csv_r.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Failed to read row, %v", err)
			}

			loc, ok := row["location"]

			if !ok {
				continue
			}

			if csv_wr == nil {

				fieldnames := make([]string, 0)

				for k, _ := range row {
					fieldnames = append(fieldnames, k)
				}

				wr, err := csvdict.NewWriter(os.Stdout, fieldnames)

				if err != nil {
					log.Fatalf("Failed to create new writer, %v", err)
				}

				csv_wr = wr
				csv_wr.WriteHeader()
			}

			seen_id, ok := seen[loc]

			if ok {
				row["wof_id"] = strconv.FormatInt(seen_id, 10)

			} else {

				search_url, err := url.JoinPath(*placeholder_endpoint, "/api/search")

				if err != nil {
					log.Fatalf("Failed to create search URL, %v", err)
				}

				req, err := http.NewRequest("GET", search_url, nil)

				if err != nil {
					log.Fatalf("Failed to create request, %v", err)
				}

				q := url.Values{}
				q.Set("text", loc)

				req.URL.RawQuery = q.Encode()

				// auth goes here...

				cl := http.Client{}
				rsp, err := cl.Do(req)

				if err != nil {
					log.Fatalf("Failed to perform request, %v", err)
				}

				var places []map[string]interface{}

				dec := json.NewDecoder(rsp.Body)

				err = dec.Decode(&places)

				if err != nil {
					log.Fatalf("Failed to decode places for %s, %v", loc, err)
				}

				if len(places) == 0 {
					continue
				}

				first := places[0]

				id := int64(first["id"].(float64))
				seen[loc] = id

				row["wof_id"] = strconv.FormatInt(id, 10)
			}

			csv_wr.WriteRow(row)
		}
	}

	csv_wr.Flush()
}
