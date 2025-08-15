package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var text string

	fmt.Print("Input one line of words (S) : ")
	reader := bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')

	vowelSeen := make(map[rune]bool)
	consonantSeen := make(map[rune]bool)
	var vowelOrder, consonantOrder []rune
	var foundVowels, foundConsonants []rune

	for _, r := range text {
		if unicode.IsSpace(r) {
			continue
		}
		r = unicode.ToLower(r)

		if strings.ContainsRune("aiueo", r) {
			foundVowels = append(foundVowels, r)
			if !vowelSeen[r] {
				vowelSeen[r] = true
				vowelOrder = append(vowelOrder, r)
			}
		} else {
			foundConsonants = append(foundConsonants, r)
			if !consonantSeen[r] {
				consonantSeen[r] = true
				consonantOrder = append(consonantOrder, r)
			}
		}
	}

	// sort sesuai urutan kemunculan
	sortLetters := func(order []rune, letters []rune) string {
		var sorted []rune
		for _, o := range order {
			for _, l := range letters {
				if o == l {
					sorted = append(sorted, l)
				}
			}
		}
		return string(sorted)
	}

	fmt.Println("Vowel Characters :", sortLetters(vowelOrder, foundVowels))
	fmt.Println("Consonant Characters :", sortLetters(consonantOrder, foundConsonants))
}
