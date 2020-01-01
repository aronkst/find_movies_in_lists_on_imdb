package main

import (
	"fmt"
)

func main() {
	document, err := loadSite("")
	if err != nil {
		panic(err)
	}
	fmt.Println(document)
}
