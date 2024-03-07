package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type News struct {
	Id    int
	Title string
	Url   string
	Theme string
}

func main() {
	var index *int = flag.Int("index", 0, "an index")
	var title *string = flag.String("title", "Unknown", "a title")
	var url *string = flag.String("url", "", "an url")
	var theme *string = flag.String("theme", "technology", "a theme of news")

	var read *bool = flag.Bool("read", false, "read file")
	var clear *bool = flag.Bool("clear", false, "clear file content")

	flag.Parse()

	if *read {
		readDataBase()
	} else if *clear {
		clearFile()
	} else {
		writeFile(*index, *title, *url, *theme)
	}
}

func readDataBase() {
	jsonText, err := os.ReadFile("data.json")
	checkErr(err)
	fmt.Println(string(jsonText))
}

func clearFile() {
	var err error = os.Truncate("data.json", 0)
	os.WriteFile("data.json", []byte("[]"), 0644)
	checkErr(err)
}

func writeFile(index int, title string, url string, theme string) {
	var all_news []News

	jsonText, err1 := os.ReadFile("data.json")
	checkErr(err1)

	var err2 error = json.Unmarshal([]byte(jsonText), &all_news)
	checkErr(err2)

	all_news = append(all_news, News{index, title, url, theme})

	result, err3 := json.Marshal(all_news)
	checkErr(err3)

	os.WriteFile("data.json", result, 0644)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
