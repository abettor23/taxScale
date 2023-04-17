package main

import (
	"fmt"
)

func main() {
	/*
	*Если прибыль до 10000 - ставка налога равна 13%
	*Начиная с 10000 и до 50000 - 20%
	*А начиная с 50000 - 30%
	 */

	fmt.Println("Введите размер прибыли:")
	var profit int
	fmt.Scan(&profit)

	if profit >= 50000 {
		tax := profit * 30 / 100
		fmt.Println("Размер налога (30%) равен:", tax)
	} else if profit >= 10000 {
		tax := profit * 20 / 100
		fmt.Println("Размер налога (20%) равен:", tax)
	} else if profit > 0 {
		tax := profit * 13 / 100
		fmt.Println("Размер налога (13%) равен:", tax)
	} else if profit == 0 {
		fmt.Println("Правда? НИчего не заработали? А если найдем?")
	} else {
		fmt.Println("Ожидалось не отрицательное число")
	}
}
