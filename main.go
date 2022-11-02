package main

import (
	"flag"
	"sokoban-puzzle-fetcher/fetcher"
	"sokoban-puzzle-fetcher/parser"
)

const (
	BASEURL="https://sokoban.info/?"
)


func main() {
	format := make(map[rune]rune)
	format[fetcher.PLAYER] = fetcher.PLAYERCHAR
	format[fetcher.BLANK]=fetcher.BLANKCHAR
	format[fetcher.OBSTACLE]=fetcher.OBSTACLECHAR
	format[fetcher.OUTSIDE]=fetcher.OBSTACLECHAR
	format[fetcher.LINEBREAK]=fetcher.LINEBREAKCHAR
	format[fetcher.GOAL]=fetcher.GOALCHAR
	format[fetcher.BOX]=fetcher.BOXCHAR
	format[fetcher.SKIP]=fetcher.BLANK
	format[fetcher.BOXONGOAL]=fetcher.BOXONGOALCHAR
	format[fetcher.PLAYERONGOAL]=fetcher.PLAYERONGOALCHAR
	url := BASEURL+"1_2"
	str,name:=fetcher.Fetch(url, format)
	parser.Parse(str,name,".")
}

func parseFlags(){
	blankPtr:=flag.String("blank",string(fetcher.BLANKCHAR),"Blank character as in output file")
	boxPtr:=flag.String("box",string(fetcher.BOXCHAR),"Box character as in output file")
	obstaclePtr:=flag.String("obs",string(fetcher.OBSTACLECHAR),"obstacle character as in output file")
	

}