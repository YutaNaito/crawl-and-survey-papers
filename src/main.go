package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://export.arxiv.org/api/query?search_query=cat:cs.RO+OR+cs.CV&start=0&max_results=100&sortBy=submittedDate&sortOrder=descending")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
