package main

import (
	"bufio"
	"fmt"
	"github.com/zs1621/leetcode_go/bookAlgorithms/ChapterSearch"
	"os"
	"strings"
	"time"
)

func main() {

	var (
		tokens []byte
		err    error

	)
	// bst

	var x = &ChapterSearch.NodeTree{nil, 0, nil}


	scanner := bufio.NewScanner(os.Stdin)
	t1 := time.Now()
	for scanner.Scan() {
		_, tokens, err = bufio.ScanLines(scanner.Bytes(), true)
		if err != nil {
			fmt.Println(err)
		}
		for _, value := range strings.Split(string(tokens), " ") {
			if len(value) < 4 {
				continue
			}
			x.Put(value)
		}
	}
	if scanner.Err() != nil {
		// handle error.
	}
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
	fmt.Println(x.Size)
}