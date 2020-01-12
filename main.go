package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"unicode/utf8"

	_ "github.com/go-sql-driver/mysql"
)

//структура для хранения данных объявления
type AdStruct struct {
	AdPrice string   `json:"ad_price"`
	AdName  string   `json:"ad_name"`
	AdBody  string   `json:"ad_body"`
	AdPhoto []string `json:"photo"`
}

//получение конкретного объявления по id
// curl --header "Content-Type: application/json" --request GET --data '{"ad_id":"1"}' http://localhost:8080/getAdById?fields=1
//or
//curl --header "Content-Type: application/json" --request GET --data '{"ad_id":"1"}' http://localhost:8080/getAdById
func getAdById(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		fmt.Fprintln(w, 405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, 500)
		return
	}

	var readId map[string]interface{}

	err = json.Unmarshal(body, &readId)
	if err != nil {
		fmt.Fprintln(w, 500)
		return
	}

	idToGet, inBody := readId["ad_id"]
	if !inBody {
		fmt.Fprintln(w, 400)
		return
	}

	bufId, ok := idToGet.(string)
	if !ok {
		fmt.Fprintln(w, 415)
		return
	}
	id, err := strconv.Atoi(bufId)

	resRow := database.QueryRow("SELECT ad_name, ad_value, ad_first_photo, ad_second_photo, ad_third_photo, ad_price  FROM ad WHERE id=?", id)
	fmt.Println(id)
	ad := AdStruct{}
	firstPhoto := ""
	secondPhoto := ""
	thirdPhoto := ""

	err = resRow.Scan(&ad.AdName, &ad.AdBody, &firstPhoto, &secondPhoto, &thirdPhoto, &ad.AdPrice)
	if err != nil {
		fmt.Fprintln(w, 500)
		return
	}

	fields := r.URL.Query().Get("fields")
	if fields == "" {
		fmt.Fprintln(w, ad.AdName, ad.AdPrice, firstPhoto)
	} else {
		fmt.Fprintln(w, ad.AdName, ad.AdBody, firstPhoto, secondPhoto, thirdPhoto, ad.AdPrice)
	}

}

//curl --header "Content-Type: application/json" --request POST --data '{"ad_price": "321", "ad_name":"wow", "ad_body":"new test", "photo": ["first", "second", "third"]}' http://localhost:8080/createAd
func createАd(w http.ResponseWriter, r *http.Request) {

	var lastId int64

	if r.Method != http.MethodPost {
		fmt.Fprintln(w, 405, lastId)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, 500, lastId)
		return
	}

	data := &AdStruct{}
	err = json.Unmarshal(body, data)
	if err != nil {
		fmt.Fprintln(w, 500, lastId)
		return
	}

	if data.AdBody == "" || data.AdName == "" || data.AdPrice == "" || data.AdPhoto == nil {

		fmt.Fprintln(w, 400, lastId)
		return
	}

	if utf8.RuneCountInString(data.AdName) > 200 || utf8.RuneCountInString(data.AdBody) > 1000 {
		fmt.Fprintln(w, 400, lastId)
		return
	}

	if len(data.AdPhoto) != 3 {

		fmt.Fprintln(w, 400, lastId)
		return
	}

	firstPhoto := data.AdPhoto[0]
	secondPhoto := data.AdPhoto[1]
	thirdPhoto := data.AdPhoto[2]

	result, err := database.Exec("INSERT INTO ad (ad_name, ad_value, ad_first_photo, ad_second_photo, ad_third_photo, ad_price) VALUES (?, ?, ?, ?, ?, ?) ",
		data.AdName,
		data.AdBody,
		firstPhoto,
		secondPhoto,
		thirdPhoto,
		data.AdPrice,
	)

	if err != nil {
		fmt.Fprintln(w, 500, lastId)
		return
	}

	lastId, err = result.LastInsertId()
	if err != nil {
		fmt.Fprintln(w, 500, lastId)
		return
	}

	fmt.Fprintln(w, 200, lastId)
	return
}

//получение списка объявлений
//curl --header "Content-Type: application/json" --request GET --data '{"page_num":"1"}' http://localhost:8080/getListOfAd
func getListOfAd(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		fmt.Fprintln(w, 405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, 500)
		return
	}

	var readPage map[string]interface{}

	err = json.Unmarshal(body, &readPage)
	if err != nil {
		fmt.Fprintln(w, 500)
		return
	}

	pageToGet, inBody := readPage["page_num"]
	if !inBody {
		fmt.Fprintln(w, 400)
		return
	}

	bufPage, ok := pageToGet.(string)
	if !ok {
		fmt.Fprintln(w, 415)
		return
	}

	page, err := strconv.Atoi(bufPage)

	str := "SELECT ad_name, ad_first_photo, ad_price FROM ad"

	priceSort := r.URL.Query().Get("price_sort")
	dataSort := r.URL.Query().Get("date_sort")

	if priceSort == "1" && dataSort == "1" {
		str = str + " ORDER BY ad_price, ad_creation_date"
	}

	if priceSort == "1" && dataSort == "0" {
		str = str + " ORDER BY ad_price, ad_creation_date DESC"
	}

	if priceSort == "0" && dataSort == "1" {
		str = str + " ORDER BY ad_price DESC, ad_creation_date"
	}

	if priceSort == "0" && dataSort == "0" {
		str = str + " ORDER BY ad_price DESC, ad_creation_date DESC"
	}

	if priceSort == "0" && dataSort == "" {
		str = str + " ORDER BY ad_price DESC"
	}

	if priceSort == "1" && dataSort == "" {
		str = str + " ORDER BY ad_price"
	}

	if priceSort == "" && dataSort == "1" {
		str = str + " ORDER BY ad_creation_time"
	}

	if priceSort == "" && dataSort == "0" {
		str = str + " ORDER BY ad_creation_time DESC"
	}

	str = str + " LIMIT ?, ?;"

	rows, err := database.Query(str, (page-1)*10, 10)
	if err != nil {
		fmt.Fprintln(w, 500)
	}
	var adPrice int64
	adName := ""
	adFirstPhoto := ""

	for rows.Next() {
		newErr := rows.Scan(&adName, &adFirstPhoto, &adPrice)
		if newErr != nil {
			fmt.Fprintln(w, newErr)
			continue
		}
		fmt.Fprintln(w, adName, adFirstPhoto, adPrice)
	}

	return
}

//база даданнынныхх
var database *sql.DB

func main() {
	db, err := sql.Open("mysql", "root:134652@tcp(localhost:3306)/adTable")
	if err != nil {
		db.Close()
		return
	}

	database = db

	http.HandleFunc("/getAdById", getAdById)
	http.HandleFunc("/createAd", createАd)
	http.HandleFunc("/getListOfAd", getListOfAd)
	http.ListenAndServe(":8080", nil)
	return
}
