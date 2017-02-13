package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
	"os/exec"
)

var version string
var installlocation string
var path string
var filename string

func init() {
	// Set up command-line flags.
	log.Print("Starting init function")
	flag.StringVar(&path, "f", "", "go source code file")
	flag.Parse()
	// Validate flags.
	if path == "" {
		log.Fatal("No .go file passed.")
	}
	os.Mkdir("tmp", 0666)
	log.Print("tmp made")
	os.Chdir("tmp")
	log.Print("change directory to tmp")


	logfile := fmt.Sprintf("buildgo_%s.log", time.Now().Format("2006-01-02_150405"))


	f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")

}

func main() {

	out, err := exec.Command("docker", "pull golang").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Ryan %s\n", out)




}
