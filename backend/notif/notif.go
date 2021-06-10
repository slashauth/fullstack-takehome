package notif

import (
	"sync"

	"github.com/getdebrief/fullstack-takehome/graph/model"
)

var subscribers sync.Map

func AddSubscriber(id string, notifChannel chan *model.PriceUpdate) error {
	existingList, ok := subscribers.Load(id)
	if !ok {
		existingList = make([]chan *model.PriceUpdate, 0)
		subscribers.Store(id, existingList)
	}

	existingList = append(existingList.([]chan *model.PriceUpdate), notifChannel)
	subscribers.Store(id, existingList)

	return nil
}

func RemoveSubscriber(id string, notifChannel chan *model.PriceUpdate) error {
	existingList, ok := subscribers.Load(id)
	if ok {
		newList := make([]chan *model.PriceUpdate, 0)
		for _, elem := range existingList.([]chan *model.PriceUpdate) {
			if elem != notifChannel {
				newList = append(newList, elem)
			}
		}

		subscribers.Store(id, newList)
	}
	return nil
}

func GetSubscribers(id string) []chan *model.PriceUpdate {
	resp, ok := subscribers.Load(id)
	if !ok {
		return []chan *model.PriceUpdate{}
	}

	return resp.([]chan *model.PriceUpdate)
}

func NotifySubscribers(id string, sess model.PriceUpdate) error {
	subs := GetSubscribers(id)
	for _, elem := range subs {
		elem <- &sess
	}
	return nil
}
