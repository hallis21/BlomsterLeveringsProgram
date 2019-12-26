package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)
// Generic function to that syncs database and webserver

const (
	host = "localhost"
	port = 5432
	user = "fwp"
	password = "backendbuisness"
	dbname = "flowerpowertrip"
)

var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)



// SyncDatabase is guud
func syncDatabase() {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		alert(fmt.Sprint("COULD NOT ACCESS DATABASE: ", err))
	}
	defer db.Close()
	if len(NoderToSync) > 0 {
		for _, e := range NoderToSync {
			if nodeNotInDB(db, e) {
				st, err := db.Prepare("INSERT INTO node(navn, failed, gateNavn, husNr, postSted, bokstav, lat, lon, dato) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9);")
				if err != nil {
					alert(fmt.Sprint("Off syncing prepar", err))
				}
				_, err = st.Exec(e.Navn, e.Failed, e.Adresse.GateNavn, e.Adresse.HusNr, e.Adresse.PostSted, e.Adresse.Bokstav, e.Lat, e.Lon,"now()")
				if err != nil {
					fmt.Println(err)
					alert("Off syncing exec")
				}

			} else {
				alert(fmt.Sprint("Name not unique: ", e.Navn))
			}
			
		}
		NoderToSync = make([]Node, 0)
			
	}


}
func nodeNotInDB(DB *sql.DB, n Node) bool {
	rows, err := DB.Query("SELECT navn FROM node;")
	if err != nil {
		return true
	}
	defer rows.Close()
	for rows.Next() {
		var navn string

		if err := rows.Scan(&navn); err != nil {
			}	
				if navn == n.Navn {
					return false;
				}
	}
	return true
}

type opt struct {
	idag bool
	nonFailed bool
}

func alert(s string) {
	fmt.Println(s)
}	
// GetAllNodes returns all nodes given the options
func getAllNodes(opts opt) []Node {
	var temp []Node
	where := ""
	if opts.idag {
		where = "n.dato = now()::date"
	}
	if opts.nonFailed {
		if where != "" {
			where = fmt.Sprint(where, " AND ")
		}
		where = fmt.Sprint(where, " n.failed = FALSE")
	}
	if where != "" {
		where = fmt.Sprint(" WHERE ", where)
	}
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		alert(fmt.Sprint("COULD NOT ACCESS DATABASE: ", err))
	}
	defer db.Close()
	rows, err := db.Query(fmt.Sprint("SELECT * FROM node as n", where))
	if err != nil {
		alert(fmt.Sprint("OOF: ", err))
	}
	defer rows.Close()
	for rows.Next() {
		var (
			navn string
			failed bool
			gatenavn string
			husnr int
			poststed string
			bokstav string
			lat float64
			lon float64
			dato string
		)
		if err := rows.Scan(&navn, &failed, &gatenavn, &husnr, &poststed, &bokstav, &lat, &lon, &dato); err != nil {
			fmt.Println(err)
		} else {
			temp = append(temp, Node{navn, failed, Address{gatenavn, husnr, poststed, bokstav}, lat, lon})
		}

	}

	return temp

}
// Removes entries that are not for this day
func purgeOld() int64 {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		alert(fmt.Sprint("COULD NOT ACCESS DATABASE: ", err))
	}
	defer db.Close()
	rs, err :=db.Exec("DELETE FROM node WHERE dato != NOW()::DATE")
	if err != nil {
		alert(fmt.Sprint("COULD NOT DELETE: ", err))
	}
	n, err := rs.RowsAffected()
	return n
}

func purgeAll(c bool) int64 {
db, err := sql.Open("postgres", psqlInfo)
if !c {
	return 0
}
	if err != nil {
		alert(fmt.Sprint("COULD NOT ACCESS DATABASE: ", err))
	}
	defer db.Close()
	rs, err :=db.Exec("DELETE FROM node WHERE TRUE")
	if err != nil {
		alert(fmt.Sprint("COULD NOT DELETE: ", err))
	}
	n, err := rs.RowsAffected()
	return n
}