package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("testcase.txt")
	// file, err := os.Open("A-large-practice.in")
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
	outfile = os.Stdout
	casenum := 0
	fmt.Fscanf(file, "%d\n", &casenum)
	for casei := 0; casei < casenum; casei++ {
		fmt.Fprintf(outfile, "Case #%d: ", casei+1)
		var input string
		fmt.Fscanf(file, "%s\n", &input)
		var output string
		for _, c := range input {
			r, l := output+string(c), string(c)+output
			if r > l {
				output = r
			} else {
				output = l
			}
		}
		fmt.Fprintf(outfile, "%s\n", output)
	}
}
