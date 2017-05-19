package main

import (
	"fmt"
	"os"

	"github.com/jodosha/stringutil"
)

func main() {
	input := os.Args[1]
	result := stringutil.Reverse(input)
	fmt.Println(result)
}
