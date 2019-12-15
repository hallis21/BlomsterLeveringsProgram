package main

import (

)
// NyTest is stuff
type NyTest struct {
	v,g,t int
}

func main() {
	b := createLeveranse(5)
	b.setAddress(23, 1890, 65.333, 45.2222, true, "LissiomVeien2", "", "Rakkestad")
	b.print()
}