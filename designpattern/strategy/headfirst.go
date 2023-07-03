package strategy

import "fmt"

type MallardDuck struct {
	quack QuackBehavior
	fly   FlyBehavior
}

func (m *MallardDuck) display() {
	fmt.Println("저는 물 오리 입니다.")
}
func NewMallardDuck() *MallardDuck {
	return &MallardDuck{
		&Quack{},
		&FlyWithWings{},
	}
}

type FlyBehavior interface {
	fly()
}
type FlyWithWings struct{}

func (f *FlyWithWings) fly() {
	fmt.Println("나는 난다 날개를 가지고 흐힣?")
}

type FlyNoWay struct{}

func (f *FlyNoWay) fly() {
	fmt.Println("나는 방법이 없다 뚜벅초다 ㅠ")
}

type QuackBehavior interface {
	quack()
}
type Quack struct{}

func (q *Quack) quack() {
	fmt.Println("꽉 운다")
}

type Squeak struct{}

func (s *Squeak) quack() {
	fmt.Println("스ㅏ카스카스캇")
}

type MuteQuack struct{}

func (m *MuteQuack) quack() {
	fmt.Println("울수가 없는걸 주륵 ....")
}

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d *Duck) performQuack() {
	d.quackBehavior.quack()
}
func (d *Duck) performFly() {
	d.flyBehavior.fly()
}

type MallardDuck2 struct {
	Duck
}

func NewMallardDuck2() *MallardDuck2 {
	return &MallardDuck2{
		Duck{&FlyWithWings{}, &Quack{}},
	}
}

func NewDuck(fly FlyBehavior, quack QuackBehavior) *Duck {
	return &Duck{
		fly,
		quack,
	}
}
