package main

import (
	 "fmt"
	 "os"
)


import  "github.com/nsf/termbox-go"
import  "github.com/mattn/go-runewidth"

// change parameters for a struct or object or somethings
func print_message(col int, row int, fg termbox.Attribute, bg termbox.Attribute, message string){
	for _, ch := range message {
		termbox.SetCell(col, row, ch, fg, bg)
		col += runewidth.RuneWidth(ch)
	}
}

func run_editor(){
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	print_message(25, 11, termbox.ColorDefault, termbox.ColorDefault, "EM - Marina's Text Editor")
	termbox.Flush()
	termbox.PollEvent()
	termbox.Close()
}

func main(){
	run_editor()
}

