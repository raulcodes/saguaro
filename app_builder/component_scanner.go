package app_builder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	path      string
	extension string
}

type Component struct {
	name           string
	rootPath       string
	componentFiles []File
}

func scanComponents(rootPath string) ([]Component, error) {
	var components []Component
	
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		extension := filepath.Ext(path)
		base := filepath.Base(path)
		if extension == ".html" {
			var files []File

			componentName := strings.Replace(base, ".html", "", 1)
			for _, component := range components {
				if componentName == component.name {
					return fmt.Errorf("Error scanning components. Make sure that all of your components have unique names")
				}
			}
			htmlFile := File{path, extension}
			files = append(files, htmlFile)
			component := Component{componentName, filepath.Dir(path), files}
			components = append(components, component)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return components, nil
}