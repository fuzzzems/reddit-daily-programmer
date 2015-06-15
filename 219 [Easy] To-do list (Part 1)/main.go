// [2015-06-15] Challenge #219 [Easy] To-do list (Part 1)
// http://www.reddit.com/r/dailyprogrammer/comments/39ws1x/20150615_challenge_218_easy_todo_list_part_1/

package main

import (
	"fmt"
)

// Struct because this allows for expansion of type of data we can associate with a note.
type Note struct {
	content string
}

type Notebook []Note

var notebook = Notebook{}

func addNote(s string) {
	notebook = append(notebook, Note{content: s})
}

func viewNotes() {
	for i := range notebook {
		fmt.Println(notebook[i].content)
	}
}

func deleteNote(s string) {
	for i := range notebook {
		if s == notebook[i].content {
			notebook = append(notebook[:i], notebook[i+1:]...)
			break
		}

		if i == len(notebook)-1 {
			fmt.Printf("Note: \"%s\" not found", s)
		}

	}
}

func main() {

	addNote("Take a shower")
	addNote("Go to work")
	viewNotes()
	addNote("Buy a new phone")
	viewNotes()
	deleteNote("Take a shower")
	viewNotes()

}
