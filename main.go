package main

import (
	"flag"
	"sokoban-puzzle-fetcher/fetcher"
	"sokoban-puzzle-fetcher/parser"
)

const (
	BASEURL = "https://sokoban.info/?"
)

var (
	puzzle = "1_2"
	name   = ""
	dst    = "."
	all    = false
)

func main() {
	parseFlags()
	format := make(map[rune]rune)
	format[fetcher.PLAYER] = fetcher.PLAYERCHAR
	format[fetcher.BLANK] = fetcher.BLANKCHAR
	format[fetcher.OBSTACLE] = fetcher.OBSTACLECHAR
	format[fetcher.OUTSIDE] = fetcher.OBSTACLECHAR
	format[fetcher.LINEBREAK] = fetcher.LINEBREAKCHAR
	format[fetcher.GOAL] = fetcher.GOALCHAR
	format[fetcher.BOX] = fetcher.BOXCHAR
	format[fetcher.SKIP] = fetcher.BLANK
	format[fetcher.BOXONGOAL] = fetcher.BOXONGOALCHAR
	format[fetcher.PLAYERONGOAL] = fetcher.PLAYERONGOALCHAR
	if all {
		url := BASEURL + puzzle
		str, queryName := fetcher.Fetch(url, format)
		if len(name) != 0 {
			parser.Parse(str, name, dst)
		} else {
			parser.Parse(str, queryName, dst)
		}
	} else {
		fetcher.FetchCollections()
	}
}

func parseFlags() {
	blankPtr := flag.String("blank", string(fetcher.BLANKCHAR), "Blank character as in output file")
	boxPtr := flag.String("box", string(fetcher.BOXCHAR), "Box character as in output file")
	obstaclePtr := flag.String("obs", string(fetcher.OBSTACLECHAR), "Obstacle character as in output file")
	playerPtr := flag.String("player", string(fetcher.PLAYERCHAR), "Player character as in output file")
	goalPtr := flag.String("goal", string(fetcher.GOALCHAR), "Goal character as in output file")
	boxOnGoalPtr := flag.String("bog", string(fetcher.BOXONGOAL), "Box on Goal character as in output file")
	playerOnGoalPtr := flag.String("pog", string(fetcher.PLAYERONGOAL), "Player on Goal as in output file")

	flag.StringVar(&puzzle, "puzzle", puzzle, "Puzzle as in <group>_<puzzle>")
	flag.StringVar(&name, "name", name, "Output file name")
	flag.StringVar(&dst, "dst", dst, "File destination")
	flag.BoolVar(&all, "all", all, "Fetch all puzzles")

	flag.Parse()

	fetcher.BLANKCHAR = []rune(*blankPtr)[0]
	fetcher.BOXCHAR = []rune(*boxPtr)[0]
	fetcher.OBSTACLECHAR = []rune(*obstaclePtr)[0]
	fetcher.PLAYERCHAR = []rune(*playerPtr)[0]
	fetcher.GOALCHAR = []rune(*goalPtr)[0]
	fetcher.BOXONGOALCHAR = []rune(*boxOnGoalPtr)[0]
	fetcher.PLAYERONGOALCHAR = []rune(*playerOnGoalPtr)[0]
}
