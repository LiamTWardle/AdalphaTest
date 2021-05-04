package main

import "fmt"

const (
	buyInstruction    = "BUY"
	investInstruction = "INVEST"
	sellInstruction   = "SELL"
	raiseInstruction  = "RAISE"
)

const (
	assetNameHeader       = "\"asset-name\": "
	instructionTypeHeader = "\"instruction-type\": "
	amountHeader          = "\"amount\": "
	currencyCodeHeader    = "\"currency-code\": "
)

type Portfolio interface {
	Withdraw(float32) string
}

type AdalphaPortfolio struct {
	assets []Asset
}

func NewAdalphaPortfolio(assets []Asset) *AdalphaPortfolio {
	p := new(AdalphaPortfolio)
	p.assets = assets
	return p
}

func (a AdalphaPortfolio) Withdraw(withdrawAmount float32) string {
	var totalPortfolioValue float32
	for _, asset := range a.assets {
		totalPortfolioValue += asset.Value()
	}
	var instructions string
	for _, asset := range a.assets {
		amountToRaise := withdrawAmount * asset.Value() / totalPortfolioValue
		instructions += WriteInstruction(asset.isin, raiseInstruction, amountToRaise, "GBP")
	}
	return instructions
}

func WriteInstruction(isin string, instructionType string, amount float32, currencyCode string) string {
	instruction := "{"
	instruction += assetNameHeader
	instruction += "\"" + isin + "\","
	instruction += instructionTypeHeader
	instruction += "\"" + instructionType + "\","
	instruction += amountHeader
	instruction += fmt.Sprintf("%f", amount) + ","
	if currencyCode != "" {
		instruction += currencyCodeHeader
		instruction += "\"" + currencyCode + "\","
	}
	instruction += "}"
	return instruction
}
