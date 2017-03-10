package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// file, err := os.Open("testcase.txt")
	file, err := os.Open("C-large-practice.in")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	outfile, outfileerr := os.OpenFile("C-large-practice.out", os.O_WRONLY|os.O_CREATE, 0600) // os.O_WRONLY|os.O_CREATE|os.O_APPEND
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
		var N int
		fmt.Fscanf(file, "%d\n", &N)
		input := make([]int, N+1)
		for i := 1; i <= N; i++ {
			fmt.Fscanf(file, "%d", &input[i])
		}
		maxcircel := 0              // count the max circle
		taillen := make([]int, N+1) // count the max composition
		for i := 1; i <= N; i++ {
			visit := make([]bool, N+1)
			j := i
			countVisit := 0
			for !visit[j] {
				visit[j] = true
				countVisit++
				if input[input[j]] != j && taillen[j]+1 > taillen[input[j]] {
					taillen[input[j]] = taillen[j] + 1
				}
				j = input[j]
			}
			if j == i {
				if countVisit > maxcircel {
					maxcircel = countVisit
				}
			}
		}
		maxcompo := 0
		visit := make([]bool, N+1)
		for i := 1; i <= N; i++ {
			if visit[i] {
				continue
			}
			if input[input[i]] == i {
				maxcompo += 2 + taillen[i] + taillen[input[i]]
				visit[i] = true
				visit[input[i]] = true
			}
		}
		output := maxcircel
		if maxcompo > output {
			output = maxcompo
		}
		fmt.Fprintf(outfile, "%d\n", output)
	}
}
