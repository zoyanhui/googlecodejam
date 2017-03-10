package main

// large dataset Time Limit Exceeded

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// file, err := os.Open("testcase.txt")
	file, err := os.Open("A-large-practice.in")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	outfile, outfileerr := os.OpenFile("A-large-practice.out", os.O_WRONLY|os.O_CREATE, 0600) // os.O_WRONLY|os.O_CREATE|os.O_APPEND
	if outfileerr != nil {
		log.Fatal(outfileerr)
		return
	}
	defer outfile.Close()
	// outfile = os.Stdout

	// init constants
	digits := [][]rune{[]rune("ZERO"), []rune("ONE"), []rune("TWO"), []rune("THREE"), []rune("FOUR"), []rune("FIVE"), []rune("SIX"), []rune("SEVEN"), []rune("EIGHT"), []rune("NINE")}
	casenum := 0
	fmt.Fscanf(file, "%d\n", &casenum)
	for casei := 0; casei < casenum; casei++ {
		fmt.Fprintf(outfile, "Case #%d: ", casei+1)
		// do solve case
		var input string
		fmt.Fscanf(file, "%s\n", &input)
		var alphabetaCount [26]int
		for _, c := range input {
			alphabetaCount[int(c-'A')]++
		}
		var phone []int
		var output []rune
		var curAbCount [26]int
		if solve(&alphabetaCount, &curAbCount, &digits, &phone, 0) {
			output = make([]rune, len(phone))
			for i, n := range phone {
				output[i] = '0' + rune(n)
			}
		}
		// output case result
		fmt.Fprintf(outfile, "%s\n", string(output))
	}
}

func equals(a *[26]int, b *[26]int) bool {
	for i := 0; i < 26; i++ {
		if (*a)[i] != (*b)[i] {
			return false
		}
	}
	return true
}

func canset(digits *[][]rune, num int, alphabetaCount *[26]int, curAbCount *[26]int) bool {
	for _, c := range (*digits)[num] {
		char := int(c - 'A')
		if (*curAbCount)[char]+1 > (*alphabetaCount)[char] {
			return false
		}
	}
	return true
}

func set(digits *[][]rune, num int, curAbCount *[26]int) {
	for _, c := range (*digits)[num] {
		char := int(c - 'A')
		(*curAbCount)[char]++
	}
}

func unset(digits *[][]rune, num int, curAbCount *[26]int) {
	for _, c := range (*digits)[num] {
		char := int(c - 'A')
		(*curAbCount)[char]--
	}
}

func solve(alphabetaCount *[26]int, curAbCount *[26]int, digits *[][]rune, phone *[]int, curnum int) bool {
	if equals(curAbCount, alphabetaCount) {
		return true
	}
	for i := curnum; i <= 9; i++ {
		if canset(digits, i, alphabetaCount, curAbCount) {
			set(digits, i, curAbCount)
			*phone = append(*phone, i)
			if solve(alphabetaCount, curAbCount, digits, phone, i) {
				return true
			}
			unset(digits, i, curAbCount)
			*phone = (*phone)[:len(*phone)-1]
		}
	}
	return false
}

// if i == 0 {
//     zeros := alphabetaCount[25] // char z nums
//     for j := 0; j < zeros; j++ {
//         set(digits, i, curAbCount)
//         *phone = append(*phone, i)
//     }
//     continue
// }
