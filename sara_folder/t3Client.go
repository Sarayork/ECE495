package sara_folder

import (
  "encoding/json"
  "log"
  "net/http"
  "fmt"
  "os"
  "bufio"
)

type Person struct {
  ID string `json:"id,omitempty"`
  Firstname string `json:"firstname,omitempty"`
  Lastname string `json:"lastname,omitempty"`
  Address *Address `json:"address,omitempty"`
}

type Address struct {
  City string `json:"city,omitempty"`
  State string `json:"state,omitempty"`
}

func main() {

  url := fmt.Sprintf("http://localhost:7000/people/2")
  req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
  // For control over HTTP client headers, redirect policy, and other settings,
	// create a Client. A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client. Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body when done reading from it
  // Defer the closing of the body
	defer resp.Body.Close()

  // Fill the record with the data from the JSON
	var text Person

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&text); err != nil {
		log.Println("zina:", err)
	}
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter name: ")
  text, _= reader.ReadString('\n')
  fmt.Println("Firstname =", text.Firstname)
	//fmt.Println("Firstname = ", test.Firstname)
	//fmt.Println("Last Name   = ", test.Lastname)
}
