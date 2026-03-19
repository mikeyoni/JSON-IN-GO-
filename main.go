package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// we need to make the structure of the json object so we can decode it

	type user struct {
		UserId    int    `json:"userId"`
		Email     string `json:"email"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	// reading the json file form  the os

	bytes, err := os.ReadFile("db.json")

	// error handling for the file reading

	if err != nil {

		fmt.Printf(" file not in directory !", err)
		return
	}

	// giving the structure to a variable to store the decoded data

	var jsondata []user

	// decoding the json data and storing it in the variable and also error handeling with error variable

	err = json.Unmarshal(bytes, &jsondata)

	// error handling for the unmarshalling of the data

	if err != nil {
		fmt.Printf("error in unmarshalling data !", err)
		return
	}

	fmt.Printf("%+v", jsondata)

	jdata, err := json.MarshalIndent(jsondata, "", "\t")

	fmt.Printf(string(jdata))

	// email := "mda78@gmail.com"

	// if email != jsondata[0].Email{

	// 	fmt.Printf(" \n email does not match ! \n")
	// 	return

	// }else{
	// 	fmt.Printf(" \n email matches ! \n ")

	// }

	// encoding the data into json format

	newUser := user{
		UserId:    2,
		Email:     "mikey2334@gmail.com",
		Title:     "qui est esse",
		Completed: false,
	}

	isexist := false

	for _, exitsuser := range jsondata {

		if exitsuser.Email == newUser.Email {
			isexist = true
			break
		}

	}


	if isexist {

		fmt.Printf(" \n email already exists ! \n")
		return

	} else {

		jsondata = append(jsondata, newUser)

		updateddata, _ := json.MarshalIndent(jsondata, "", "\t")

		fmt.Printf(string(updateddata))

		err = os.WriteFile("db.json", updateddata, 0644)
		if err != nil {
			fmt.Printf("error in writing to file !", err)
			return
		}
	}

	// decoding the json data

	// var haha user

	// err = json.Unmarshal(jdata, &haha)

	// if err != nil {

	// 	fmt.Printf("error in unmarshalling data !", err)
	// }

	// fmt.Printf("%+v", haha)

}
