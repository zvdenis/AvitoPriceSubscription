package AvitoWatcher

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type SubscriptionManager struct {
	DB *sql.DB
}

type Item struct {
	Id    int64
	Url   string
	Price int
}

// Добавляет информацию о подписке в БД
func (subManager SubscriptionManager) addSubscription(subscription Subscription, price int, id int64) {
	_, err := subManager.DB.Exec("insert items(id, price, url) values(?, ?, ?)", id, price, subscription.Url)

	if strings.Contains(err.Error(), "Duplicate") {
		fmt.Println("Duplicate item")
	} else if err != nil {
		fmt.Println(err)
	}

	_, err = subManager.DB.Exec("insert subscriptions(item_id, email) values(?, ?)", id, subscription.Email)
	if strings.Contains(err.Error(), "Duplicate") {
		fmt.Println("Duplicate subscription")
	} else if err != nil {
		fmt.Print(err)
	}

}

//Получает цену объявления
func (subManager SubscriptionManager) GetPrice(item Item) int {
	price, _, err := GetPrice(item.Url)

	if err != nil {
		fmt.Println("failed update price")
		return item.Price
	}

	return price
}

//Обновляет цену в БД
func (subManager SubscriptionManager) UpdatePrice(id int64, newPrice int) {
	_, err := subManager.DB.Exec("update items set price = ? where id = ?", newPrice, id)
	if err != nil {
		panic(err)
	}
}

//Создает рассылку и сообщение
func (subManager SubscriptionManager) MakeMailing(item Item) {
	rows, err := subManager.DB.Query("select email from subscriptions where item_id = ?", item.Id)
	if err != nil {
		panic(err)
	}
	var emails []string
	var email string
	for rows.Next() {
		rows.Scan(&email)
		emails = append(emails, email)
	}
	message := "new price is: \r\n" + strconv.Itoa(item.Price) + " \r\n Link: \r\n" + item.Url
	SendMail(emails, message)
}

//Возвращает список отслеживаемых предметов
func (subManager SubscriptionManager) GetItemsList() []Item {
	var items []Item

	rows, err := subManager.DB.Query("select * from items")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var item Item
		err := rows.Scan(&item.Id, &item.Price, &item.Url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		items = append(items, item)
	}
	return items
}
