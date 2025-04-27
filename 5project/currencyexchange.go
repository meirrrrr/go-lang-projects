package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const apiKey = "API KEY"
const apiUrl = "https://v6.exchangerate-api.com/v6/%s/latest/%s"

type ApiResponse struct {
	ConversionRate map[string]float64 `json:"conversion_rates"`
}

func getExchangeRate(fromCurrency, toCurrency string) (float64, error) {
	url := fmt.Sprintf(apiUrl, apiKey, fromCurrency)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to make a request to the API: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return 0, fmt.Errorf("failed to parse the API response: %v", err)
	}

	if rate, ok := apiResponse.ConversionRate[toCurrency]; ok {
		return rate, nil
	}
	return 0, fmt.Errorf("invalid target currency: %s", toCurrency)
}

func convertCurrency(amount float64, fromCurrency, toCurrency string) (float64, error) {
	rate, err := getExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		return 0, err
	}

	convertedAmount := amount * rate
	return convertedAmount, nil
}

func main() {
	fmt.Print("Enter the amount to convert: ")
	var amountStr string
	_, err := fmt.Scanln(&amountStr)
	if err != nil {
		log.Fatalf("Invalid input: %v\n", err)
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		log.Fatalf("Invalid amount. Please enter a valid number.\n")
	}

	fmt.Print("Enter the source currency code (e.g., USD): ")
	var fromCurrency string
	_, err = fmt.Scanln(&fromCurrency)
	if err != nil || len(fromCurrency) != 3 {
		log.Fatalf("Invalid source currency code. Please enter a 3-letter code.\n")
	}

	fmt.Print("Enter the target currency code (e.g., EUR): ")
	var toCurrency string
	_, err = fmt.Scanln(&toCurrency)
	if err != nil || len(toCurrency) != 3 {
		log.Fatalf("Invalid target currency code. Please enter a 3-letter code.\n")
	}

	convertedAmount, err := convertCurrency(amount, strings.ToUpper(fromCurrency), strings.ToUpper(toCurrency))
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Printf("%.2f %s is equivalent to %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)
}
