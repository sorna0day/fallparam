package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: jsparam -d <website>")
		return
	}

	switch os.Args[1] {
	case "-h":
		fmt.Println("Usage: jsparam -d <website>")
		fmt.Println("-h\t\tShow this help message")
		fmt.Println("-v\t\tShow version")
	case "-r":
		deleteTool()
	case "-up":
		updateTool()
	default:
		fmt.Println("Invalid command")
	}
}

func deleteTool() {
	// Get the operating system
	os := runtime.GOOS

	// Delete the tool based on the operating system
	switch os {
	case "windows":
		err := os.Remove("C:\\path\\to\\jsparam.exe")
		if err != nil {
			log.Fatal(err)
		}
	case "linux":
		err := os.Remove("/usr/local/bin/jsparam")
		if err != nil {
			log.Fatal(err)
		}
	case "darwin":
		err := os.Remove("/usr/local/bin/jsparam")
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Unsupported operating system")
	}
}

func updateTool() {
	// Download and install the latest version of the tool from GitHub
	fmt.Println("Updating jsparam...")
	// Add code to download and install the latest version of the tool from GitHub
	fmt.Println("jsparam updated successfully!")
}
