# avito-tech-task

# JSON API для сайта объявлений

## Дизайн сервиса

Язык - Golang.    
СУБД - MySQL

### Методы
|Метод HTTP|URL|Действие|
|---|---|---|
|GET|/getListOfAd?price_sort=1&date_sort=0|Получить список объявлений|
|GET|/getAdById?fields=1|Получить объявление по id|
|POST|/createAd|Создать объявление|

Параметр page в методе для получения списка объявлений является обязательным.

Для сортировки по дате и цене в запрос следует передать параметры date_sort и price_sort. 

|Значение параметра|Сортировка|
|---|---|
|0|По убыванию|
|1|По возрастанию|


### Примеры запросов
Метод создания объявления
```
curl --header "Content-Type: application/json" --request POST --data '{"ad_price": "321", "ad_name":"wow", "ad_body":"new test", "photo": ["first", "second", "third"]}' http://localhost:8080/createAd
```
Метод получения списка объявлений
```
curl --header "Content-Type: application/json" --request GET --data '{"page_num":"1"}' http://localhost:8080/getListOfAd
curl --header "Content-Type: application/json" --request GET --data '{"page_num":"1"}' http://localhost:8080/getListOfAd?date_sort=1
curl --header "Content-Type: application/json" --request GET --data '{"page_num":"1"}' http://localhost:8080/getListOfAd?date_sort=1&price_sort=0
```
Метод получения конкретного объявления
```
curl --header "Content-Type: application/json" --request GET --data '{"ad_id":"1"}' http://localhost:8080/getAdById
curl --header "Content-Type: application/json" --request GET --data '{"ad_id":"1"}' http://localhost:8080/getAdById?fields=1
```
