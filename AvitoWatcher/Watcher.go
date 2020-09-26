package AvitoWatcher

import (
	"fmt"
	"time"
)

type Watcher struct {
	SubManager SubscriptionManager
}

//Проверяет объявления из списка на изменение цены
func (watcher Watcher) WatchAvito() {
	ticker := time.NewTicker(time.Minute)
	for t := range ticker.C {
		fmt.Println("Scan at:  ", t)
		items  := watcher.SubManager.GetItemsList()

		for _, el := range items {
			newPrice, _, err := GetPrice(el.Url)
			if err != nil{
				fmt.Println("can't parse item with id: ", el.Id)
				continue
			}

			if newPrice != el.Price {
				watcher.SubManager.UpdatePrice(el.Id, newPrice)
				el.Price = newPrice
				watcher.SubManager.MakeMailing(el)
			}
		}
	}
}
