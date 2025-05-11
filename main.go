package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var apiKey string = "4bf24a447dfc553655f37dce" //CACHE !!!

type ExchangeResponse struct {
	BaseCode       string  `json:"base_code"`
	TargetCode     string  `json:"target_code"`
	ConversionRate float64 `json:"conversion_rate"`
}

type CurrencyCode struct {
	Result         string     `json:"result"`
	SupportedCodes [][]string `json:"supported_codes"`
}

func main() {
	fmt.Println("Liste des devises disponibles :", CurrencyCodeList())

	entryCurrency := EntryCurrency()
	targetCurrency := TargetCurrency()
	amount := Amount()

	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/pair/%s/%s", apiKey, entryCurrency, targetCurrency)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données :", err)
		return
	}
	defer resp.Body.Close()

	var data ExchangeResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON :", err)
		return
	}

	fmt.Printf("1 %s = %.2f %s\n", data.BaseCode, data.ConversionRate, data.TargetCode)

	convertedAmount := amount * data.ConversionRate

	fmt.Printf("%.2f %s = %.2f %s\n", amount, entryCurrency, convertedAmount, targetCurrency)

}

func EntryCurrency() string {
	fmt.Print("Entrez la devise de départ : ")
	var currency string
	fmt.Scanln(&currency)
	return strings.TrimSpace(strings.ToUpper(currency))
}

func TargetCurrency() string {
	fmt.Print("Entrez la devise cible : ")
	var currency string
	fmt.Scanln(&currency)
	return strings.TrimSpace(strings.ToUpper(currency))
}

func Amount() float64 {
	fmt.Print("Entrez le montant à convertir : ")
	var amount float64
	fmt.Scanln(&amount)
	return amount
}

func CurrencyCodeList() []string {
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/codes", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données :", err)
		return []string{}
	}
	defer resp.Body.Close()

	var data CurrencyCode
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON :", err)
		return []string{}
	}

	var currencyCodes []string
	for _, codePair := range data.SupportedCodes {
		currencyCodes = append(currencyCodes, codePair[0])
	}

	return currencyCodes
}
