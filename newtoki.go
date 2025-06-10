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
	// Request 객체 생성
	req, err := http.NewRequest("GET", randnumber, nil)
	if err != nil {
		panic(err)
	}

	// 전달받은 헤더만 추가
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Client 객체에서 Request 실행
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
	// 결과 출력
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
	// 파일에 텍스트 이어쓰기
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
	var waitTime int = 10 // 기본값 10초로 설정
	fmt.Print("Enter filename: ")
	fmt.Scanln(&fileName)
	fmt.Print("Enter wait time (seconds) [default: 10]: ")
	waitTimeInput, _ := fmt.Scanln(&waitTime)
	if waitTimeInput == 0 { // 사용자가 입력하지 않은 경우
		waitTime = 10
	}

	// 허용할 헤더 목록 정의
	allowedHeaders := map[string]bool{
		"accept":          true,
		"accept-language": true,
		"connection":      true,
		"user-agent":      true,
		"cookie":          true,
		"host":            true,
	}

	fmt.Println("Enter headers")
	scanner := bufio.NewScanner(os.Stdin)

	// 첫 번째 줄이 :authority 형식인지 확인
	scanner.Scan()
	firstLine := strings.TrimSpace(scanner.Text())

	var getPath string
	headers := make(map[string]string)

	if strings.HasPrefix(strings.ToLower(firstLine), ":") {
		// :authority 형식으로 처리
		key := strings.ToLower(firstLine)
		scanner.Scan()
		value := strings.TrimSpace(scanner.Text())

		// :authority를 Host로 변환
		if key == ":authority" {
			headers["host"] = value
		}

		// 나머지 헤더 처리
		for {
			scanner.Scan()
			headerLine := scanner.Text()
			if strings.TrimSpace(headerLine) == "" {
				break
			}

			key := strings.ToLower(strings.TrimSpace(headerLine))
			scanner.Scan()
			value := strings.TrimSpace(scanner.Text())

			// :path를 GET 요청으로 변환
			if key == ":path" {
				getPath = value
				continue
			}

			if allowedHeaders[key] {
				headers[key] = value
			} else {
				fmt.Printf("Disallowed header: %s\n", key)
			}
		}
	} else {
		// 기존 GET 형식으로 처리
		getParts := strings.Fields(firstLine)
		if len(getParts) < 2 {
			fmt.Println("Invalid GET request format")
			return
		}
		getPath = getParts[1]

		for {
			scanner.Scan()
			headerLine := scanner.Text()
			if strings.TrimSpace(headerLine) == "" {
				break
			}

			parts := strings.SplitN(headerLine, ":", 2)
			if len(parts) == 2 {
				key := strings.ToLower(strings.TrimSpace(parts[0]))
				value := strings.TrimSpace(parts[1])

				if allowedHeaders[key] {
					headers[key] = value
				} else {
					fmt.Printf("Disallowed header: %s\n", key)
				}
			}
		}
	}

	// Host 헤더에서 도메인 추출
	host := headers["host"]
	if host == "" {
		fmt.Println("Host header is required")
		return
	}

	// URL 생성
	randnumber := "https://" + host + getPath

	u, err := url.Parse(randnumber)
	if err != nil {
		panic(err)
	}
	// 마지막 '/' 위치까지 baseURL 추출
	lastSlash := strings.LastIndex(u.Path, "/")
	baseURL := u.Scheme + "://" + u.Host + u.Path[:lastSlash+1]

	req, err := http.NewRequest("GET", randnumber, nil)
	if err != nil {
		panic(err)
	}

	// 입력받은 헤더를 request에 바로 추가
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Client 객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 결과 출력
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`<option value="(\d+)"`)
	matches = re.FindAllStringSubmatch(string(body), -1)

	for i := len(matches) - 1; i >= 0; i-- {
		novelURL := baseURL + matches[i][1]
		fmt.Printf("Progress: %d/%d (%.2f%%)\n", len(matches)-i, len(matches), float64(len(matches)-i)/float64(len(matches))*100)
		time.Sleep(time.Duration(waitTime) * time.Second)
		save(fileName+".txt", novelURL, headers)
	}
}
