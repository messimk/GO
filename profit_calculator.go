package main

import "fmt"

func ss() {
	// var taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("expenses: ")
	taxRate := getUserInput("TaxRate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	fmt.Println("EBT:", ebt)
	fmt.Println("Profit after tax:", profit)
	fmt.Println("EBT/Profit ratio:", ratio)

	// formattedFv := fmt.Sprintf("Future value: %.1f\n", revenue)
	// formattedRFv := fmt.Sprintf("Future value (adjusted for Inflation)\n %.1f\n:", profit)

	// fmt.Print(formattedFv, formattedRFv)
}
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio

}
func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
