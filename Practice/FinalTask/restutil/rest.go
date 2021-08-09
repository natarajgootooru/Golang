package restutil

import (
	"io"
	"log"
	"net/http"
)

const restEndpoint string = "https://jsonplaceholder.typicode.com"

func DoGET(queryUri string) []byte {

	resp, err := http.Get(restEndpoint + queryUri)
	if err != nil {
		log.Fatal("Unable to query data for URI: " + queryUri)
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Unable to query data for URI: " + queryUri)
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Unable to pull data from REST API")
	}
	return body
}
