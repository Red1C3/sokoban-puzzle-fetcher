package main

import (
	"sokoban-puzzle-fetcher/fetcher"
)

func main(){
	url:="https://sokoban.info/?1_2"
	fetcher.Fetch(url)
}