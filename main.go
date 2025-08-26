package main

import "fmt"

func main() {
	const usdToEur = 0.86
	const usdToRub = 80.45

	//Рассчитать EUR в RUB на основании первых двух переменных-констант

	const eurToUsd = 1 / usdToEur
	const eurToRub = usdToRub * eurToUsd
	fmt.Printf("Один евро в рублях составит: %.2f руб.", eurToRub)

}
