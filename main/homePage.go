package main

type homePage interface {
	Hello() string
}

type helloWorldHomePage struct {
	helloMessage string
}

func NewHomePage() *helloWorldHomePage {
	h := new(helloWorldHomePage)
	h.helloMessage = "Hello World"
	return h
}

func (h helloWorldHomePage) Hello() string {
	return h.helloMessage
}
