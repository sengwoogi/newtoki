package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var text2 = ""
var matches [][]string

func save(fileName string, randnumber string, headers map[string]string) {
	req, err := http.NewRequest("GET", randnumber, nil)
	if err != nil {
		panic(err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received non-200 response code: %d\n", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("Error creating document:", err)
		return
	}
	text2 = ""
	doc.Find("#novel_content").Each(func(i int, s *goquery.Selection) {
		s.Find("p").Each(func(j int, p *goquery.Selection) {
			text := p.Text()
			text2 = text2 + strings.TrimSpace(text) + "\n\n"
		})
	})
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString("\n")
	if _, err := f.WriteString(text2); err != nil {
		log.Fatal(err)
	}
	if err != nil {
		panic(err)
	}
}

func main() {
	var fileName string
	fmt.Print("Enter filename: ")
	fmt.Scanln(&fileName)

	allowedHeaders := map[string]bool{
		"Accept":          true,
		"Accept-Language": true,
		"Connection":      true,
		"User-Agent":      true,
		"Cookie":          true,
		"Host":            true,
	}

	fmt.Println("Enter headers")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	getLine := scanner.Text()
	getParts := strings.Fields(getLine)
	if len(getParts) < 2 {
		fmt.Println("Invalid GET request format")
		return
	}
	getPath := getParts[1]

	headers := make(map[string]string)
	for {
		scanner.Scan()
		headerLine := scanner.Text()
		if strings.TrimSpace(headerLine) == "" {
			break
		}
		parts := strings.SplitN(headerLine, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if allowedHeaders[key] {
				headers[key] = value
			} else {
				fmt.Printf("Disallowed header: %s\n", key)
			}
		}
	}

	host := headers["Host"]
	if host == "" {
		fmt.Println("Host header is required")
		return
	}

	randnumber := "https://" + host + getPath

	u, err := url.Parse(randnumber)
	if err != nil {
		panic(err)
	}

	lastSlash := strings.LastIndex(u.Path, "/")
	baseURL := u.Scheme + "://" + u.Host + u.Path[:lastSlash+1]

	req, err := http.NewRequest("GET", randnumber, nil)
	if err != nil {
		panic(err)
	}


	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`<option value="(\d+)"`)
	matches = re.FindAllStringSubmatch(string(body), -1)

	for i := len(matches) - 1; i >= 0; i-- {
		novelURL := baseURL + matches[i][1]
		fmt.Printf("Progress: %d/%d (%.2f%%)\n", len(matches)-i, len(matches), float64(len(matches)-i)/float64(len(matches))*100)
		time.Sleep(10 * time.Second)
		save(fileName+".txt", novelURL, headers)
	}
}
