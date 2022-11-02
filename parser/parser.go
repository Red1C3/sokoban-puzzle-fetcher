package parser

import (
	"encoding/json"
	"log"
	"os"
	"sokoban-puzzle-fetcher/fetcher"
)

func Parse(s string,fileName string,dst string){
	puzzle:=make([][]string,1)
	i:=0
	for _,ch:=range s{
		switch ch{
		case fetcher.LINEBREAKCHAR:
			i+=1
			puzzle=append(puzzle,make([]string,0))
		default:
			puzzle[i]=append(puzzle[i],string(ch))
		}
	}
	puzzle=puzzle[:len(puzzle)-1]
	file,err:=os.Create(dst+"/"+fileName+".json")
	if err!=nil{
		log.Fatalf("Failed to create file %s",fileName)
	}
	defer file.Close()
	enc:=json.NewEncoder(file)
	err=enc.Encode(puzzle)
	if err!=nil{
		log.Fatal("Failed to encode puzzle")
	}
}