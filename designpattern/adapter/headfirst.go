package adapter

import (
	"fmt"
	"github.com/labstack/gommon/log"
)

/**
기존 시스템 인터페이스 vs 업체에서 제공하는 인터페이스 가 다르다...
*/

type Duck interface {
	quack()
	fly()
}

type MallardDuck struct {
}

func (m *MallardDuck) quack() {
	fmt.Println("quakckkkckkckck")
}
func (m *MallardDuck) fly() {
	fmt.Println("오리는 날고있습니다.")
}

type Turkey interface {
	gobble()
	fly()
}
type WildTurkey struct {
}

func (w *WildTurkey) gobble() {
	fmt.Println("goooooooblelbelbel")
}
func (w *WildTurkey) fly() {
	fmt.Println("짤븡날개로 못나는디;;;")
}

type TurkeyAdapter struct {
	t Turkey
}

func (t *TurkeyAdapter) quack() {
	t.t.gobble()
}
func (t *TurkeyAdapter) fly() {
	t.t.fly()
}

/**
어댑터 패턴은 특정 클래스 인터페이스를 클라이언트에서 요구하는 다른 인터페이스로 변환 합니다. 인터페이스가 호환되지않아 같이 쓸 수 없었던 클래스를 사용할 수 있게 도와준다.
*/

type Iterator interface {
	hasNext()
	next()
	remove()
}
type Enumeration interface {
	hasMoreElements()
	nextElement()
}

type EnumerationIterator struct {
	i Iterator
}

func (e *EnumerationIterator) hasMoreElements() {
	e.i.hasNext()
}
func (e *EnumerationIterator) nextElement() {
	e.i.next()
}
func (e *EnumerationIterator) remove() {
	log.Errorf("unsupported error")
}
