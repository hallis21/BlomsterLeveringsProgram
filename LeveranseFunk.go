package main
// Div funksjonalitet rundt leveranse 
// Med untak av opprettelse av objectet / structen
import (
	"net/http"
	"fmt"
	"time"
	"encoding/json"
)
const url string = "https://nominatim.openstreetmap.org/search?q="
const format string = "&format=json&polygon=1&addressdetails=1"
var client http.Client = http.Client{Timeout: 30 * time.Second,}

func (l *leveranse) addCoords() error {
	// var coords coordinate 
	resp, err := client.Get(url+l.formatAddressLev()+format)
	if err != nil {
		return err
	}
	var temp struct {
		lat string "json:'0:lat'"
		lon string "json:'0:lon'"
	}
	err = json.NewDecoder(resp.Body).Decode(&temp)
	fmt.Println(temp)
	if err != nil {
		return err
	}
	return nil
}

func (l * leveranse) formatAddressLev() string {
	gateNavn := l.Adr.gateNavn
	husNummer, postNummer := l.Adr.husNummer, l.Adr.postNummer
	return fmt.Sprint(husNummer,"+"+gateNavn+",+",postNummer,",+Norge")
}

func main() {
	b := createLeveranse(5)
	b.setAddress(23, 1890, "haldenveien", "", "Rakkestad")
	err := b.addCoords()
	if err != nil {
		fmt.Println("oof")
	}
}