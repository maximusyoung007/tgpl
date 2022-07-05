package chapter1

import (
	"fmt"
	"strings"
	"time"
)

func Ex1_3() {
	res1 := ""
	res2 := ""
	s := []string{"a", "b", "c", "d", "Process", "finished", "with", "exit", "code", "0"}
	t1 := time.Now()
	for _, val := range s {
		res1 += val + " "
	}
	t2 := time.Now()
	fmt.Println(t2.Nanosecond() - t1.Nanosecond())

	t3 := time.Now()
	res2 = strings.Join(s, " ")
	t4 := time.Now()
	fmt.Println(t4.Nanosecond() - t3.Nanosecond())
	fmt.Println(res1)
	fmt.Println(res2)
}
