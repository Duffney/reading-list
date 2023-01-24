package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Book struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Published int
	Pages     string
	Genres    []string
	Version   int32
}

type Envelope struct {
	Book Book `json:"book"`
}

func (b *Book) Get(id int) (*Book, error) {
	endpoint := fmt.Sprintf("http://localhost:4000/v1/books/%d", id)
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal the JSON data into the envelope struct
	var envelope Envelope
	err = json.Unmarshal(body, &envelope)
	if err != nil {
		fmt.Println(err)
	}

	// Extract the book data from the envelope
	book := envelope.Book

	// Print the struct
	return &book, nil
}
