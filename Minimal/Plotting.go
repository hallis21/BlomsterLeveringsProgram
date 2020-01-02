package main

import (
	"time"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
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

func createGEOJSONRoute(nodes []NodeContainer) string {
	toReturn := "{\"type\": \"FeatureCollection\",\"features\": ["
	first := true
	// Sort array by waypointindex
	sortedArray := make([]NodeContainer, len(nodes))
	for _, n := range nodes {
		if !(n.waypointIndex >= len(nodes)) {
			sortedArray[n.waypointIndex] = n
		}
	}
	for _, n := range sortedArray {
		st := ","
		if first {
			st = ""
			first = false
		}
		nstr := createSingularGOEJSONRoute(n)
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

func createSingularGOEJSONRoute(nc NodeContainer) string {
	node := nc.node
	toReturn := fmt.Sprintf("{\"type\": \"Feature\",\"geometry\": {\"type\": \"Point\",\"coordinates\": [%v,%v]},\"properties\": {\"name\": \"%v\",\"adresse\":\"%v\",\"marker-color\": \"#ff0000\", \"index\": \"%v\"}}", node.Lon, node.Lat, node.Navn, node.Adresse.toString(), nc.waypointIndex)
	if node.Navn == "" {
		return ""
	}
	return toReturn
}




// NodeContainer is a cointaer for route calculating
type NodeContainer struct {
	node Node
	index int
	waypointIndex int
}
// AreaContainer holds nodes from a specific area to plot optimal route
type AreaContainer struct {
	navn string
	noder []Node
	shortestRoute []NodeContainer
}

func (ac *AreaContainer) add(n Node) {
	ac.noder = append(ac.noder, n)
}

// Shortest route with start node as home
func shortestRoute(nodes []NodeContainer) []NodeContainer {

	// 
	// nn := NodeContainer{Node{"Hjem", false, Address{"Hjem", 0, "Raxta", ""},59.413895, 11.355041},0,0}
	client = http.Client{Timeout: 30 * time.Second} 
	url := "http://router.project-osrm.org/trip/v1/driving/11.355041,59.413895"
	coords := ""
	first := false
	for i, eC := range nodes {
		eC.index = i
		e := eC.node
		p := ";"
		if first {
			p = ""
			first = false
		}
		coords = fmt.Sprintf("%v%v%v,%v",coords, p, e.Lon, e.Lat)
	}
	fmt.Println(fmt.Sprintf("%v%v?overview=false", url, coords))
	// match JSON to nodeContainers
	resp, err := client.Get(fmt.Sprintf("%v%v?overview=false", url, coords))
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var temp map[string]interface{}
	// fmt.Println(string(b))
	err = json.Unmarshal(b, &temp)
	if err != nil {
		panic(err)
	}
	code := temp["code"]
	// If it cant plot it returns an empty array
	if code != "Ok" {
		fmt.Println("OOY")
		return make([]NodeContainer, 0)
	}
	waypoints := temp["waypoints"].([]interface{})
	for i := range waypoints {
		if i != 0 {
			// Set waypoint index
			w := waypoints[i].(map[string]interface{})
			wIndex := w["waypoint_index"].(float64)
			
			// node.waypointIndex = int(wIndex)
			nodes[i-1].waypointIndex = int(wIndex)
			// fmt.Println(node.waypointIndex)
			// Assert equal coords
		}
	}




	return nodes
}

// Returns a list of containers that represent a given area on the map
func createBulk(nodes []Node) []AreaContainer {
	// AreaContainers for main areas and gokk
	area1 := AreaContainer{}
	area1.navn = "Sentrum"
	area2 := AreaContainer{}
	area2.navn = "Degernes retning"
	area3 := AreaContainer{}
	area3.navn = "Åsen retning"
	area4 := AreaContainer{}
	area4.navn = "Strømfossveien ish"
	area5 := AreaContainer{}
	area5.navn = "(Bregneveien ish)"
	area6 := AreaContainer{}
	area6.navn = "Rest / Gokk"


	for _, n := range nodes {
		switch findArea(n) {
		case 1:
			area1.add(n)
		case 2:
			area2.add(n)
		case 3:
			area3.add(n)
		case 4:
			area4.add(n)
		case 5:
			area5.add(n)
		default:
			area6.add(n)
		}
	}
	



	return make([]AreaContainer, 0)
}

func findArea(node Node) int {



	return 0
}