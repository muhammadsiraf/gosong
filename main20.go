package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name":"john wick", "Age":27}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age :", data.Age)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age : ", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Println("user : ", decodedData["Name"])
	fmt.Println("user : ", decodedData["Age"])

	var jsonString1 = `[
		{"Name":"john pantau", "Age":27}, 
		{"Name":"ethan 0", "Age":32} 
	]`

	var data3 []User

	err = json.Unmarshal([]byte(jsonString1), &data3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1:", data3[0].FullName)
	fmt.Println("user 2:", data3[1].FullName)

}
