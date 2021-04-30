package main

type HomePage interface {
	Hello() string
}

type HelloWorldHomePage struct {
	helloMessage string
}

func NewHomePage() *HelloWorldHomePage {
	h := new(HelloWorldHomePage)
	h.helloMessage = "Hello World"
	return h
}

func (h HelloWorldHomePage) Hello() string {
	return h.helloMessage
}
