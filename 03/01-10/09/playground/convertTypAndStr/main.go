package main

import "fmt"

type typ string

const (
	typS typ = "s"
	typA typ = "a"
)

var typs = map[typ]struct{}{typS: struct{}{}, typA: struct{}{}}

func main() {
	s := "s"

	t := typ(s)
	fmt.Println(t)

	tC := typ("c")
	fmt.Println(tC)

	t2, ok := typs[t]
	fmt.Println(t2, ok)

	t3, ok := typs[tC]
	fmt.Println(t3, ok)

}
