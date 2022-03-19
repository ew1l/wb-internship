package main

/*
	Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.
*/

import "fmt"

type IWeapon interface {
	GetName() string
	GetPower() int
}

type Weapon struct {
	Name  string
	Power int
}

func (g *Weapon) GetName() string {
	return g.Name
}

func (g *Weapon) GetPower() int {
	return g.Power
}

type AK47 struct {
	Weapon
}

func NewAK47() IWeapon {
	return &AK47{
		Weapon: Weapon{
			Name:  "AK-47",
			Power: 3,
		},
	}
}

type Glock18 struct {
	Weapon
}

func NewGlock18() IWeapon {
	return &Glock18{
		Weapon: Weapon{
			Name:  "Glock 18",
			Power: 1,
		},
	}
}

func GetWeapon(w string) (IWeapon, error) {
	switch w {
	case "AK-47":
		return NewAK47(), nil
	case "Glock 18":
		return NewGlock18(), nil
	default:
		return nil, fmt.Errorf("Wrong weapon!")
	}
}

func Details(weapon IWeapon) string {
	return fmt.Sprintf("Gun: %s\nPower: %d\n", weapon.GetName(), weapon.GetPower())
}

func main() {
	weapons := []string{"AK-47", "Glock 18", "M16"}

	for _, w := range weapons {
		weapon, err := GetWeapon(w)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(Details(weapon))
	}
}

// Gun: AK-47
// Power: 3
//
// Gun: Glock 18
// Power: 1
//
// Wrong weapon!
