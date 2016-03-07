// [2016-03-07] Challenge #257 [Easy] In what year were most presidents alive?
// https://www.reddit.com/r/dailyprogrammer/comments/49aatn/20160307_challenge_257_easy_in_what_year_were/

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	file, err := os.Open("input.csv")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	// Get rid of the headers
	reader.Read()

	years, l := make(map[int]int, 0), 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		birth, death := getDateRange(record[1], record[3])

		for start := birth; start <= death; start++ {
			years[start] = years[start] + 1

			if l < years[start] {
				l = years[start]
			}
		}
	}

	for k, v := range years {
		if l == v {
			fmt.Println(k)
		}
	}

}

func getDateRange(born string, death string) (int, int) {
	return formatDate(born), formatDate(death)
}

func formatDate(s string) int {
	if s == "" {
		s = time.Now().Format("Jan 02 2006")
	}

	t := strings.Fields(s)

	year, err := strconv.Atoi(t[2])
	if err != nil {
		panic(err)
	}

	return year
}
