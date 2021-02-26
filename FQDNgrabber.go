//This program gets all associated sub domains from a provided
// domain that the user gives in the form of an argument
// Inputs are domians like "youtube.com" or "google.com"

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	//call getDOM function
	dom := getDOM()
	fmt.Printf("\nURL given: %v\n\n", dom)

	//Call to getPage function
	page, err := getPage(dom)

	//Print out error if needed
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	display(page)
}

//Grabs URL string from arguemtns given by user
func getDOM() string {
	//Creates args variable to put user input into
	args := os.Args[1:]

	//If no argument is given, we will return a proper error
	if len(args) == 0 {
		fmt.Println("No URL provided")
		os.Exit(0)
	}

	//Creating another variable to test for integer
	args2, err := strconv.Atoi(args[0])

	//If there is an integer, we will provide an error
	if err == nil {
		fmt.Printf("%v is not a URL\n", args2)
		os.Exit(0)
	}

	//When proper input is given, we pass it back to the main function
	dom := args[0]

	return dom
}

//Calls to bufferover and gives us our wanted JSON information
func getPage(dom string) (string, error) {
	//Creates the full URL
	url := "https://dns.bufferover.run/dns?q=." + dom
	//Gets the page in JSON format
	res, err := http.Get(url)

	//If theres an error, print out error
	if err != nil {
		return "", err
	}
	//Reallocate resources
	defer res.Body.Close()

	//Reads the body of the JSON file to the page variable
	page, err := ioutil.ReadAll(res.Body)

	//If theres an error, print out error
	if err != nil {
		return "", err
	}

	//Returns the JSON file in string format
	return string(page), nil
}

//Displays the wanted JSON given by bufferover
func display(page string) {

	type Data struct {
		FDNS []string `json:"FDNS_A"`
		RDNS []string `json:"RDNS"`
	}

	dec := json.NewDecoder(strings.NewReader(page))

	//If theres an error, print out error
	var d Data
	if err := dec.Decode(&d); err != nil {
		return
	}

	//Prints out whats in the FNDS_A feild
	for i := 0; i < len(d.FDNS); i++ {
		//Use split to break ip away from url
		str := strings.Split(d.FDNS[i], ",")
		fmt.Println(str[1])
	}

	for i := 0; i < len(d.RDNS); i++ {
		//Use split to break ip away from url
		str := strings.Split(d.RDNS[i], ",")
		fmt.Println(str[1])
	}
}
