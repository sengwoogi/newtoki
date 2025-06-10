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
