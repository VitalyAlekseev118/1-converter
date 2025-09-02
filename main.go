package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {

	var sum float64
	for {
		choiceValuta, choiceValutaResult, err := menuUser()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Print("Введите сумму для расчета: ")
		_, err = fmt.Scan(&sum)
		if err != nil {
			fmt.Println("Произошла ошибка при вводе суммы: ", err)
			continue
		}

		calculate, err := calculationResult(sum, choiceValuta, choiceValutaResult)
		if err != nil {
			fmt.Println("Ошибка расчета: ", err)
			continue
		}

		fmt.Printf("Результат расчета %.2f %s\n", calculate, strings.ToUpper(choiceValutaResult))

		fmt.Print("Хотите продолжить (y/n)? ")
		var continueInput string
		fmt.Scan(&continueInput)
		if continueInput != "y" {
			fmt.Println("Всего хорошего")
			break
		}
	}
}

func menuUser() (string, string, error) {
	var choiceValuta string
	var choiceValutaResult string

	fmt.Println("__Меню__")
	fmt.Print("Выберете валюту из списка: USD/EUR/RUB - ")
	fmt.Scan(&choiceValuta)

	fmt.Print("Выберете валюту для расчета: ")
	fmt.Scan(&choiceValutaResult)

	choiceValuta = strings.ToUpper(choiceValuta)
	choiceValutaResult = strings.ToUpper(choiceValutaResult)

	if !isValidCurrency(choiceValuta) && !isValidCurrency(choiceValutaResult) {
		return " ", " ", errors.New("Введена не корректная валюта, просьба попробовать еще раз")
	}

	if choiceValuta == choiceValutaResult {
		return " ", " ", errors.New("Ошибка, введены одинаковые валюты")
	}

	return choiceValuta, choiceValutaResult, nil
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
