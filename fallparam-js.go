package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
)

func main() {
    url := flag.String("d", "", "website URL")
    help := flag.Bool("h", false, "help")
    flag.Parse()

    if *help {
        flag.PrintDefaults()
        return
    }

    if *url == "" {
        fmt.Println("Please provide a website URL with the -d flag")
        return
    }

    resp, err := http.Get(*url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    re1 := regexp.MustCompile(`\bvar\s+\w+\b`)
    re2 := regexp.MustCompile(`\bconst\s+\w+\b`)
    re3 := regexp.MustCompile(`\blet\s+\w+\b`)
    re4 := regexp.MustCompile(`\bJSON\.parse\(`)

    vars := re1.FindAllString(string(body), -1)
    consts := re2.FindAllString(string(body), -1)
    lets := re3.FindAllString(string(body), -1)
    jsons := re4.FindAllString(string(body), -1)

    fmt.Println("JavaScript variables names:")
    for _, v := range vars {
        fmt.Println(v)
    }
    for _, c := range consts {
        fmt.Println(c)
    }
    for _, l := range lets {
        fmt.Println(l)
    }

    fmt.Println("JSON object in JavaScript files:")
    for _, j := range jsons {
        fmt.Println(j)
    }
}
