package app_server

import (
	"fmt"
	"net/http"
)

func ServeApp() error {
  fs := http.FileServer(http.Dir("./public"))
  http.Handle("/", fs)

	fmt.Println("Listening on port 3000...")
	fmt.Println("Your app is now available at http://localhost:3000")
  err := http.ListenAndServe(":3000", nil)
  if err != nil {
    return fmt.Errorf("Error encountered while serving app")
	}

	return nil
}