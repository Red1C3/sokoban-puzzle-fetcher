package fetcher

import (
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"regexp"
	"strconv"
	"strings"
)

// <option value='1' selected="selected">Original &amp; Extra &nbsp; (90)</option>
const (
	boardRegexStr      = `="[.#$*@+x _!]+"`
	collectionRegexStr = `<option value='[0-9]+'.*\([0-9]+\)`
	quotesRegexStr=`'[0-9]+'`
	numRegexStr=`[0-9]+`
	parRegexStr=`\([0-9]+\)`
)

const (
	BASEURL = "https://sokoban.info/"
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

func Fetch(url string, f map[rune]rune) (string, string) {
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
	boardStr = strings.Map(mapper, boardStr)

	fileURL, err := url2.Parse(url)
	if err != nil {
		log.Fatalf("Failed to parse %s", url)
	}
	query := fileURL.RawQuery

	return boardStr, query
}

func mapper(r rune) rune {
	if ru, ok := format[r]; ok {
		return ru
	}
	return -1
}

func FetchCollections() []int {
	resp, err := client.Get(BASEURL)
	if err != nil {
		log.Fatalf("Failed to get a response from %s", BASEURL)
	}
	defer resp.Body.Close()

	var htmlFile strings.Builder
	_, err = io.Copy(&htmlFile, resp.Body)
	if err != nil {
		log.Fatalf("Failed to copy response body to string")
	}

	collectionRegex, err := regexp.Compile(collectionRegexStr)
	if err != nil {
		log.Fatalf("Failed to compile regex %s", collectionRegexStr)
	}
	collectionStrings := collectionRegex.FindAllString(htmlFile.String(), -1)
	quotesRegex,err:=regexp.Compile(quotesRegexStr)
	if err!=nil{
		log.Fatalf("Failed to compile regex %s",quotesRegexStr)
	}
	numRegex,err:=regexp.Compile(numRegexStr)
	if err!=nil{
		log.Fatalf("Failed to compile regex %s",numRegexStr)
	}
	parRegex,err:=regexp.Compile(parRegexStr)
	if err!=nil{
		log.Fatalf("Failed to compile regex %s",parRegexStr)
	}
	collections := make([]int, 132)
	for _, s := range collectionStrings {
		quotedIndex := quotesRegex.FindString(s)
		strIndex := numRegex.FindString(quotedIndex)
		index, err := strconv.ParseInt(strIndex, 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse %s to int", strIndex)
		}
		parVal:=parRegex.FindString(s)
		strVal:=numRegex.FindString(parVal)
		val, err := strconv.ParseInt(strVal, 10, 32)
		if err != nil {
			log.Fatalf("Failed to parse %s to int", strVal)
		}
		collections[index-1] = int(val)
	}
	return collections
}
