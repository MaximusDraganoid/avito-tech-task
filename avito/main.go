package main

import (
	"avito/api"
	"database/sql"
	"net/http"
)

func main() {
	var err error
	api.Database, err = sql.Open("mysql", "root:134652@tcp(localhost:3306)/adTable")
	if err != nil {
		if api.Database != nil {
			api.Database.Close()
		}
		return
	}

	err = api.Database.Ping()
	if err != nil {
		api.Database.Close()
		return
	}

	http.HandleFunc("/getAdById", api.GetAdById)
	http.HandleFunc("/createAd", api.CreateАd)
	http.HandleFunc("/getListOfAd", api.GetListOfAd)
	http.ListenAndServe(":8080", nil)
	return
}
