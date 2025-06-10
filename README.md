# Newtoki Novel Downloader

This program, `newtoki.go`, is a command-line tool designed to download novel content from websites structured similarly to Newtoki. It fetches chapters sequentially based on IDs found in the initial novel page and saves the extracted text content to a local file.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Go (Golang) needs to be installed on your system. You can download it from the official Go download page: [https://golang.org/dl/](https://golang.org/dl/)

### Running the Program

To run the program directly, use the following command:

```bash
go run newtoki.go
```

## Usage

1.  The program will first prompt you to "Enter filename:". This will be the base name for the output file (e.g., if you enter "mynovel", the output will be "mynovel.txt").
2.  Next, it will prompt "Enter headers:".
3.  You need to paste a full HTTP GET request, including the request line and necessary headers.
4.  **Crucially, the `Host` header must be present.** Other headers like `User-Agent`, `Cookie`, etc., can be included as needed.
5.  Here is an example of the header input:

    ```
    GET /novel/12345 HTTP/1.1
    Host: yourtargetdomain.com
    User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36
    Cookie: your_cookie_if_needed
    ```
6.  The program will then download the novel content and save it to `[filename].txt`, showing progress along the way.

## Built With

* [Go (Golang)](https://golang.org/) - The programming language used
* [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery) - Go library for HTML parsing

## Authors

This script was created to download novel content.

## License

This project is unlicensed.

---

# Newtoki 소설 다운로더 (Korean)

이 프로그램(`newtoki.go`)은 Newtoki와 유사한 구조의 웹사이트에서 소설 콘텐츠를 다운로드하도록 설계된 명령줄 도구입니다. 초기 소설 페이지에 있는 ID를 기반으로 순차적으로 챕터를 가져오고 추출된 텍스트 콘텐츠를 로컬 파일에 저장합니다.

## 시작하기 (Korean)

이 지침은 개발 및 테스트 목적으로 로컬 시스템에서 프로젝트를 복사하고 실행하는 방법을 안내합니다.

### 사전 필요 조건 (Korean)

시스템에 Go (Golang)가 설치되어 있어야 합니다. 공식 Go 다운로드 페이지([https://golang.org/dl/](https://golang.org/dl/))에서 다운로드할 수 있습니다.

### 프로그램 실행 (Korean)

프로그램을 직접 실행하려면 다음 명령을 사용하십시오:

```bash
go run newtoki.go
```

## 사용법 (Korean)

1.  프로그램은 먼저 "Enter filename:"이라는 메시지를 표시합니다. 이것은 출력 파일의 기본 이름이 됩니다 (예: "mynovel"을 입력하면 출력은 "mynovel.txt"가 됩니다).
2.  다음으로 "Enter headers:"라는 메시지가 표시됩니다.
3.  요청 라인과 필요한 헤더를 포함한 전체 HTTP GET 요청을 붙여넣어야 합니다.
4.  **중요: `Host` 헤더가 반드시 있어야 합니다.** `User-Agent`, `Cookie` 등과 같은 다른 헤더는 필요에 따라 포함할 수 있습니다.
5.  헤더 입력의 예는 다음과 같습니다:

    ```
    GET /novel/12345 HTTP/1.1
    Host: yourtargetdomain.com
    User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36
    Cookie: your_cookie_if_needed
    ```
    ![image](https://github.com/user-attachments/assets/1100e4b1-194f-4392-bcb1-c8232dbd897d)
    POST /novel/18494822 HTTP/3
    Host: booktoki468.com ... 붙여넣기

    ![image](https://github.com/user-attachments/assets/204ef835-72d9-4129-b30a-6d8cec043ac7)
    :authority
    booktoki468.com .. 붙여넣기

    
7.  그러면 프로그램이 소설 콘텐츠를 다운로드하고 `[filename].txt`에 저장하며 진행 상황을 보여줍니다.

## 만들어진 환경 (Korean)

* [Go (Golang)](https://golang.org/) - 사용된 프로그래밍 언어
* [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery) - HTML 구문 분석을 위한 Go 라이브러리

## 작성자 (Korean)

이 스크립트는 소설 콘텐츠를 다운로드하기 위해 만들어졌습니다.
