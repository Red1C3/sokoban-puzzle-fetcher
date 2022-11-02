package main

import (
	"sokoban-puzzle-fetcher/fetcher"
	"sokoban-puzzle-fetcher/parser"
)

func main() {
	format := make(map[rune]rune)
	format[fetcher.PLAYER] = fetcher.PLAYERCHAR //TODO support custom formats
	format[fetcher.BLANK]=fetcher.BLANKCHAR
	format[fetcher.OBSTACLE]=fetcher.OBSTACLECHAR
	format[fetcher.OUTSIDE]=fetcher.OBSTACLECHAR
	format[fetcher.LINEBREAK]=fetcher.LINEBREAKCHAR
	format[fetcher.GOAL]=fetcher.GOALCHAR
	format[fetcher.BOX]=fetcher.BOXCHAR
	format[fetcher.SKIP]=fetcher.BLANK
	format[fetcher.BOXONGOAL]=fetcher.BOXONGOALCHAR
	format[fetcher.PLAYERONGOAL]=fetcher.PLAYERONGOALCHAR
	url := "https://sokoban.info/?1_2"
	str:=fetcher.Fetch(url, format) //TODO get file name
	parser.Parse(str,"meow.json",".")
}
