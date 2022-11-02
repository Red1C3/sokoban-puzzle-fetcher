package fetcher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"regexp"
	"strings"
)

const (
	boardRegexStr = `="[.#$*@+x _!]+"`
)

const (
	PLAYER       = '@'
	OBSTACLE     = '#'
	OUTSIDE      = 'x'
	LINEBREAK    = '!'
	GOAL         = '.'
	BLANK        = ' '
	BOX          = '$'
	SKIP         = '_' //NOT SURE HOW IT IS DIFFERENT FROM OUTSIDE
	BOXONGOAL    = '*'
	PLAYERONGOAL = '+'
)

var (
	client = http.Client{}
	format map[rune]rune
)

func Fetch(url string, f map[rune]rune) (string,string) {
	format = f
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed to get a response from %s", url)
	}
	defer resp.Body.Close()

	var htmlFile strings.Builder
	_, err = io.Copy(&htmlFile, resp.Body)
	if err != nil {
		log.Fatalf("Failed to copy response body to string")
	}

	boardRegex, err := regexp.Compile(boardRegexStr)
	if err != nil {
		log.Fatalf("Failed to compile regex %s", boardRegexStr)
	}
	boardStr := boardRegex.FindString(htmlFile.String())
	fmt.Println(boardStr)
	boardStr = strings.Map(mapper, boardStr)
	fmt.Println(boardStr)

	fileURL,err:=url2.Parse(url)
	if err!=nil{
		log.Fatalf("Failed to parse %s",url)
	}
	query:=fileURL.RawQuery

	return boardStr,query
}

func mapper(r rune) rune {
	if ru, ok := format[r]; ok {
		return ru
	}
	return -1
}
