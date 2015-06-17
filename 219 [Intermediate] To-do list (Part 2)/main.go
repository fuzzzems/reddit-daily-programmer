// [2015-06-17] Challenge #219 [Intermediate] To-do list (Part 2)
// http://www.reddit.com/r/dailyprogrammer/comments/3a64hq/20150617_challenge_219_intermediate_todo_list/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Note struct {
	Content string
	Cats    []string
}

type Notebook []Note

var notebook = Notebook{}

func addNote(s string, cat ...string) {

	if len(cat) == 0 {
		cat = append(cat, "Uncategorized")
		notebook = append(notebook, Note{Content: s, Cats: cat})
	} else {
		notebook = append(notebook, Note{Content: s, Cats: cat})
	}
}

func viewAllNotes() {

	fmt.Println("--------[View All Notes]--------")

	for i := range notebook {
		fmt.Print(notebook[i].Content)
		for j := range notebook[i].Cats {
			fmt.Print(" ", notebook[i].Cats[j])
		}
		fmt.Println()
	}

}

func viewByCat(s ...string) {

	fmt.Printf("--------%s--------\n", s)

	for i := range notebook {

		for j := range s {

			if !containsCat(s[j], notebook[i].Cats) {
				break
			}

			if j == len(s)-1 {
				fmt.Println(notebook[i].Content)
			}

		}

	}

}

func containsCat(s string, cats []string) bool {
	for i := range cats {
		if s == cats[i] {
			return true
		}
	}

	return false
}

func updateNote(origin string, s string) {
	for i := range notebook {
		if notebook[i].Content == origin {
			notebook[i].Content = s
			break
		}
	}
}

func deleteNote(s string) {
	for i := range notebook {
		if s == notebook[i].Content {
			notebook = append(notebook[:i], notebook[i+1:]...)
			break
		}

		if i == len(notebook)-1 {
			fmt.Printf("Note: \"%s\" not found", s)
		}

	}
}

func main() {

	file, _ := ioutil.ReadFile("./data.json")

	err := json.Unmarshal(file, &notebook)

	if err != nil {
		ioutil.WriteFile("./data.json", nil, 0644)
	}

	addNote("A pixel is not a pixel is not a pixel", "Programming")
	addNote("The Scheme Programming Language", "Programming")
	addNote("Memory in C", "Programming")

	addNote("Modes in Folk Music", "Music")
	addNote("The use of the Melodic Minor Scale", "Music")

	addNote("Create Sine Waves in C", "Programming", "Music")
	addNote("Haskell's School of Music", "Programming", "Music")
	addNote("Algorithmic Symphonies from one line of code", "Programming", "Music")

	viewByCat("Programming")
	viewByCat("Music")
	viewByCat("Programming", "Music")

	updateNote("Create Sine Waves in C", "Create Sine Waves in Python")

	viewAllNotes()

	b, _ := json.Marshal(notebook)
	ioutil.WriteFile("./data.json", b, 0644)

}
