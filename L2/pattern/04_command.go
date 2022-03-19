package main

/*
	Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты, позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.
*/

import "fmt"

type Command interface {
	Execute()
}

type Button struct {
	Command
}

func (b *Button) Press() {
	b.Command.Execute()
}

type Device interface {
	On()
	Off()
}

type OnCommand struct {
	Device
}

func (oc *OnCommand) Execute() {
	oc.Device.On()
}

type OffCommand struct {
	Device
}

func (oc *OffCommand) Execute() {
	oc.Device.Off()
}

type TV struct {
	IsRunning bool
}

func (t *TV) On() {
	t.IsRunning = true
	fmt.Println("TV on")
}

func (t *TV) Off() {
	t.IsRunning = false
	fmt.Println("TV off")
}

func main() {
	TV := new(TV)

	onCommand := &OnCommand{
		Device: TV,
	}

	offCommand := &OffCommand{
		Device: TV,
	}

	onButtonTV := &Button{
		Command: onCommand,
	}
	onButtonTV.Press()

	offButtonTV := &Button{
		Command: offCommand,
	}
	offButtonTV.Press()

	onButtonRemote := &Button{
		Command: onCommand,
	}
	onButtonRemote.Press()

	offButtonRemote := &Button{
		Command: offCommand,
	}
	offButtonRemote.Press()
}

// TV on
// TV off
// TV on
// TV off
