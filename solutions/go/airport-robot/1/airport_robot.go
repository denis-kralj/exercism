package airportrobot

import "fmt"

type Greeter interface {
    LanguageName() string
    Greet(string) string
}

func SayHello(name string, greeter Greeter) string{
    return fmt.Sprintf("I can speak %v: %v", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct {
}

func (target Italian) LanguageName() string {
    return "Italian"
}
func (target Italian) Greet(name string) string {
    return fmt.Sprintf("Ciao %v!", name)
}

type Portuguese struct {
}

func (target Portuguese) LanguageName() string {
    return "Portuguese"
}
func (target Portuguese) Greet(name string) string {
    return fmt.Sprintf("Olá %v!", name)
}

