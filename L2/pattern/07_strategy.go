package main

/*
	Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс.
*/

import "fmt"

type RouteStrategy interface {
	BuildRoute(string, string)
}

type RoadStrategy struct{}

func (rs *RoadStrategy) BuildRoute(A string, B string) {
	fmt.Println("Road route")
}

type PublicTransportStrategy struct{}

func (pts *PublicTransportStrategy) BuildRoute(A string, B string) {
	fmt.Println("Public transport route")
}

type WalkingStrategy struct{}

func (ws *WalkingStrategy) BuildRoute(A string, B string) {
	fmt.Println("Walking route")
}

type Navigator struct {
	Route RouteStrategy
}

func (n *Navigator) SetRoute(route RouteStrategy) {
	n.Route = route
}

func (n *Navigator) BuildRoute(A string, B string) {
	n.Route.BuildRoute(A, B)
}

func main() {
	roadRoute := new(RoadStrategy)
	publicTransportRoute := new(PublicTransportStrategy)
	walkingRoute := new(WalkingStrategy)

	navigator := Navigator{}
	navigator.SetRoute(roadRoute)
	navigator.BuildRoute("-29.038175, 134.796176", "-27.042017, 133.443484")

	navigator.SetRoute(publicTransportRoute)
	navigator.BuildRoute("-33.879051, 151.195760", "-33.892874, 151.214213")

	navigator.SetRoute(walkingRoute)
	navigator.BuildRoute("-34.505367, 139.630075", "-34.703138, 139.458966")
}

// Road route
// Public transport route
// Walking route
