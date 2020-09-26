package AvitoWatcher

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//Ключевые слова для поиска на странице
const priceStartKeyWord = "\"dynx_price\":"
const priceEndKeyWord = ",\"dynx_category\":"

const idStartKeyWord = "\"itemID\":"
const idEndKeyWord = ",\"vertical\":"

//Выделяет строку между ключевыми словами
func findBetween(source string, start string, end string) (string, error) {
	l := strings.Index(source, start)
	r := strings.Index(source, end)

	if r <= l {
		return "", errors.New("can't find keyword")
	}
	return source[l+len(start) : r], nil
}

//Получает цену, ID объявления (Разные ссылки могут вести на один товар)
func GetPrice(address string) (int, int64, error) {

	response, err := http.Get(address)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	priceString, err := findBetween(bodyString, priceStartKeyWord, priceEndKeyWord)
	if err != nil {
		return -1, 0, errors.New("wrong url")
	}

	price, err := strconv.Atoi(priceString)
	if err != nil{
		return -1, 0, errors.New("can't parse price")
	}

	idString, err := findBetween(bodyString, idStartKeyWord, idEndKeyWord)

	if err != nil {
		return -1, 0, errors.New("wrong url")
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return -1, 0, errors.New("can't parse id")
	}

	return price, id, nil
}