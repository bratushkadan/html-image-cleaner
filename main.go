package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func processHTMLFile(inputFile string) (string, error, bool) {
	// Чтение содержимого файла
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return "", err, false
	}

	// Создание нового документа из HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		return "", err, false
	}

	// Найти все img теги и удалить атрибут crossorigin
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		s.RemoveAttr("crossorigin")
	})

	// Получить измененный HTML
	modifiedHTML, err := doc.Html()
	if err != nil {
		return "", err, false
	}

	return modifiedHTML, nil, string(data) != modifiedHTML
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	files, err := os.ReadDir(cwd)
	if err != nil {
		log.Fatalf("Failed to read current working directory: %v", err)
	}

	htmlFiles := make([]string, 0)
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".html") {
			htmlFiles = append(htmlFiles, file.Name())
		}
	}

	if len(htmlFiles) == 0 {
		log.Fatalf("No .html files in current directory")
	}

	for _, filename := range htmlFiles {
		modifiedHTML, err, transformed := processHTMLFile(filename)
		if err != nil {
			log.Fatalf("Error processing HTML file: %v", err)
		}

		if !transformed {
			continue
		}
		err = os.WriteFile(filename, []byte(modifiedHTML), 0644)
		if err != nil {
			log.Fatalf("Error writing the modified HTML to file: %v", err)
		}
		fmt.Printf("Modified HTML saved to %s\n", filename)
	}
}
