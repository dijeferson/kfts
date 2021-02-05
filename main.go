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
		panic(err)
	}

	proc := provider.Processor{}
	result, err := proc.Process(*contents)

	if err != nil {
		panic(err)
	}

	for _, v := range *result {
		if v.Note != "" {
			fmt.Printf("Book: %v - Note: %+v\n", v.Book, v.Note)
		}
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
