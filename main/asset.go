package main

type Asset struct {
	id          int
	isin        string
	description string
	quantity    float32
	price       float32
}

func (a Asset) Value() float32 {
	return a.quantity * a.price
}

func NewAdalphaAssets() []Asset {
	var a [6]Asset

	a[0].id = 0
	a[0].isin = "IE00B52L4369"
	a[0].description = "BlackRock Institutional Cash Series Sterling Liquidity Agency Inc"
	a[0].quantity = 44
	a[0].price = 1

	a[1].id = 1
	a[1].isin = "GB00BQ1YHQ70"
	a[1].description = "Threadneedle UK Property Authorised Investment Net GBP 1 Acc"
	a[1].quantity = 22
	a[1].price = 2

	a[2].id = 2
	a[2].isin = "GB00B3X7QG63"
	a[2].description = "Vanguard FTSE U.K. All Share Index Unit Trust Accumulation"
	a[2].quantity = 44
	a[2].price = 0.5

	a[3].id = 3
	a[3].isin = "GB00BG0QP828"
	a[3].description = "Legal & General Japan Index Trust C Class Accumulation"
	a[3].quantity = 11
	a[3].price = 1

	a[4].id = 4
	a[4].isin = "GB00BPN5P238"
	a[4].description = "Vanguard US Equity Index Institutional Plus GBP Accumulation"
	a[4].quantity = 44
	a[4].price = 0.75

	a[5].id = 5
	a[5].isin = "IE00B1S74Q32"
	a[5].description = "Vanguard U.K. Investment Grade Bond Index Fund GBP Accumulation"
	a[5].quantity = 264
	a[5].price = 0.25

	return a[0:6]
}
