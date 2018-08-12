package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	strMap := strings.Fields(s)
	mapLength := len(strMap)
	ret := make(map[string]int)

	for idx := 0; idx < mapLength; idx++ {
		(ret[strMap[idx]])++
	}

	return ret
}

func main() {
	fmt.Println(WordCount("Hello World Test World Test"))
}
