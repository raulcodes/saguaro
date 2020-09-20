package app_watcher

import (
	"fmt"
	"log"
	"time"

	"github.com/radovskyb/watcher"
  "github.com/raulcodes/saguaro/app_builder"
)

func WatchDocuments() {
	w := watcher.New()

	w.FilterOps(watcher.Write)


	go func() {
		for {
			select {
			case event := <-w.Event:	
				fmt.Println(event) // Print the event's info.
				app_builder.BuildApp()
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	if err := w.Add("./document"); err != nil {
		log.Fatalln(err)
	}

	// // Print a list of all of the files and folders currently
	// // being watched and their paths.
	// for path, f := range w.WatchedFiles() {
	// 	fmt.Printf("%s: %s\n", path, f.Name())
	// }

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
