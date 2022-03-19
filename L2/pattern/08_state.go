package main

/*
	Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего состояния.
*/

import "fmt"

type State interface {
	Render(*Page)
}

type Page struct {
	State
}

func (p *Page) ChangeState(s State) {
	p.State = s
}

func (p *Page) Publish() {
	p.State.Render(p)
}

type DraftState struct{}

func (ds *DraftState) Render(p *Page) {
	fmt.Println("Draft")
	p.ChangeState(new(ModerationState))
}

type ModerationState struct{}

func (ms *ModerationState) Render(p *Page) {
	fmt.Println("Moderation")
	p.ChangeState(new(PublicationState))
}

type PublicationState struct{}

func (ps *PublicationState) Render(p *Page) {
	fmt.Println("Publication")
}

func main() {
	page := &Page{State: new(DraftState)}

	page.Publish()
	page.Publish()
	page.Publish()
}

// Draft
// Moderation
// Publication
