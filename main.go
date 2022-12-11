package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	//"gopkg.in/yaml.v2"
)

type rawAuthors struct {
	Authors []Author `json:"authors"`
}

type Author struct {
	IDENTIFICATION string `json:"key"`
}

type rawWorks struct {
	Entries []Title `json:"entries"`
}

type Title struct {
	NAME string `json:"title"`
}

type Name struct {
	NAME string `json:"personal_name"`
}

type Revisions struct {
	REVISION int `json:"revision"`
}

func main() {

	var ISBN string

	fmt.Printf("Hi ! Welcome to the simple Book Fetcher application \n")
	fmt.Printf("Please input ISBN number of desired book \n")

	fmt.Scan(&ISBN)

	bookAPI := fmt.Sprintf("https://openlibrary.org/isbn/%v.json", ISBN)
	response, err := http.Get(bookAPI)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data rawAuthors
	json.Unmarshal(responseData, &data)

	authorIdentifier := data.Authors[0]

	plainAPI := "https://openlibrary.org"
	//for fetching works
	newAPI := fmt.Sprintf("%s%s/works.json", plainAPI, authorIdentifier.IDENTIFICATION)

	// for fetching names
	newAPI1 := fmt.Sprintf("%s%s.json", plainAPI, authorIdentifier.IDENTIFICATION)

	// getting works
	response2, err := http.Get(newAPI)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData2, err := io.ReadAll((response2.Body))

	if err != nil {
		log.Fatal(err)
	}

	var data2 rawWorks
	json.Unmarshal(responseData2, &data2)

	//getting name
	response3, err := http.Get(newAPI1)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData3, err := io.ReadAll(response3.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data3 Name
	json.Unmarshal(responseData3, &data3)

	//getting revisions
	response4, err := http.Get(newAPI1)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData4, err := io.ReadAll(response4.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data4 Revisions
	json.Unmarshal(responseData4, &data4)

	fmt.Printf("Name of the author is %s \n", data3.NAME)
	fmt.Printf("His/her works are  : %s\n With number of revisions : %v\n", data2.Entries, data4.REVISION)

}
