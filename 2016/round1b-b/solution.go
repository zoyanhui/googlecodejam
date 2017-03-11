package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// file, err := os.Open("testcase2.txt")
	file, err := os.Open("B-small-practice.in")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	outfile, outfileerr := os.OpenFile("B-small-practice.out", os.O_WRONLY|os.O_CREATE, 0600) // os.O_WRONLY|os.O_CREATE|os.O_APPEND
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
		// do solve case
		var coderstr string
		var jammerstr string
		fmt.Fscanf(file, "%s %s\n", &coderstr, &jammerstr)
		coders := []rune(coderstr)
		jammers := []rune(jammerstr)
		var diff int
		for idx, c := range coders {
			j := jammers[idx]
			// ci := int(c - '0')
			// ji := int(jammers[idx] - '0')
			if diff == 0 {
				if c == '?' && j == '?' {
					if idx >= len(coders)-1 || (coders[idx+1] == '?' || jammers[idx+1] == '?') {
						coders[idx], jammers[idx] = '0', '0'
					} else {
						ndiff := int(coders[idx+1]-'0') - int(jammers[idx+1]-'0')
						if ndiff >= -5 && ndiff <= 5 {
							coders[idx], jammers[idx] = '0', '0'
						} else if ndiff > 5 {
							coders[idx], jammers[idx] = '0', '1'
							diff = -1
						} else {
							coders[idx], jammers[idx] = '1', '0'
							diff = 1
						}
					}
				} else if c == '?' {
					if idx >= len(coders)-1 || (coders[idx+1] == '?' || jammers[idx+1] == '?') {
						coders[idx] = j
					} else {
						ndiff := int(coders[idx+1]-'0') - int(jammers[idx+1]-'0')
						if ndiff >= -5 && ndiff <= 5 {
							coders[idx] = j
						} else if ndiff > 5 {
							if j > '0' {
								coders[idx] = j - rune(1)
								diff = -1
							} else {
								coders[idx] = '0'
							}
						} else {
							if j < '9' {
								coders[idx] = j + rune(1)
								diff = 1
							} else {
								coders[idx] = '9'
							}
						}
					}
				} else if j == '?' {
					if idx >= len(coders)-1 || (coders[idx+1] == '?' || jammers[idx+1] == '?') {
						jammers[idx] = c
					} else {
						ndiff := int(coders[idx+1]-'0') - int(jammers[idx+1]-'0')
						if ndiff >= -5 && ndiff <= 5 {
							jammers[idx] = c
						} else if ndiff > 5 {
							if c < '9' {
								jammers[idx] = c + rune(1)
								diff = -1
							} else {
								jammers[idx] = '9'
							}
						} else {
							if c > '0' {
								jammers[idx] = c - rune(1)
								diff = 1
							} else {
								jammers[idx] = '0'
							}
						}
					}
				} else {
					diff = int(c-'0') - int(j-'0')
				}
			} else if diff > 0 {
				if c == '?' && j == '?' {
					coders[idx], jammers[idx] = '0', '9'
				} else if c == '?' {
					coders[idx] = '0'
				} else if j == '?' {
					jammers[idx] = '9'
				} else {
					continue
				}
			} else {
				if c == '?' && j == '?' {
					coders[idx], jammers[idx] = '9', '0'
				} else if c == '?' {
					coders[idx] = '9'
				} else if j == '?' {
					jammers[idx] = '0'
				} else {
					continue
				}
			}
		}
		// output case result
		fmt.Fprintf(outfile, "%s %s\n", string(coders), string(jammers))
	}
}
