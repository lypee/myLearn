package main

import "log"

type StrategyType int

const (
	StrategyTypeInvalid StrategyType = iota
	StrategyTypeNormal
	StrategyTypeRebate
	StrategyTypeDiscount
)

var (
	RebateCondition = 1
)

func init() {

}

var (
	strategyMap = map[StrategyType]CashContext{
		StrategyTypeNormal:   CashNormal{},
		StrategyTypeRebate:   CashRebate{},
		StrategyTypeDiscount: CashDiscount{},
	}
)

type CashStrategy struct {
	cs CashContext
}
type CashContext interface {
	AcceptCash(price float64) float64
}

type CashNormal struct{}

func (this CashNormal) AcceptCash(price float64) float64 {
	if price < 0.0 {
		return 0.0
	}
	return price
}

type CashRebate struct {
	Condition float64
	Rebate    float64
}

func (this CashRebate) AcceptCash(price float64) float64 {
	if price > this.Condition {
		price -= this.Rebate
	}
	if price < 0.0 {
		return 0.0
	}

	return price
}

type CashDiscount struct {
	Condition float64
	Discount  float64
}

func (this CashDiscount) AcceptCash(price float64) float64 {
	if this.Discount <= 0 {
		this.Discount = 10.0
	}
	if price > this.Condition {
		price /= this.Discount
	}
	if price < 0.0 {
		return 0.0
	}
	return price
}

func main() {
	price := 100.0
	rebateStrategy := strategyMap[StrategyTypeRebate]
	price = rebateStrategy.AcceptCash(price)
	log.Println(price)
	strategyTypeDiscount := strategyMap[StrategyTypeDiscount]
	price = strategyTypeDiscount.AcceptCash(price)
	log.Println(price)
}
