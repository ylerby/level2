package main

import "fmt"

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(_ *Square) {
	fmt.Println("Вычисление площади квадрата")
}

func (a *AreaCalculator) visitForCircle(_ *Circle) {
	fmt.Println("Вычисление площади окружности")
}
func (a *AreaCalculator) visitForRectangle(_ *Rectangle) {
	fmt.Println("Вычисление площади прямоугольника")
}

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(_ *Square) {
	fmt.Println("Вычисление координат средней точки для квадрата")
}

func (a *MiddleCoordinates) visitForCircle(_ *Circle) {
	fmt.Println("Вычисление координат средней точки окружности")
}
func (a *MiddleCoordinates) visitForRectangle(_ *Rectangle) {
	fmt.Println("Вычисление координат средней точки прямоугольника")
}

type _ interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Квадрат"
}

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Окружность"
}

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForRectangle(t)
}

func (t *Rectangle) getType() string {
	return "Прямоугольник"
}

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
