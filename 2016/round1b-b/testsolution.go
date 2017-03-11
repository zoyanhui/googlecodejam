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
		codersChange := make([]bool, len(coders))
		jammersChange := make([]bool, len(jammers))
		var diff int
		for idx, c := range coders {
			j := jammers[idx]
			// ci := int(c - '0')
			// ji := int(jammers[idx] - '0')
			if diff == 0 {
				if c == '?' && j == '?' {
					codersChange[idx], jammersChange[idx] = true, true
					coders[idx], jammers[idx] = '0', '0'
				} else if c == '?' {
					codersChange[idx] = true
					coders[idx] = j
				} else if j == '?' {
					jammersChange[idx] = true
					jammers[idx] = c
				} else {
					ndiff := int(c-'0') - int(j-'0')
					if ndiff >= -5 && ndiff <= 5 {
						diff = ndiff
					} else if ndiff > 5 {
						if shiffup(&coders, &jammers, idx-1, -1, &codersChange, &jammersChange) {
							diff = -1
						} else {
							diff = ndiff
						}
					} else {
						if shiffup(&coders, &jammers, idx-1, 1, &codersChange, &jammersChange) {
							diff = 1
						} else {
							diff = ndiff
						}
					}
				}
			} else if diff > 0 {
				if c == '?' && j == '?' {
					codersChange[idx], jammersChange[idx] = true, true
					coders[idx], jammers[idx] = '0', '9'
				} else if c == '?' {
					codersChange[idx] = true
					coders[idx] = '0'
				} else if j == '?' {
					jammersChange[idx] = true
					jammers[idx] = '9'
				} else {
					continue
				}
			} else {
				if c == '?' && j == '?' {
					codersChange[idx], jammersChange[idx] = true, true
					coders[idx], jammers[idx] = '9', '0'
				} else if c == '?' {
					codersChange[idx] = true
					coders[idx] = '9'
				} else if j == '?' {
					jammersChange[idx] = true
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

func shiffup(coders *[]rune, jammers *[]rune, idx int, val int, codersChange *[]bool, jammersChange *[]bool) bool {
	if idx < 0 {
		return false
	}
	if val < 0 {
		if (*codersChange)[idx] {

		}

		if (*coders)[idx] > '0' {
			if (*codersChange)[idx] {
				(*coders)[idx] -= rune(1)
				return true
			}
			if (*jammers)[idx] < '9' {
				if (*jammersChange)[idx] {
					(*jammers)[idx] += rune(1)
					return true
				}
				return false
			}
			return false
		} else {
			if shiffup(coders, jammers, idx-1, val, codersChange, jammersChange) {
				return true
			}
			if (*jammers)[idx] < '9' {
				if (*jammersChange)[idx] {
					(*jammers)[idx] += rune(1)
					return true
				}
				return false
			} else {
				if shiffup(coders, jammers, idx, val, codersChange, jammersChange) {
					return true
				}
			}
		}
	} else {
		if (*jammers)[idx] > '0' {
			if (*jammersChange)[idx] {
				(*jammers)[idx] -= rune(1)
				return true
			}
			if (*coders)[idx] < '9' {
				if (*codersChange)[idx] {
					(*coders)[idx] += rune(1)
					return true
				}
				return false
			} else {
				if shiffup(coders, jammers, idx-1, val, codersChange, jammersChange) {
					return true
				}
			}
			return false
		} else {
			if shiffup(coders, jammers, idx-1, val, codersChange, jammersChange) {
				return true
			}
			if (*coders)[idx] < '9' {
				if (*codersChange)[idx] {
					(*coders)[idx] += rune(1)
					return true
				}
				return false
			} else {
				if shiffup(coders, jammers, idx-1, val, codersChange, jammersChange) {
					return true
				}
			}
		}
	}
	return false
}
