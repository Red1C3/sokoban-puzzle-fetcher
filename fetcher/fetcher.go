package fetcher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const (
	boardRegexStr = `var Board\s*="[.#$*@+x _!]+"`
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

func Fetch(url string, f map[rune]rune) string {
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
	fmt.Println(boardStr)

	return boardStr
}

func mapper(r rune) rune {
	if ru, ok := format[r]; ok {
		return ru
	}
	return -1
}
