# sokoban-puzzle-fetcher
A tool that fetches and formats sokoban puzzles from https://sokoban.info/

## Features
- Files are output as JSON 2D arrays
- Get all the puzzles from the site in a single command
- Specify the symbols of the output

## Usage
Check out ```./sokoban-puzzle-fetcher -h```

## Example
Inputting puzzle 5_16, the output:
```
[
["🟥","🟥","🟥","🟥","🟥","🟥","🟥","🟥","🟥"],
["🟥","🟥","🔲","🔲","🟥","🟥","🟥","🟥","🟥"],
["🟥","🟥","🔲","🔲","🔲","🟩","🔲","🔲","🟥"],
["🟥","🟥","🟩","🟥","🔲","🟥","🔲","🔲","🟥"],
["🟥","🟥","🔲","🔲","😿","🔲","🔲","🟥","🟥"],
["🟥","🔲","🔲","🟥","⭕","🟥","🟩","🟥","🟥"],
["🟥","🔲","🔲","🟩","⭕","🔲","🔲","🟥","🟥"],
["🟥","🟥","🟥","🔲","⭕","🔲","🟥","🟥","🟥"],
["🟥","🟥","🟥","🟥","🟥","🟥","🟥","🟥","🟥"]
]

```

## Building
```go build```

## License
[UNLICENSE](./UNLICENSE)