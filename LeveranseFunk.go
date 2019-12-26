package main
// Div funksjonalitet rundt leveranse 
// Med untak av opprettelse av objectet / structen
import (
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"
)
const url string = "https://nominatim.openstreetmap.org/search?q="
const format string = "&format=json&polygon=1&addressdetails=1"
var client http.Client = http.Client{Timeout: 30 * time.Second,}

func (l *leveranse) addCoords() error {
	// var coords coordinate 
	resp, err := client.Get(url+l.formatAddressLevForHTTP()+format)
	// fmt.Println(url+l.formatAddressLev()+format)
	if err != nil {
		return err
	}
	b, error := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if error != nil {
		return error
	}
	var temp []interface{}
	err = json.Unmarshal(b, &temp)
	m := temp[0].(map[string]interface{})
	// big oofer here, må få satt type string
	var adr = m["address"].(map[string]interface{})
	var name string = m["display_name"].(string)
	var road string = adr["road"].(string)
	var nr string = adr["house_number"].(string)
	var code string = adr["postcode"].(string)
	var lat string = m["lat"].(string)
	var lon string = m["lon"].(string)
	hentet := retrived{false,name, road, nr, code, lat, lon}

	l.Adr.hentetUt = hentet
	l.Adr.provdHentet = true

	if err != nil {
		l.Adr.hentetUt = retrived{}
		l.Adr.hentetUt.feilet = true
		l.Adr.provdHentet = true
		return err
	}

	return nil
}

func (l *leveranse) formatAddressLevForHTTP() string {
	gateNavn := l.Adr.gateNavn
	husNummer, postNummer := l.Adr.husNummer, l.Adr.postNummer
	return fmt.Sprint(husNummer,"+"+gateNavn+",+",postNummer,",+Norge")
}

func (l *leveranse) checkCorrect() error {
	// Initial checks
	var err error
	if !l.Adr.provdHentet {
		err = l.addCoords()
	}
	if l.Adr.hentetUt.feilet {
		l.CorrectGPS = false
		return err
	}
	// Check consistency
	adr := l.Adr
	h := l.Adr.hentetUt
	if strings.ToLower(adr.gateNavn) != strings.ToLower(h.gateNavn) || strconv.Itoa(adr.husNummer) != h.husNummer || strconv.Itoa(adr.postNummer) != h.postNummer {
		l.CorrectGPS = false
		return err
	}
	// fmt.Println(h.gateNavn)
	l.CorrectGPS = true
	return err
}

// func main() {
// 	b := createLeveranse(5)
// 	b.setAddress(21, 1890, "haldenveien", "")
// 	err := b.checkCorrect()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(b.CorrectGPS)
// }