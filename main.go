package main

import (

)
// NyTest is stuff
type NyTest struct {
	v,g,t int
}

func main() {
	b := createLeveranse(5)
	b.setAddress(23, 1890, "haldenveien", "", "Rakkestad")
	b.print()
}