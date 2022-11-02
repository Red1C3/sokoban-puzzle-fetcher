package main

import (
	"sokoban-puzzle-fetcher/fetcher"
)

func main() {
	format := make(map[rune]rune)
	format[fetcher.PLAYER] = '@' //TODO
	url := "https://sokoban.info/?1_2"
	fetcher.Fetch(url, format)
}
