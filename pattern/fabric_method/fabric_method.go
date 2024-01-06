package main

import "fmt"

type transport struct {
	name  string
	speed uint
}

func (t *transport) setName(n string) {
	t.name = n
}

func (t *transport) getName() string {
	return t.name
}

func (t *transport) setSpeed(s uint) {
	t.speed = s
}

func (t *transport) getSpeed() uint {
	return t.speed
}

type electricScooter struct {
	transport
}

func newElectricScooter() iTransport {
	return &electricScooter{
		transport: transport{
			name:  "самокат",
			speed: 4,
		},
	}
}

func getTransport(tt string) (iTransport, error) {
	if tt == "самокат" {
		return newElectricScooter(), nil
	}
	if tt == "квадрокоптер" {
		return newQuadcopter(), nil
	}
	return nil, fmt.Errorf("Неверный тип ")
}

type iTransport interface {
	setName(n string)
	getName() string
	setSpeed(s uint)
	getSpeed() uint
}

type quadcopter struct {
	transport
}

func newQuadcopter() iTransport {
	return &quadcopter{
		transport: transport{
			name:  "квадрокоптер",
			speed: 14,
		},
	}
}

func main() {

	scooter, _ := getTransport("самокат")
	quad, _ := getTransport("квадрокоптер")

	fmt.Println(scooter)
	fmt.Println(quad)
}
