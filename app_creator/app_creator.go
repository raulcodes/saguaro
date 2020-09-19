package app_creator

import (
	"fmt"
	"os"
)

func CreateApp(name string) error {
	// Future
	// err := createComponentsDir(name)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("Successfully created components directory")
	
	err := createDocumentDir(name)
	if err != nil {
		return err
	}
	fmt.Println("Successfully created document directory")

	err = createPublicDir(name)
	if err != nil {
		return err
	}
	fmt.Println("Successfully created public directory")
	
	fmt.Printf("Done creating Saguaro project: %v\n", name)
	return nil
}

// func createComponentsDir(name string) error {
// 	componentDir := fmt.Sprintf("%s/components", name)
// 	err := os.MkdirAll(componentDir, 0755)
// 	if err != nil {
// 		return err
// 	}

// 	componentFilePath := fmt.Sprintf("%s/components/HelloWorld.html", name)
// 	componentFile, err := os.Create(componentFilePath)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = componentFile.Write([]byte(HelloWorld))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func createDocumentDir(name string) error {
	documentDir := fmt.Sprintf("%s/document", name)
	err := os.MkdirAll(documentDir, 0755)
	if err != nil {
		return err
	}

	bodyFilePath := fmt.Sprintf("%s/document/body.html", name)
	bodyFile, err := os.Create(bodyFilePath)
	if err != nil {
		return err
	}

	_, err = bodyFile.Write([]byte(BodyTemplate))
	if err != nil {
		return err
	}

	headFilePath := fmt.Sprintf("%s/document/head.html", name)
	headFile, err := os.Create(headFilePath)
	if err != nil {
		return err
	}

	_, err = headFile.Write([]byte(HeadTemplate))
	if err != nil {
		return err
	}
	
	return nil
}

func createPublicDir(name string) error {
	publicDir := fmt.Sprintf("%s/public", name)
	err := os.Mkdir(publicDir, 0755)
	if err != nil {
		return err
	}
	
	return nil
}