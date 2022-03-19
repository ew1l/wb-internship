package main

/*
	Цепочка обязанностей — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков.
*/

import "fmt"

type Patient struct {
	RegistrationDone bool
	DoctorDone       bool
	MedicineDone     bool
	PaymentDone      bool
}

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}

type Reception struct {
	Next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.Next.Execute(p)
	}

	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	r.Next.Execute(p)
}

func (r *Reception) SetNext(next Department) {
	r.Next = next
}

type Doctor struct {
	Next Department
}

func (d *Doctor) Execute(p *Patient) {
	if p.DoctorDone {
		fmt.Println("Doctor checkup already done")
		d.Next.Execute(p)
	}

	fmt.Println("Doctor checking patient")
	p.DoctorDone = true
	d.Next.Execute(p)
}

func (d *Doctor) SetNext(next Department) {
	d.Next = next
}

type Medical struct {
	Next Department
}

func (m *Medical) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.Next.Execute(p)
	}

	fmt.Println("Medical giving medicine to patient")
	p.MedicineDone = true
	m.Next.Execute(p)
}

func (m *Medical) SetNext(next Department) {
	m.Next = next
}

type Cashier struct {
	Next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("Payment done")
	}

	fmt.Println("Cashier getting money from patient")
}

func (c *Cashier) SetNext(next Department) {
	c.Next = next
}

func main() {
	patient := new(Patient)

	reception := new(Reception)
	doctor := new(Doctor)
	medical := new(Medical)
	cashier := new(Cashier)

	reception.SetNext(doctor)
	doctor.SetNext(medical)
	medical.SetNext(cashier)

	reception.Execute(patient)
}

// Reception registering patient
// Doctor checking patient
// Medical giving medicine to patient
// Cashier getting money from patient
