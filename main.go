package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {

	for {
		inputSourceCurrency := inputSourceCurrency()

		sum, err := inputSum()
		if err != nil {
			fmt.Println(err)
			continue
		}
		targetCurrency := targetCurrency()

		err = verificationInputValute(inputSourceCurrency, targetCurrency)
		if err != nil {
			fmt.Println(err)
			continue
		}

		calculate, err := calculationResult(sum, inputSourceCurrency, targetCurrency)
		if err != nil {
			fmt.Println(err)
			continue

		}

		fmt.Printf("Результат расчета %.2f %s\n", calculate, strings.ToUpper(targetCurrency))

		fmt.Print("Хотите продолжить (y/n)? ")
		var continueInput string
		fmt.Scan(&continueInput)
		if continueInput != "y" {
			fmt.Println("Всего хорошего")
			break
		}

	}
}

func inputSourceCurrency() string {
	for {
		fmt.Print("Введите исходную валюту (USD/EUR/RUB): ")
		var currency string
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)
		if isValidCurrency(currency) {
			return currency
		}
		fmt.Println("Ошибка! Попробуйте снова")
	}
}

func inputSum() (float64, error) {
	for {
		var sum float64
		fmt.Print("Введите сумму для расчета: ")
		_, err := fmt.Scan(&sum)

		if err != nil {
			fmt.Println("Ошибка! Введено не число. Попробуйте снова.")
			continue // Повторяем ввод суммы
		}
		if sum <= 0 {
			return 0.0, errors.New("Введено отрицательное число, работает только с положительными")
		}
		return sum, nil
	}

}

func targetCurrency() string {
	for {
		fmt.Print("Введите валюту для конвертации: ")
		var currencyTarget string
		fmt.Scan(&currencyTarget)
		currencyTarget = strings.ToUpper(currencyTarget)

		if isValidCurrency(currencyTarget) {
			return currencyTarget
		}
		fmt.Println("Ошибка! Попробуйте сноваfffff")
	}
}

func verificationInputValute(choiceValuta, choiceValutaResult string) error {

	if !isValidCurrency(choiceValuta) || !isValidCurrency(choiceValutaResult) {
		return errors.New("Введена не корректная валюта, просьба попробовать еще раз")
	}

	if choiceValuta == choiceValutaResult {
		return errors.New("Ошибка, введены одинаковые валюты")
	}

	return nil
}

func isValidCurrency(currency string) bool {
	if currency == "USD" || currency == "EUR" || currency == "RUB" {
		return true
	}
	return false
}

func calculationResult(sum float64, inputValuta string, inputConvertValuta string) (float64, error) {

	var calculate float64

	switch {
	case inputValuta == "EUR" && inputConvertValuta == "USD":
		calculate = sum * 1.17
	case inputValuta == "USD" && inputConvertValuta == "EUR":
		calculate = sum * 0.86
	case inputValuta == "EUR" && inputConvertValuta == "RUB":
		calculate = sum * 94.05
	case inputValuta == "RUB" && inputConvertValuta == "EUR":
		calculate = sum * 0.0106
	case inputValuta == "USD" && inputConvertValuta == "RUB":
		calculate = sum * 79.65
	case inputValuta == "RUB" && inputConvertValuta == "USD":
		calculate = sum * 0.013
	default:
		return 0.0, errors.New("Конвертация между указанными валютами не поддерживается")
	}

	return calculate, nil
}
