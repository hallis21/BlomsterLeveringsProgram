package main

// import (
// 	"net/http"
// 	"strconv"
// 	"github.com/labstack/echo"
// 	"encoding/json"
// )

// func main() {
// 	e := echo.New()
// 	e.GET("/getGeoJSON", getGeoJSON)
// 	e.POST("/newNode", addNode)

// 	e.Logger.Fatal(e.Start(":1323"))
// }

// func addNode(c echo.Context) error {
// 	name := c.Param("name")
// 	streetName := c.Param("streetName")
// 	nrP := c.Param("streetNr")
// 	streetNr, _ := strconv.ParseInt(nrP, 0, 64)
// 	town := c.Param("town")
// 	letter := c.Param("letter")
// 	n := createNode(name, streetName, int(streetNr), letter, town)
// 	r := "Plotted"
// 	if n.Lat == 0.0 {
// 		r ="Unplotted"
// 	}

// 	return c.String(http.StatusCreated, r)
// }

// func getGeoJSON (c echo.Context) error {
// 	r := createGEOJSON(getAllNodes(opt{true, true}))

// 	return c.String(http.StatusCreated, r)
// }

// func getNodes(c echo.Context) error {
// 	p1, _ := strconv.ParseBool(c.Param("plotted"))
// 	p2, _ := strconv.ParseBool(c.Param("today"))
// 	nodes := getAllNodes(opt{p1, p2})
// 	r, _ := json.Marshal(nodes)
// 	return c.String(http.StatusOK, string(r))
// }