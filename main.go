package main

func main() {
	sales := []int{12_000, 8_000, 15_000, 8_000, 10_000}
	for _, s := range sales {
		println(bonus(s))
	}
}

func bonus(s int) int {
	if s >= 10_000 {
		b := (s / 100) * 5
		return b
	}
	return 0
}
