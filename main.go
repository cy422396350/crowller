package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func getEncoding(reader io.Reader) (determineEncoding encoding.Encoding) {
	peek, err := bufio.NewReader(reader).Peek(1024)
	if err != nil {
		panic(err)
	}
	determineEncoding, _, _ = charset.DetermineEncoding(peek, "charset")
	return
}

func main() {
	response, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println("response status code",response.StatusCode)
	}

	determineEncoding := getEncoding(response.Body)

	body := transform.NewReader(response.Body,determineEncoding.NewDecoder())

	bytes, err := ioutil.ReadAll(body)

	fmt.Println(string(bytes))
}