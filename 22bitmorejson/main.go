package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //(alias) i want to send this data to the json so we use this and this is sending it with the name coursename
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // '-' this dash will remove the this field to reflect to the user whoever consumes the api
	Tags     []string `json:"tags,omitempty"` //will not show the field which is empty
}

func main() {
	fmt.Println("Handling the json content")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 199, "GuptaClasses.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 99, "GuptaClasses.in", "bcd123", []string{"full-stack", "js"}},
		{"Golang Bootcamp", 599, "GuptaClasses.in", "efg123", nil},
	}

	//package this data as JSON data

	// finalJson,err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
        "coursename": "ReactJS Bootcamp",
        "Price": 199,
        "website": "GuptaClasses.in",
        "tags": ["web-dev","js"]
    }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)    //checking whether we have created the correct json format data

	if checkValid {
		fmt.Println("JSON was Valid\n")
		json.Unmarshal(jsonDataFromWeb,&lcoCourse)    //here instance is a word for the struct only nothing else 
		// fmt.Printf("%#v\n", lcoCourse)
	}else{
		fmt.Println("JSON was not Valid")
	}


	// some cases where you want to add data in key-value pair

	var myOnlineData map[string]interface{}      // actually we dont have any idea about the data in which format it is coming from the web so we are using the type interface (struct)
	json.Unmarshal(jsonDataFromWeb,&myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)



	for k,v :=range myOnlineData{
		fmt.Printf("Key is %v and value is %v and type is %T\n",k,v,v)
	}
}




