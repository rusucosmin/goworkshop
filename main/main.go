package main

import "fmt"

import (
	"io/ioutil"
	"encoding/json"
	"github.com/rusucosmin/goworkshop/model"
)

func main() {
	fileContent, err := ioutil.ReadFile("model/books.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileContent, &model.Books)

	if err != nil {
		panic(err)
	}

	serializedData, err := json.Marshal(model.Books)
	if err != nil {
		panic(err)
	}

	fmt.Println("The serialized books are:")
	fmt.Println(string(serializedData))

	fileContent, err = ioutil.ReadFile("model/authors.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileContent, &model.Authors)

	if err != nil {
		panic(err)
	}

	serializedData, err = json.Marshal(model.Authors)
	if err != nil {
		panic(err)
	}

	fmt.Println("The serialized books are:")
	fmt.Println(string(serializedData))

}
