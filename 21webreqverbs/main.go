package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web reqverbs:")

		// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest() 
}

func PerformGetRequest() { //task: to create a get function which will take the url as parameter
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code ", response.StatusCode)

	fmt.Println("Content length is: ", response.ContentLength)

	// content, _ := io.ReadAll(response.Body)

	var responseString strings.Builder
	content,_ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ",byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
	{
		"coursename":"Let's go with golang",
		"price":100,
		"pltform":"LearnCodeOnline"
	}
	`)

	response,err := http.Post(myurl,"application/json",requestBody)
	if err !=nil{
		panic(err)
	}
	defer response.Body.Close()

	content,_ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}


func PerformPostFormRequest()  {
	const myurl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname","yash")
	data.Add("lastname","gupta")
	data.Add("email","yg@gmil.com")

	response,err := http.PostForm(myurl, data)

	if err!=nil {
		panic(err)
	}
	defer response.Body.Close()

	content,_ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}