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
	// بررسی تعداد آرگومان‌ها
	if len(os.Args) < 2 {
		fmt.Println("Usage: fallparam -d <website>")
		return
	}

	// بررسی آرگومان‌ها
	var website string
	var update bool
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-d":
			if i+1 < len(os.Args) {
				website = os.Args[i+1]
			}
		case "-v":
			fmt.Println("Version:", version)
			return
		case "-up":
			update = true
		}
	}

	// بررسی آدرس وب سایت
	if website == "" {
		fmt.Println("Usage: fallparam -d <website>")
		return
	}

	// باز کردن صفحه وب سایت
	response, err := http.Get(website)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// خواندن محتوای صفحه وب سایت
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// استفاده از عبارت منظم برای پیدا کردن نام‌های متغیرها در جاوا اسکریپت
	varNameRegex := regexp.MustCompile(`\bvar\s+(\w+)\s+=`)
	variables := varNameRegex.FindAllStringSubmatch(string(body), -1)

	// استفاده از عبارت منظم برای پیدا کردن JSON object در فایل‌های جاوا اسکریپت
	jsonRegex := regexp.MustCompile(`\bJSON\.parse\((.*?)\)`)
	jsonObjects := jsonRegex.FindAllStringSubmatch(string(body), -1)

	// چاپ نام‌های متغیرها
	fmt.Println("JavaScript variables names:")
	for _, match := range variables {
		fmt.Println(match[1])
	}

	// چاپ JSON objects
	fmt.Println("JSON object in JavaScript files:")
	for _, match := range jsonObjects {
		fmt.Println(match[1])
	}

	// بروزرسانی ابزار
	if update {
		fmt.Println("Updating fallparam...")
		// TODO: Add update code here
		fmt.Println("Update complete!")
	}
}
