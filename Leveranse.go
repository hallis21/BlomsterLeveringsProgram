package main

import (
	"time"
	"fmt"
)

type leveranse struct {
	LeveranseID int
	DatoLagd time.Time // Dato lagt til i systemet
	Adr adresse
	CorrectGPS bool // Om adressen ble funnet på kartet - om ting stemmer overens
	Pers person
	LeveringsTid *time.Time
	ErLevert *bool
	Attempts *[]attempt
	Kommentar *string
}
// Alt settes ved lesing fra database
// ved innlegg av ny leveranse kjører man bare en update som henter ut
func createLeveranse(ID int) leveranse {
	l := leveranse{LeveranseID: ID, DatoLagd: time.Now()}
	return l
}
// Setters
func (l *leveranse) setAddress(husNummer, postNummer int, gateNavn, leiglighetsNr string) {
	l.Adr = adresse{gateNavn, husNummer,leiglighetsNr, postNummer, retrived{}, false}
}

func (l *leveranse) setPerson(PersonID int, Fornavn, Etternavn, Telefonnummer, Kommentar string){
	l.Pers = person{PersonID, Fornavn, Etternavn, Telefonnummer, Kommentar}
}
func (l *leveranse) print() {
	fmt.Println("ID: ", l.LeveranseID)
	fmt.Println("Adresse:", l.Adr.gateNavn,l.Adr.husNummer,",",l.Adr.postNummer)
}

type adresse struct {
	gateNavn string
	husNummer int
	leiglighetsNr string // 205 || 2A || A
	postNummer int // typeCheck 4 siffer ved input (ork)
	hentetUt retrived
	provdHentet bool
}
type person struct {
	PersonID int
	Fornavn string
	Etternavn string
	Telefonnummer string // typecheck ved input
	Kommentar string
}

type attempt struct {
	LeveranseID int // referanse til leveransen
	Dato time.Time
	Kommentar string
}

type retrived struct {
	feilet bool
	displayName string
	gateNavn string
	husNummer string 
	postNummer string
	Lat string
	Lon string
}

