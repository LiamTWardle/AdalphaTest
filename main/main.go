package main

func main() {
	var assets = NewAdalphaAssets()
	var portfolio = NewAdalphaPortfolio(assets)
	var server = NewServer(portfolio)
	server.Listen()
}
