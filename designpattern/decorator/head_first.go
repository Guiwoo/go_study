package decorator

import "fmt"

// 모든걸 서브클래스 해결 => 상속으로 한다 는 뉘앙스
// 데코레이터는 자신이 장식하고 있는 객체에게 어떤 행동을 위임하는 일 말고도 추가 작업을 수행할 수 있다.
//객체에 추가요소를 동적으로 더할 수 있다. 데코레이터를 사용하면 서브클래스를 만들때 보다 훨씬유연하게 기능을 확장할수 있다.

type beverage interface {
	cost() float32
	getDescription() string
}

type HouseBlend struct {
	description string
}

func (h *HouseBlend) cost() float32 {
	return 0.89
}
func (h *HouseBlend) getDescription() string {
	return h.description
}

type DarkRost struct {
	description string
}

func (d *DarkRost) cost() float32 {
	return 1.32
}
func (d *DarkRost) getDescription() string {
	return d.description
}

type Milk struct {
	b beverage
}

func (m *Milk) cost() float32 {
	return m.b.cost() + 12.3
}
func (m *Milk) getDescription() string {
	return m.b.getDescription() + " milk"
}

type Whip struct {
	b beverage
}

func (w *Whip) cost() float32 {
	return w.b.cost() + 15.5
}
func (w *Whip) getDescription() string {
	return w.b.getDescription() + " whip"
}

func StartHead() {
	//case 1
	a := &HouseBlend{"house coffee"}
	b := &Milk{a}
	c := &Whip{b}

	fmt.Println(c.cost(), c.getDescription())
	//case2
	var B beverage
	B = &HouseBlend{"house"}
	B = &Milk{B}
	B = &Whip{B}

	fmt.Println(B.cost())
}
