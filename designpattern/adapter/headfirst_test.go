package adapter

import "testing"

func duckTest(duck Duck) {
	duck.quack()
	duck.fly()
}

func Test_01(t *testing.T) {
	var d Duck
	duck := &MallardDuck{}

	w := &WildTurkey{}
	d = &TurkeyAdapter{w}

	duckTest(duck)

	duckTest(d)

}
