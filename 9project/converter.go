package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/russross/blackfriday/v2"
)

func convertMarkdownToHTML(mdContent string) []byte {
	htmlContent := blackfriday.Run([]byte(mdContent))
	return htmlContent
}

func readMarkdownFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}

func writeHTMLToFile(filePath string, htmlContent []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(htmlContent)
	return err
}

func main() {
	mdFilePath := "example.md" 
	mdContent, err := readMarkdownFile(mdFilePath)
	if err != nil {
		log.Fatalf("Error reading markdown file: %v", err)
	}

	htmlContent := convertMarkdownToHTML(mdContent)

	htmlFilePath := "output.html" 
	err = writeHTMLToFile(htmlFilePath, htmlContent)
	if err != nil {
		log.Fatalf("Error writing HTML file: %v", err)
	}

	fmt.Println("Conversion complete. HTML file saved as:", htmlFilePath)
}
