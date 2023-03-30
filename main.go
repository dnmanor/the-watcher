package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	const TEST_FOLDER_PATH string = "./test_folder"
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Add(TEST_FOLDER_PATH)

	if err != nil {
		log.Fatalf("Unable to watch %s - Error: %v", TEST_FOLDER_PATH, &err)
	}

	log.Print("Listening for change events")

	for {
		select {
		case event := <-watcher.Events:
			log.Println("event:", event)
		}
	}

}
