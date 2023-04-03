package main

import (
	"context"
	"flag"
	"log"
	"os/exec"
	"sync"
	"time"

	"the-watcher/utils"

	"github.com/fsnotify/fsnotify"
)

var mu sync.Mutex

func buildProject(ctx context.Context, command string) {

	mu.Lock()
	defer mu.Unlock()

	log.Print("Starting build process")

	// Create a new context with a timeout of 8 seconds
	ctx, cancel := context.WithTimeout(ctx, 8*time.Second)
	defer cancel()

	// Run the build command
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	_, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("failed to run build command %v", err)
	}

	log.Print("Build process completed")
}

func main() {

	// const TEST_FOLDER_PATH string = "./test_folder"

	dirToBeObervedPtr := flag.String("dir", "", "Directory to observe changes")
	buildCommandPtr := flag.String("command", "", "Command to run each time changes are observed in a directory")

	flag.Parse()

	if *dirToBeObervedPtr == "" {
		log.Fatalln("[Insufficient args] please provide directory to watch with -dir flag ")
	}

	if *buildCommandPtr == "" {
		log.Fatalln("[Insufficient args] please provide build command with -command flag ")
	}

	ctx := context.Background()
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Add(*dirToBeObervedPtr)

	if err != nil {
		log.Fatalf("Unable to watch %s - Error: %v", *dirToBeObervedPtr, &err)
	}

	log.Print("Listening for change events")

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				log.Print("[WARNING] problem with changes stream")
			}

			lastChangedSources := map[string]struct{}{event.Name: {}}

			if !utils.ShouldRebuild(event.Name, event.Op) {
				continue
			}

			buildProject(ctx, *buildCommandPtr)

			log.Printf("{lastChangedSources}: %v", lastChangedSources)

			log.Println("event:", event)

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}

}
