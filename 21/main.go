package main

import "fmt"

// евро вилка
type EuponeanPlug struct{}

// сообщаем об успешном подключении
func (e *EuponeanPlug) Connect() {
	fmt.Println("Евро вилка подключена к итальянской розентке")
}

// интерфейс неподходящей розетки
type PowerSocket interface {
	Connect()
}

// переходник для вилки
type PlugAdapter struct {
	EuponeanPlug *EuponeanPlug
}

//пдключение к розетке, используя старую вилку
func (p *PlugAdapter) ConnectWithAdapter() {
	p.EuponeanPlug.Connect()
}

func main() {
	euponeanPlug := &EuponeanPlug{}
	adapter := &PlugAdapter{EuponeanPlug: euponeanPlug}
	adapter.ConnectWithAdapter()
}
