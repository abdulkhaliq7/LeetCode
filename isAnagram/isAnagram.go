package main

import "fmt"


func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freqs := make(map[rune]int)

	freqt := make(map[rune]int)

	for _, s1 := range s {
		freqs[s1]++
		fmt.Println(freqs)
	}

	for _, t1 := range t {
		freqt[t1]++
	}

	for char, count := range freqs {
		if freqt[char] != count {
			return false
		}
	}

	return true

}

func main() {
	s := "halp"
	t := "hala"

	fmt.Println(isAnagram(s, t))
}
