package main

func main() {}

func bonus(sales []int) int {
	minSaleValue := 10_000
	for _, sale := range sales {
		if sale >= minSaleValue {
			percent := 100
			bonusValue := 5
			managerBonus := (sale / percent) * bonusValue
			return managerBonus
		}
	}
	return 0
}
