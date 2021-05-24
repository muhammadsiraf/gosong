package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type pelajar struct {
	ID    string
	Name  string
	Grade int
}

var baseURL = "http://localhost:8000"

func main() {
	var users, err = fetchUser()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	for _, each := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}

	var user1, err1 = fetchUser1("E001")
	if err1 != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)
}

func fetchUser() ([]pelajar, error) {
	var err error
	var client = &http.Client{}
	var data []pelajar

	request, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchUser1(ID string) (pelajar, error) {
	var err error
	var client = &http.Client{}
	var data pelajar

	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/user", payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}
