package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://docs.google.com/document/d/1t0Og6qLbxzSL4cNQaiwMzfGYlbanZe-CLVI2LyIYKFI/edit?tab=t.0#heading=h.r203u1gs2kvg"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
