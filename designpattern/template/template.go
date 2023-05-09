package template

import "fmt"

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
