package fetch

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func getEncoding(reader io.Reader) (determineEncoding encoding.Encoding) {
	peek, err := bufio.NewReader(reader).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	determineEncoding, _, _ = charset.DetermineEncoding(peek, "charset")
	return
}

func Fetch(url string) ([]byte,error){
	response, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("status code is not ok ,code is %d",response.StatusCode)
	}

	determineEncoding := getEncoding(response.Body)

	body := transform.NewReader(response.Body,determineEncoding.NewDecoder())

	bytes, err := ioutil.ReadAll(body)

	return bytes,err
}