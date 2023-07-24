package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: fallparam -d <website>")
		return
	}

	switch os.Args[1] {
	case "-h":
		fmt.Println("Usage: fallparam -d <website>")
		fmt.Println("-h\t\tShow this help message")
		fmt.Println("-v\t\tShow version")
		fmt.Println("update-\t\tUpdate to the latest version")
		fmt.Println("-r\t\tRemove fallparam from the server")
	case "-v":
		fmt.Println("fallparam version", version)
	case "update-":
		fmt.Println("Updating fallparam...")
		cmd := exec.Command("git", "pull")
		cmd.Dir = "/root/go/bin/fallparam "
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("fallparam updated successfully")
	case "-r":
		fmt.Println("Removing fallparam...")
		cmd := exec.Command("rm", "-rf", "/root/go/bin/fallparam ")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("fallparam removed successfully")
	case "-d":
		if len(os.Args) < 3 {
			fmt.Println("Usage: fallparam -d <website>")
			return
		}

		website := os.Args[2]

		response, err := http.Get(website)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		varNameRegex := regexp.MustCompile(`\bvar\s+(\w+)\s+=`)
		variables := varNameRegex.FindAllStringSubmatch(string(body), -1)

		jsonRegex := regexp.MustCompile(`\bJSON\.parse\((.*?)\)`)
		jsonObjects := jsonRegex.FindAllStringSubmatch(string(body), -1)

		fmt.Println("JavaScript variables names:")
		for _, match := range variables {
			fmt.Println(match[1])
		}

		fmt.Println("JSON object in JavaScript files:")
		for _, match := range jsonObjects {
			fmt.Println(match[1])
		}
	default:
		fmt.Println("Usage: fallparam -d <website>")
	}

	if runtime.GOOS == "linux" {
		fmt.Println("This is a Linux system")
	} else if runtime.GOOS == "windows" {
		fmt.Println("This is a Windows system")
	} else if runtime.GOOS == "darwin" {
		fmt.Println("This is a macOS system")
	}

	if len(os.Args) > 1 && os.Args[1] == "fallparam" {
		fmt.Println("fallparam")
	}
}
