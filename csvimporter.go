package main

import (
	"encoding/csv"
	"fmt"
	//"io"
	"log"
	"os"
	//"strings"
)

func start(fileName string)([][]string){
	file, err := os.Open(fileName)
	if err != nil{
		fmt.Println("There was an error")
		log.Fatal(err)
	}
	reader := csv.NewReader(file)

	firstLine, err := reader.ReadAll()
	if err != nil{
		fmt.Println("There was an error")
		log.Fatal(err)
	}
	return firstLine
}