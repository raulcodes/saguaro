package app_builder

import (
	"fmt"
	"io/ioutil"
	"os"
)

func BuildApp() error {
	if _, err := os.Stat("document/body.html"); os.IsNotExist(err) {
		return fmt.Errorf("Could not find a valid Saguaro app in this directory")
	}
	if _, err := os.Stat("document/head.html"); os.IsNotExist(err) {
		return fmt.Errorf("Could not find a valid Saguaro app in this directory")
	}

	index, err := os.Create("public/index.html")
	if err != nil {
		return fmt.Errorf("An error occured when building index.html")
	}

	// htmlComponentFiles, err := scanComponents("components/")
	// if err != nil {
	// 	return fmt.Errorf("An error occured when scanning components")
	// }
	// for _, file := range htmlComponentFiles {
	// 	fmt.Println(file)
	// }

	headData, _ := ioutil.ReadFile("document/head.html")
	index.Write(headData)

	bodyData, _ := ioutil.ReadFile("document/body.html")
	index.Write(bodyData)

	// fmt.Print(validateHTML([]byte("<hello")))

	fmt.Println("App built successfully.")
	return nil
}
