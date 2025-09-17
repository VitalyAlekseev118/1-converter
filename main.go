package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {

	for {
		inputSourceTargetCurrency, err := inputSourceTargetCurrency()
		if err != nil {
			fmt.Println(err)
			continue
		}

		sum, err := inputSum()
		if err != nil {
			fmt.Println(err)
			continue
		}

		mapConverter, err := calculationResult(sum, inputSourceTargetCurrency)
		if err != nil {
			fmt.Println(err)
			continue

		}

		for key, value := range *mapConverter {
			fmt.Printf("Результат расчета %s : %.2f\n", key, value)
		}

		fmt.Print("Хотите продолжить (y/n)? ")
		var continueInput string
		fmt.Scan(&continueInput)
		if continueInput != "y" {
			fmt.Println("Всего хорошего")
			break
		}

	}
}

func inputSourceTargetCurrency() (string, error) {
	for {
		fmt.Print("Введите пару валют для конвертации (USD/EUR/RUB) используя /: ")
		var currency string
		fmt.Scan(&currency)

		currency = strings.ToUpper(currency)
		if isValidCurrency(currency) {
			fmt.Println("Приянто")
			return currency, nil
		} else {
			return " ", errors.New("Ошибка! Попробуйте снова")
		}
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

func isValidCurrency(currency string) bool {
	if currency == "USD/EUR" || currency == "USD/RUB" || currency == "EUR/USD" || currency == "EUR/RUB" ||
		currency == "RUB/EUR" || currency == "RUB/USD" {
		return true
	}
	fmt.Println("Нет такой пары")
	return false
}

func calculationResult(sum float64, inputSourceTargetCurrency string) (*map[string]float64, error) {

	m := make(map[string]float64)

	switch inputSourceTargetCurrency {
	case "EUR/USD":
		m[inputSourceTargetCurrency] = sum * 1.17
	case "USD/EUR":
		m[inputSourceTargetCurrency] = sum * 0.86
	case "EUR/RUB":
		m[inputSourceTargetCurrency] = sum * 94.05
	case "RUB/EUR":
		m[inputSourceTargetCurrency] = sum * 0.0106
	case "USD/RUB":
		m[inputSourceTargetCurrency] = sum * 79.65
	case "RUB/USD":
		m[inputSourceTargetCurrency] = sum * 0.013
	default:
		return nil, errors.New("Конвертация между указанными валютами не поддерживается")
	}

	return &m, nil
}
