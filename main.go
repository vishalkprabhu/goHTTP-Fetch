package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Geo list
type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

//Company list
type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

//Address list
type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

//User list
type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Company  Company `json:"company"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
}

func getResponseData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	return contents
}

func main() {
	contents := getResponseData("http://jsonplaceholder.typicode.com/users")
	var users []User
	json.Unmarshal(contents, &users)

	//print users geo location
	for _, user := range users {
		fmt.Printf("Username: %s - EMAIL: %s - Lat: %s Lng: %s \n", user.Name, user.Email, user.Address.Geo.Lat, user.Address.Geo.Lng)
	}
}
