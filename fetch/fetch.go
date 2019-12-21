package fetch

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"time"
)

func getEncoding(reader *bufio.Reader) (determineEncoding encoding.Encoding) {
	peek, err := reader.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	determineEncoding, _, _ = charset.DetermineEncoding(peek, "charset")
	return
}

var limit = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-limit
	request, e := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36)")
	if e != nil {
		panic(e)
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bufReder := bufio.NewReader(response.Body)

	determineEncoding := getEncoding(bufReder)

	body := transform.NewReader(bufReder, determineEncoding.NewDecoder())

	bytes, err := ioutil.ReadAll(body)

	return bytes, err
}
