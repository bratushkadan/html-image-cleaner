package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func processHTMLFile(inputFile string) (string, error) {
	// Чтение содержимого файла
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return "", err
	}

	// Создание нового документа из HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		return "", err
	}

	// Найти все img теги и удалить атрибут crossorigin
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		s.RemoveAttr("crossorigin")
	})

	// Получить измененный HTML
	modifiedHTML, err := doc.Html()
	if err != nil {
		return "", err
	}

	return modifiedHTML, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: <input_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	// Генерация имени выходного файла
	// ext := filepath.Ext(inputFile)
	// baseName := strings.TrimSuffix(inputFile, ext)
	// outputFile := baseName + ".edited" + ext
	outputFile := inputFile

	modifiedHTML, err := processHTMLFile(inputFile)
	if err != nil {
		log.Fatalf("Error processing HTML file: %v", err)
	}

	// Запись измененного HTML в новый файл
	err = ioutil.WriteFile(outputFile, []byte(modifiedHTML), 0644)
	if err != nil {
		log.Fatalf("Error writing the modified HTML to file: %v", err)
	}

	fmt.Printf("Modified HTML saved to %s\n", outputFile)
}
