package main

import (
	"fmt"
	"math/rand"
)

type Bot struct {
	name    string
	money   int
	Bidding []*Commodity
	Owner   []*Commodity
}

func (bot *Bot) toBidding(commodity *Commodity, cs []*Commodity) bool {
	c := 0
	for i := 0; i < len(bot.Owner); i++ {
		c += bot.Owner[i].money
	}
	for i := 0; i < len(cs); i++ {
		if cs[i].TmpOwner == bot {
			c += cs[i].money
		}
	}
	yuer := bot.money - c
	if yuer <= 0 || yuer <= commodity.money {
		fmt.Println(bot.name, "余额不足")
		return false
	}

	jiajia := rand.Intn(yuer-commodity.money) + 1
	fmt.Printf("%s的余额为%d,加价%d\n", bot.name, yuer, jiajia)
	commodity.TmpOwner = bot
	commodity.money += jiajia
	commodity.ask = 0

	fmt.Printf("%s加价%d成功,%s最新价格%d\n", bot.name, jiajia, commodity.name, commodity.money)
	return true
}

type Commodity struct {
	name     string
	money    int
	Owner    *Bot
	TmpOwner *Bot
	ask      int
}

func main() {

	bot1 := &Bot{name: "[机器人1]", money: 100, Bidding: make([]*Commodity, 0, 0), Owner: make([]*Commodity, 0, 0)}
	bot2 := &Bot{name: "[机器人2]", money: 100, Bidding: make([]*Commodity, 0, 0), Owner: make([]*Commodity, 0, 0)}
	bot3 := &Bot{name: "[机器人3]", money: 100, Bidding: make([]*Commodity, 0, 0), Owner: make([]*Commodity, 0, 0)}

	bots := []*Bot{bot1, bot2, bot3}
	c1 := &Commodity{"[豆沙包1]", 1, nil, nil, 0}
	c2 := &Commodity{"[豆沙包2]", 1, nil, nil, 0}
	c3 := &Commodity{"[豆沙包3]", 1, nil, nil, 0}
	cs := []*Commodity{c1, c2, c3}

	for t := 0; t < 6; t++ {
		loop := true
		for i := 0; i < 3; i++ { //商品
			fmt.Println("======================")
			if cs[i].Owner == nil {
				fmt.Println("竞拍商品：", cs[i].name)
				for j := 0; j < 3; j++ { //机器人
					if cs[i].TmpOwner == bots[j] {
						continue
					}
					if !bots[j].toBidding(cs[i], cs) {
						cs[i].ask += 1
					}
				}
				if cs[i].ask >= 2 {
					cs[i].Owner = cs[i].TmpOwner
					cs[i].TmpOwner = nil
					cs[i].Owner.Owner = append(cs[i].Owner.Owner, cs[i])
					fmt.Printf("❤️%s已被%s拍得，价格为:%d\n", cs[i].name, cs[i].Owner.name, cs[i].money)
				}
				loop = false
			}
		}

		if loop {
			break
		}
	}
	fmt.Println("拍卖结束")
}
