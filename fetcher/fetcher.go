package fetcher

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const(
	boardRegexStr=`var Board\s*="[.#$*@+x _!]+"`
)

var (
	client=http.Client{}
)

func Fetch(url string) string{
	resp,err:=client.Get(url)
	if err!=nil{
		log.Fatalf("Failed to get a response from %s",url)
	}
	defer resp.Body.Close()

	var htmlFile strings.Builder
	_,err=io.Copy(&htmlFile,resp.Body)
	if err!=nil{
		log.Fatalf("Failed to copy response body to string")
	}

	boardRegex,err:=regexp.Compile(boardRegexStr)
	if err!=nil{
		log.Fatalf("Failed to compile regex %s",boardRegexStr)
	}
	boardStr:=boardRegex.FindString(htmlFile.String())
	log.Println(boardStr)
	return boardStr
}