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
	// digits := [][]rune{[]rune("ZERO"), []rune("ONE"), []rune("TWO"), []rune("THREE"), []rune("FOUR"), []rune("FIVE"), []rune("SIX"), []rune("SEVEN"), []rune("EIGHT"), []rune("NINE")}

	// var abs [26][]int
	// for num, d := range digits {
	// 	for _, c := range d {
	// 		abs[int(c-'A')] = append(abs[int(c-'A')], num)
	// 	}
	// } // EFGHINORSTUVWXZ
	// for ab, nums := range abs {
	// 	if len(nums) > 0 {
	// 		fmt.Fprintf(outfile, "%c: %x", ab+'A', nums)
	// 		fmt.Fprintf(outfile, "\n")
	// 	}
	// }

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
		var numCount [10]int
		numCount[0] = alphabetaCount[int('Z'-'A')]
		numCount[2] = alphabetaCount[int('W'-'A')]
		numCount[4] = alphabetaCount[int('U'-'A')]
		numCount[6] = alphabetaCount[int('X'-'A')]
		numCount[7] = alphabetaCount[int('S'-'A')] - numCount[6]
		numCount[8] = alphabetaCount[int('G'-'A')]
		numCount[3] = alphabetaCount[int('H'-'A')] - numCount[8]
		numCount[1] = alphabetaCount[int('O'-'A')] - numCount[0] - numCount[2] - numCount[4]
		numCount[5] = alphabetaCount[int('V'-'A')] - numCount[7]
		numCount[9] = alphabetaCount[int('I'-'A')] - numCount[5] - numCount[6] - numCount[8]
		var phone []int
		for num, count := range numCount {
			for i := 0; i < count; i++ {
				phone = append(phone, num)
			}
		}
		output := make([]rune, len(phone))
		for i, n := range phone {
			output[i] = '0' + rune(n)
		}
		// output case result
		fmt.Fprintf(outfile, "%s\n", string(output))
	}
}
