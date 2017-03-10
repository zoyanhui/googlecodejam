package main

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
	casenum := 0
	fmt.Fscanf(file, "%d\n", &casenum)
	for casei := 0; casei < casenum; casei++ {
		fmt.Fprintf(outfile, "Case #%d: ", casei+1)
		var input string
		fmt.Fscanf(file, "%s\n", &input)
		var output string
		for _, c := range input {
			if len(output) == 0 {
				output = string(c)
				continue
			}
			if c > (rune)(output[0]) {
				output = string(c) + output
			} else if c < (rune)(output[0]) {
				output = output + string(c)
			} else {
				i := 0
				for ; i < len(output)-1; i++ {
					if output[i] < output[i+1] {
						output = output + string(c)
						break
					} else if output[i] > output[i+1] {
						output = string(c) + output
						break
					} else {
						continue
					}
				}
				if i == len(output)-1 {
					output = output + string(c)
				}
			}
		}
		fmt.Fprintf(outfile, "%s\n", output)
	}
}
