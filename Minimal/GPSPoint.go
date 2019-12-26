package main
import (
	"fmt"
)
// Struct med adresse/navn og gps-coords

// Node for address
type Node struct {
	Navn string
	Failed bool
	Adresse Address
	Lat float64
	Lon float64
}
// Address is container
type Address struct {
	GateNavn string
	HusNr int
	PostSted string
	Bokstav string
}
func (a *Address) toString() string {
	return fmt.Sprint(a.GateNavn," ", a.HusNr,", ", a.PostSted)
}

