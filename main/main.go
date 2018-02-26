package main

import (
	"fmt"
	"github.com/rusucosmin/goworkshop/importer"
	"github.com/rusucosmin/goworkshop/web"
	"github.com/rusucosmin/goworkshop/model"
)

func main() {
	model.Authors = importer.ImportAuthors()
	fmt.Printf("Imported authors are: %s\n", model.Authors)
	model.Books = importer.ImportBooks()
	fmt.Printf("Imported books are: %s\n", model.Books)
	web.StartServer()
}
