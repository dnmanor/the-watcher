# WATCHER

This is a simple Go program that listens for changes in a specified directory and runs a command each time a change is detected.

## Prerequisites

Go installed on your machine

## Usage

- Clone this repository to your local machine
- Open a terminal and navigate to the directory where you cloned the repository
- Run the command go run main.go -dir <directory to observe changes> -command "<command to run each time changes are observed>" replacing <directory to observe changes> with the path to the directory you want to observe and <command to run each time changes are observed> with the command you want to run each time a change is detected in the directory.

## Example

go run main.go -dir ./test-folder -command "go build"

In the above example, the program will observe changes in the directory test-folder and run the command go build each time a change is detected.
