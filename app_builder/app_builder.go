package app_builder

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/net/html"
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

	buildAndWriteComponent("document/head.html", index)
	buildAndWriteComponent("document/body.html", index)

	fmt.Println("App built successfully.")
	return nil
}

func buildAndWriteComponent(componentPath string, destination *os.File) error {
	data, err := ioutil.ReadFile(componentPath)
	if err != nil {
		return fmt.Errorf("Error building component at %s", componentPath)
	}

	err = validateHTML(data)
	if err != nil {
		return err
	}

	destination.Write(data)

	return nil
}

func validateHTML(htmlData []byte) error {
	htmlDataReader := strings.NewReader(string(htmlData))

	_, err := html.Parse(htmlDataReader)
	if err != nil {
		return fmt.Errorf("Error parsing HTML")
	}

	return nil
}
