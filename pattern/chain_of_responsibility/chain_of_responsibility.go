package main

import "fmt"

type department interface {
	execute(*patient)
	setNext(department)
}

type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Регистрация пациента завершена")
		r.next.execute(p)
		return
	}
	fmt.Println("Регистрация пациента")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Осмотр у врача проведен")
		d.next.execute(p)
		return
	}
	fmt.Println("Осмотр у врача")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Лекарство выдано")
		m.next.execute(p)
		return
	}
	fmt.Println("Выдача лекарства")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Оплата произведена")
	}
	fmt.Println("Произведение оплаты")
}

func (c *cashier) setNext(next department) {
	c.next = next
}

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func main() {
	cashier := &cashier{}

	medical := &medical{}
	medical.setNext(cashier)

	doctor := &doctor{}
	doctor.setNext(medical)

	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "Ivan"}
	reception.execute(patient)
}
