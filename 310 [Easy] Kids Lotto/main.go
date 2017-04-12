// [2017-04-10] Challenge #310 [Easy] Kids Lotto
// https://www.reddit.com/r/dailyprogrammer/comments/64jesw/20170410_challenge_310_easy_kids_lotto/

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	input := "Rebbeca Gann;Latosha Caraveo;Jim Bench;Carmelina Biles;Oda Wilhite;Arletha Eason;Theresa Kaczorowski;Jane Cover;Melissa Wise;Jaime Plascencia;Sacha Pontes;Tarah Mccubbin;Pei Rall;Dixie Rosenblatt;Rosana Tavera;Ethyl Kingsley;Lesia Westray;Vina Goodpasture;Drema Radke;Grace Merritt;Lashay Mendenhall;Magali Samms;Tiffaney Thiry;Rikki Buckelew;Iris Tait;Janette Huskins;Donovan Tabor;Jeremy Montilla;Sena Sapien;Jennell Stiefel"
	numNames := 3

	names := strings.Split(input, ";")

	shuffledNames := shuffle(names)

	// Add numNames to the end of the slice to prevent out of bounds errors.
	shuffledNames = append(shuffledNames, shuffledNames[:numNames]...)

	for i := 0; i < len(names); i++ {
		n := i + 1
		fmt.Printf("%s > %s \n", shuffledNames[i], strings.Join(shuffledNames[n:n+numNames], "; "))
	}
}

func shuffle(data []string) []string {

	rand.Seed(time.Now().UTC().UnixNano())

	for i := range data {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
	return data

}
