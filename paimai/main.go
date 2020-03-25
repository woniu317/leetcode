package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var lockItem = [3]sync.Mutex{}
var lockMoney = sync.Mutex{}

type saleMoney struct {
	money      int32
	consumerId int32
}

var money = [3]int32{100, 100, 100}
var sale = [3]saleMoney{{0, -1}, {0, -1}, {0, -1}}

func main() {
	go consumer(0)
	go consumer(1)
	go consumer(2)
	for {
		if saleOver() {
			break
		}
	}
}

func saleOver() bool {
	lastMoney := [3]int32{}
	sameTime := 0
	for {
		time.Sleep(5 * time.Microsecond)
		i := 0
		for ; i < 3; i++ {
			if lastMoney[i] != sale[i].money {
				sameTime = 0
				break
			}
		}
		if i == 3 {
			sameTime++
		}
		if sameTime == 5 {
			return true
		}
		copyMoney(&lastMoney)
	}
}

func copyMoney(lastMoney *[3]int32) {
	lockMoney.Lock()
	defer lockMoney.Unlock()
	for i := 0; i < 3; i++ {
		(*lastMoney)[i] = sale[i].money
		fmt.Printf("%d-%d,", sale[i].money, sale[i].consumerId)
	}
	fmt.Println()
}

func consumer(id int32) {
	for {
		itemId := rand.Intn(3)
		addMoneyItem(itemId, id)
	}
}

func addMoneyItem(itemId int, consumerId int32) {
	lockItem[itemId].Lock()
	defer lockItem[itemId].Unlock()
	if sale[itemId].consumerId == consumerId {
		return
	}
	if sale[itemId].money >= money[consumerId] {
		return
	}
	lockMoney.Lock()
	defer lockMoney.Unlock()
	if sale[itemId].money >= money[consumerId] {
		return
	}
	atomic.StoreInt32(&sale[itemId].consumerId, consumerId)
	atomic.AddInt32(&sale[itemId].money, 1)
	atomic.AddInt32(&money[consumerId], -1)
}
