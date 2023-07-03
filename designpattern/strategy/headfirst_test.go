package strategy

import (
	"fmt"
	"testing"
)

func Test_01(t *testing.T) {
	var d Duck
	d = Duck(NewMallardDuck2().Duck)
	d.performFly()
	d.performQuack()
}

func Test_02(t *testing.T) {
	tt := NewTextProcessor(&MarkdownListStrategy{})
	tt.AppendList([]string{"park", "gui", "woo"})

	fmt.Println(tt)

	tt.Reset()

	tt.SetOutputFormat(Html)
	tt.AppendList([]string{"park", "gui", "woo"})
	fmt.Println(tt)
}
