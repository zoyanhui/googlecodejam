package main

import (
	// "bufio"
	"fmt"
	// "io"
	// "io/ioutil"
	"log"
	"os"
)

// func main() {
// 	fmt.Println("Hello, playground")
// }

func main() {
	// file, err := os.Open("testcase.txt")
	file, err := os.Open("B-large-practice.in")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	outfile, outfileerr := os.OpenFile("B-large-practice.out", os.O_WRONLY|os.O_CREATE, 0600) // os.O_WRONLY|os.O_CREATE|os.O_APPEND
	if outfileerr != nil {
		log.Fatal(outfileerr)
		return
	}
	defer outfile.Close()
	// outfile = os.Stdout
	casenum := 0
	// scanner := bufio.NewScanner(file)
	// scanner.Text()  // read line
	fmt.Fscanf(file, "%d\n", &casenum)
	for casei := 0; casei < casenum; casei++ {
		fmt.Fprintf(outfile, "Case #%d: ", casei+1)
		heights := make([]int, 2501)
		N := 0
		fmt.Fscanf(file, "%d\n", &N)
		for j := 0; j < 2*N-1; j++ {
			for i := 0; i < N; i++ {
				each := 0
				fmt.Fscanf(file, "%d", &each)
				heights[each] ^= 1
			}
		}
		for idx, h := range heights {
			if h == 1 {
				fmt.Fprintf(outfile, "%d ", idx)
			}
		}
		fmt.Fprintf(outfile, "\n")
	}
	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
}
