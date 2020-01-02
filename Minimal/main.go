package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"strings"
	"bufio"
	"os"
	"strconv"
)
// Vindu med points som viser hvor en node er

const (
	url = "https://ws.geonorge.no/adresser/v1/sok?"
	opts = ""
)
// NoderToSync er noder som skal legges til i databasen
var NoderToSync []Node


var client = http.Client{Timeout: 30 * time.Second}
func setCoords(n *Node) error {
	bokstav := n.Adresse.Bokstav
	if n.Adresse.Bokstav == "?" {
		bokstav = ""
	}
	sok := fmt.Sprintf("%vobjtype=Vegadresse&adressenavn=%v&nummer=%v&bokstav=%v&poststed=%v", url, strings.Replace(n.Adresse.GateNavn, " ", "%", 1000), n.Adresse.HusNr, bokstav, n.Adresse.PostSted)
	resp, err := client.Get(sok)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print("yo2")
		return err
	}
	var temp map[string]interface{}
	// fmt.Println(string(b))
	err = json.Unmarshal(b, &temp)
	if err != nil {
		fmt.Print(err)
		return err
	}
	m := temp["metadata"].(map[string]interface{})
	// meta := m["metadata"].(map[string]interface{})
	antall := m["totaltAntallTreff"].(float64)
	if antall == 0 {
		myErr := &plotError{"Null treff: "+n.Navn}
		n.Lat = 0.0
		n.Lon = 0.0
		return myErr
	}
	adr := temp["adresser"].([]interface{})
	if antall > 1 {
		fmt.Println("Flere treff funnet: ", n.Navn)
	}
	adr0 := adr[0].(map[string]interface{})

	deg := adr0["representasjonspunkt"].(map[string]interface{})
	n.Lat = deg["lat"].(float64)
	n.Lon = deg["lon"].(float64)

	return err
}
// plotError
type plotError struct {message string}
func (p *plotError) Error() string {
	return p.message
}


//Finner lat lon fra OSM
func createNode(navn, gateNavn string, husNr int, bokstav string, postSted string) Node {
	n := Node{}
	n.Navn = navn
	n.Adresse.GateNavn = gateNavn
	n.Adresse.HusNr = husNr
	if bokstav == "" {
		bokstav = "?"
	}
	n.Adresse.Bokstav = bokstav
	n.Adresse.PostSted = postSted
	err := setCoords(&n)
	if err != nil {
		fmt.Println(err)
		n.Failed = true
	}
	NoderToSync = append(NoderToSync, n)
	syncDatabase()
	return n
}

func createNodesInTerminal() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("Ny node (Skrive Exit for å avslutte): ")
		fmt.Print("Navn: ")
		navn, _ := reader.ReadString('\n')
		navn = strings.Replace(navn, "\n", "", 1)
		if navn == "Exit" {
			break
		}
		fmt.Print("Gatenavn (ikke nummer): ")
		gatenavn, _ := reader.ReadString('\n')
		gatenavn = strings.Replace(gatenavn, "\n", "", 1)
		var nr int64
		var err error
		for true {
			fmt.Print("Nummer: ")
			strNr, _ := reader.ReadString('\n')
			strNr = strings.Replace(strNr, "\n", "", 1)
			nr, err = strconv.ParseInt(strNr, 10, 64)
			if err == nil {
				break
			}
			fmt.Println("Ugyldig nummer")
		}
		fmt.Print("Bokstav: ")
		bokstav, _ := reader.ReadString('\n')
		bokstav = strings.Replace(bokstav, "\n", "", 1)
		fmt.Print("Poststed: ")
		poststed, _ := reader.ReadString('\n')
		poststed = strings.Replace(poststed, "\n", "", 1)
		if navn == "" || gatenavn == "" || poststed == "" {
			fmt.Println("Ugyldig input u know")
		} else {
			if bokstav == "" {
				bokstav = "?"
			}
			createNode(navn, gatenavn, int(nr), bokstav, poststed)
		}
	}
}


func main() {
	// createNode("Ringsby2", "Haldenveien", 21, "?", "Rakkestad")
	// createNode("Yoyoyoy2", "Peer Gynts vei", 2, "?", "Rakkestad")
	// purgeAll(true)
	// createNodesInTerminal()
	// n := b
	// b = n
	// fmt.Println(n.Lat, n.Lon)
	noder := getAllNodes(opt{false, true})
	// shortestRoute(noder)
	// fmt.Println(createGEOJSON(getAllNodes(opt{true, true})))
	// fmt.Println("\n")
	nC := make([]NodeContainer, len(noder))
	for i, e := range noder {
		newC := NodeContainer{}
		newC.node = e
		nC[i] = newC
	}
	n := shortestRoute(nC)
	s := createGEOJSONRoute(n)
	fmt.Println(s)
	// fmt.Println(n.Lon)
	// var noder []node
	// noder = append(noder, n, b)
	// fmt.Printf(createGEOJSON(noder))
	// getCoords(n)


	// p1 := true
	// p2 := true
	// nodes := getAllNodes(opt{p1, p2})
	// r, _ := json.Marshal(nodes)
	// fmt.Println(string(r))
}