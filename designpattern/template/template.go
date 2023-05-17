package template

import (
	"fmt"
	"sort"
)

type Coffee struct{}

func (c *Coffee) prepareRecipe() {
	c.boilWater()
	c.brewCoffeeGrinds()
	c.pourInCup()
	c.addSugarAndMilk()
}

func (c *Coffee) boilWater() {
	fmt.Println("Boiling water")
}
func (c *Coffee) brewCoffeeGrinds() {
	fmt.Println("Dripping Coffee through filter")
}

func (c *Coffee) pourInCup() {
	fmt.Println("Pouring into cup")
}

func (c *Coffee) addSugarAndMilk() {
	fmt.Println("Adding Sugar and Milk")
}

type Tea struct{}

func (t *Tea) prepareRecipe() {
	t.boilWater()
	t.steepTeaBag()
	t.pourInCup()
	t.addLemon()
}

func (t *Tea) boilWater() {
	fmt.Println("Boiling water")
}

func (t *Tea) steepTeaBag() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) pourInCup() {
	fmt.Println("Pouring into cup")
}

func (t *Tea) addLemon() {
	fmt.Println("Adding Lemon")
}

type Beverage interface {
	brew()
	addCondiments()
}

type CaffeineBeverage struct {
	Beverage
}

func (c *CaffeineBeverage) prepare() {
	c.boilWater()
	c.Beverage.brew()
	c.pourInCup()
	c.Beverage.addCondiments()
}
func (c *CaffeineBeverage) boilWater() {
	fmt.Println("Boiling water")
}
func (c *CaffeineBeverage) pourInCup() {
	fmt.Println("Pouring into cup")
}
func (c *CaffeineBeverage) wantIt() bool {
	return true
}

type Coffee2 struct{}

func (c *Coffee2) brew() {
	fmt.Println("Dripping Coffee through filter")
}
func (c *Coffee2) addCondiments() {
	fmt.Println("Adding Sugar and Milk")
}

type Tea2 struct{}

func (t *Tea2) brew() {
	fmt.Println("Steeping the tea")
}
func (t *Tea2) addCondiments() {
	fmt.Println("Adding Lemon")
}

type Duck struct {
	name   string
	weight int
}

func Start() {
	coffee := &CaffeineBeverage{&Coffee2{}}
	tea := &CaffeineBeverage{&Tea2{}}

	fmt.Println("Making coffee...")
	coffee.prepare()

	fmt.Println("\nMaking tea...")
	tea.prepare()

	list := []Duck{
		{"Daffy", 8},
		{"Dewey", 2},
		{"Howard", 7},
		{"Louie", 2},
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].weight < list[j].weight
	})

	fmt.Println(list)
}
