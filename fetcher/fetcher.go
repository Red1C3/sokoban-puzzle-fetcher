package fetcher

import (
	"fmt"
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
	fmt.Println(boardStr)
	boardStr = strings.Map(mapper, boardStr)
	fmt.Println(boardStr)

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

	collections := make([]int, 132)
	for _, s := range collectionStrings {
		quotedIndex, err := regexp.MatchString("'[0-9]+'", s)
		if err != nil {
			log.Fatalf("Failed to match string %s", s)
		}
		strIndex, err := regexp.MatchString("[0-9]+", quotedIndex)
		if err != nil {
			log.Fatalf("Faild to match string %s", quotedIndex)
		}
		index, err := strconv.ParseInt(strIndex, 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse %s to int", strIndex)
		}
		parVal, err := regexp.MatchString(`\([0-9]+\)`, s)
		if err != nil {
			log.Fatalf("Failed to match string %s", s)
		}
		strVal, err := regexp.MatchString("[0-9]+", parVal)
		if err != nil {
			log.Fatalf("Failed to match string %s", parVal)
		}
		val, err := strconv.ParseInt(strVal, 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse %s to int", strVal)
		}
		collections[index] = val
	}
	fmt.Println(collections)
	return collections
}
