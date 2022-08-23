package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dijeferson/kfts/provider"
)

func main() {
	contents, err := readFile("sample.txt")
	if err != nil {
		log.Fatalln("Error reading input file. ", err)
	}

	proc := provider.Processor{}
	result, err := proc.Process(*contents)

	if err != nil {
		log.Fatalln("Error processing input file", err)
	}

	// table := termtables.CreateTable()
	// table.SetModeTerminal()
	// table.AddHeaders("Title", "Note")

	// for _, v := range *result {
	// 	table.AddRow(v.Book.Title, v.Note)
	// 	table.AddSeparator()
	// 	// if v.Note != "" {
	// 	// fmt.Printf("Book: %v - Note: %+v\n", v.Book, v.Note)
	// 	// }
	// }

	// fmt.Println(table.Render())

	for idx, value := range *result {
		fmt.Printf("%v > %v\n", idx, value)
	}
}

func readFile(path string) (*string, error) {
	f, err := os.Open("sample.txt")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	contents := string(b)
	return &contents, nil
}
