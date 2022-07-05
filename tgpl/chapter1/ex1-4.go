package chapter1

import (
	"bufio"
	"fmt"
	"os"
)

//没有解决，此方法读不到文件，不知道啥意思
func Ex1_4() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Println(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Println("%d\t%s\n", n, line)
			}
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
