package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("[" + strings.Trim("      a   a    a  ", " ") + "]")
	fmt.Println("[" + strings.Join([]string{"a", "b", "c"}, "-") + "]")
}
