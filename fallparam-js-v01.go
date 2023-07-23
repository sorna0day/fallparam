package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
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
		// TODO: Implement update command
	case "-r":
		fmt.Println("Removing fallparam...")
		// TODO: Implement remove command
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
}
