package main

func main() {}

func bonus(sale int) int {
	minSaleValue := 10_000
	if sale >= minSaleValue {
		percent := 100
		bonusValue := 5
		managerBonus := (sale / percent) * bonusValue
		return managerBonus
	}
	return 0
}
