package chapter1

import (
	"fmt"
	"os"
)

func Echo2() {
	for index, args := range os.Args[1:] {
		fmt.Println(index)
		fmt.Println(args)
	}
}
