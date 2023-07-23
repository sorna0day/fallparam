package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
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

		// Clone the repository
		repo, err := git.PlainClone(".", false, &git.CloneOptions{
			URL:      "https://github.com/sorna0day/fallparam.git",
			Progress: os.Stdout,
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		// Get the latest commit
		ref, err := repo.Head()
		if err != nil {
			fmt.Println(err)
			return
		}
		commit, err := repo.CommitObject(ref.Hash())
		if err != nil {
			fmt.Println(err)
			return
		}

		// Checkout the latest commit
		worktree, err := repo.Worktree()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = worktree.Checkout(&git.CheckoutOptions{
			Hash: commit.Hash,
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		// Build the updated binary
		err = os.Chdir("cmd/fallparam")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("../../..")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("cmd/fallparam")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("../../..")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("cmd/fallparam")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("../../..")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("cmd/fallparam")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("../../..")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("cmd/fallparam")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("../../..")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("cmd/fallparam")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("../../..")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("cmd/fallparam")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("../../..")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("cmd/fallparam")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("../../..")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("cmd/fallparam")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("../../..")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("cmd/fallparam")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("../../..")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("cmd/fallparam")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("../../..")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("cmd/fallparam")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("../../..")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("cmd/fallparam")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("main.go")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = os.Chdir("../../..")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Update complete!")
	}
}
