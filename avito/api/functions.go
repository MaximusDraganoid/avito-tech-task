// functions.go
package api

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

//база даданнынныхх
var Database *sql.DB

//структура для хранения данных объявления
type AdStruct struct {
	AdPrice string   `json:"ad_price"`
	AdName  string   `json:"ad_name"`
	AdBody  string   `json:"ad_body, omitempty"`
	AdPhoto []string `json:"photo"`
}

type AdResponse struct {
	ResultCode int   `json:"http_error_code"`
	GetingId   int64 `json:"id"`
}

func marshalAndWriteId(w io.Writer, errCode int, resId int64) {
	res := &AdResponse{
		ResultCode: errCode,
		GetingId:   resId,
	}

	answer, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintln(w, answer)
}

type ErrAnswer struct {
	err int `json:"errCode`
}

func murshalErrAdnWrite(w io.Writer, errCode int) {
	errRes := &ErrAnswer{
		err: errCode,
	}

	answer, err := json.Marshal(errRes)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintln(w, answer)
}

//получение конкретного объявления по id
// curl --header "Content-Type: application/json" --request GET --data '{"ad_id":"1"}' http://localhost:8080/getAdById?fields=1
//or
//curl --header "Content-Type: application/json" --request GET --data '{"ad_id":"1"}' http://localhost:8080/getAdById
func GetAdById(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		murshalErrAdnWrite(w, 405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		murshalErrAdnWrite(w, 500)
		return
	}

	var readId map[string]interface{}

	err = json.Unmarshal(body, &readId)
	if err != nil {
		murshalErrAdnWrite(w, 500)
		return
	}

	idToGet, inBody := readId["ad_id"]
	if !inBody {
		murshalErrAdnWrite(w, 400)
		return
	}

	bufId, ok := idToGet.(string)
	if !ok {
		murshalErrAdnWrite(w, 415)
		return
	}
	id, err := strconv.Atoi(bufId)

	if err != nil {
		murshalErrAdnWrite(w, 500)
	}

	resRow := Database.QueryRow("SELECT ad_name, ad_value, ad_first_photo, ad_second_photo, ad_third_photo, ad_price  FROM ad WHERE id=?", id)
	fmt.Println(id)
	ad := AdStruct{}
	firstPhoto := ""
	secondPhoto := ""
	thirdPhoto := ""

	err = resRow.Scan(&ad.AdName, &ad.AdBody, &firstPhoto, &secondPhoto, &thirdPhoto, &ad.AdPrice)
	if err != nil {
		murshalErrAdnWrite(w, 500)
		return
	}

	fields := r.URL.Query().Get("fields")
	if fields == "" {
		ad.AdBody = ""
		ad.AdPhoto = append(ad.AdPhoto, firstPhoto)
		fmt.Fprintln(w, ad.AdName, ad.AdPrice, firstPhoto)
		answer, err := json.Marshal(&ad)
		if err != nil {
			murshalErrAdnWrite(w, 500)
			return
		}
		fmt.Fprintln(w, answer)

	} else {
		answer, err := json.Marshal(&ad)
		if err != nil {
			murshalErrAdnWrite(w, 500)
			return
		}
		fmt.Fprintln(w, answer)
	}

}

//curl --header "Content-Type: application/json" --request POST --data '{"ad_price": "321", "ad_name":"wow", "ad_body":"new test", "photo": ["first", "second", "third"]}' http://localhost:8080/createAd
func CreateАd(w http.ResponseWriter, r *http.Request) {

	var lastId int64

	if r.Method != http.MethodPost {
		marshalAndWriteId(w, 400, lastId)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		marshalAndWriteId(w, 500, lastId)
		return
	}

	data := &AdStruct{}
	err = json.Unmarshal(body, data)
	if err != nil {
		marshalAndWriteId(w, 500, lastId)
		return
	}

	if data.AdBody == "" || data.AdName == "" || data.AdPrice == "" || data.AdPhoto == nil {

		marshalAndWriteId(w, 400, lastId)
		return
	}

	if utf8.RuneCountInString(data.AdName) > 200 || utf8.RuneCountInString(data.AdBody) > 1000 {
		marshalAndWriteId(w, 400, lastId)
		return
	}

	if len(data.AdPhoto) != 3 {
		marshalAndWriteId(w, 400, lastId)
		return
	}

	firstPhoto := data.AdPhoto[0]
	secondPhoto := data.AdPhoto[1]
	thirdPhoto := data.AdPhoto[2]

	result, err := Database.Exec("INSERT INTO ad (ad_name, ad_value, ad_first_photo, ad_second_photo, ad_third_photo, ad_price) VALUES (?, ?, ?, ?, ?, ?) ",
		data.AdName,
		data.AdBody,
		firstPhoto,
		secondPhoto,
		thirdPhoto,
		data.AdPrice,
	)

	if err != nil {
		marshalAndWriteId(w, 500, lastId)
		return
	}

	lastId, err = result.LastInsertId()
	if err != nil {
		marshalAndWriteId(w, 500, lastId)
		return
	}

	marshalAndWriteId(w, 200, lastId)
	return
}

type ShortDateOfAd struct {
	AdPrice      int64  `json:"ad_price"`
	AdName       string `json:"ad_name"`
	AdFirstPhoto string `json:"ad_photo"`
}

type ArrOfShortDateOfAd struct {
	Array []ShortDateOfAd `json:"ad_array"`
}

func (adArr *ArrOfShortDateOfAd) Add(ad ShortDateOfAd) {
	adArr.Array = append(adArr.Array, ad)
}

//получение списка объявлений
//curl --header "Content-Type: application/json" --request GET --data '{"page_num":"1"}' http://localhost:8080/getListOfAd
func GetListOfAd(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		murshalErrAdnWrite(w, 405)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		murshalErrAdnWrite(w, 500)
		return
	}

	var readPage map[string]interface{}

	err = json.Unmarshal(body, &readPage)
	if err != nil {
		murshalErrAdnWrite(w, 500)
		return
	}

	pageToGet, inBody := readPage["page_num"]
	if !inBody {
		murshalErrAdnWrite(w, 400)
		return
	}

	bufPage, ok := pageToGet.(string)
	if !ok {
		murshalErrAdnWrite(w, 415)
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

	rows, err := Database.Query(str, (page-1)*10, 10)
	if err != nil {
		murshalErrAdnWrite(w, 500)
	}

	var ad ShortDateOfAd
	var adArr ArrOfShortDateOfAd

	for rows.Next() {
		newErr := rows.Scan(&ad.AdName, &ad.AdFirstPhoto, &ad.AdPrice)
		if newErr != nil {
			fmt.Println(newErr)
			continue
		}
		adArr.Add(ad)
	}

	answer, err := json.Marshal(adArr)
	if err != nil {
		murshalErrAdnWrite(w, 500)
		return
	}

	fmt.Fprintln(w, answer)
	return
}
