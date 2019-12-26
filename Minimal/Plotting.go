package main

import (
	"time"
	"net/http"
	"fmt"
)

func createGEOJSON(nodes []Node) string {
	toReturn := "{\"type\": \"FeatureCollection\",\"features\": ["
	first := true
	for _, n := range nodes {
		st := ","
		if first {
			st = ""
			first = false
		}
		nstr := createSingularGOEJSON(n)
		if nstr != "" {
			toReturn = fmt.Sprint(toReturn, st+"\n", nstr)
		}
	}
	toReturn = fmt.Sprint(toReturn, "]}\n")
	return toReturn
}

func createSingularGOEJSON(node Node) string {
	toReturn := fmt.Sprintf("{\"type\": \"Feature\",\"geometry\": {\"type\": \"Point\",\"coordinates\": [%v,%v]},\"properties\": {\"name\": \"%v\",\"adresse\":\"%v\",\"marker-color\": \"#ff0000\"}}", node.Lon, node.Lat, node.Navn, node.Adresse.toString())
	if node.Navn == "" {
		return ""
	}
	return toReturn
}


func shortestRoute(nodes []Node) {
	client = http.Client{Timeout: 30 * time.Second} 
	url := "http://router.project-osrm.org/trip/v1/driving/"
	coords := ""
	first := true
	for _, e := range nodes {
		p := ";"
		if first {
			p = ""
			first = false
		}
		coords = fmt.Sprintf("%v%v%v,%v",coords, p, e.Lon, e.Lat)
	}
	// resp, _ := client.Get(fmt.Sprintf("v%v%?overview=false", url, coords))
	fmt.Println(fmt.Sprintf("%v%v?overview=false", url, coords))

}