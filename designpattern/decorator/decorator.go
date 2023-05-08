package decorator

import "fmt"

/**
Want to augment an object with additional functionality
Do not want to rewrite or alter existing code (OCP)
Want to keep new functionality separate (SRP)

"Facilitates the addition of behaviors to individual objects through embedding"
*/

type Bird struct {
	Age int
}

func (b *Bird) Fly() {
	if b.Age >= 10 {
		fmt.Println("Flying!")
	}
}

type Lizard struct {
	Age int
}

func (l *Lizard) Crawl() {
	if l.Age < 10 {
		fmt.Println("Crawling!")
	}
}

type Dragon struct {
	Bird
	Lizard
}

func (d *Dragon) Age() int {
	return d.Bird.Age
}
func (d *Dragon) SetAge(age int) {
	d.Bird.Age = age
	d.Lizard.Age = age
}

type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird2 struct {
	age int // age is lower case
}

func (b *Bird2) Age() int       { return b.age }
func (b *Bird2) SetAge(age int) { b.age = age }

func (b *Bird2) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying!")
	}
}

type Lizard2 struct {
	age int
}

func (l *Lizard2) Age() int       { return l.age }
func (l *Lizard2) SetAge(age int) { l.age = age }
func (l *Lizard2) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling!")
	}
}

type Dragon2 struct {
	bird   Bird2
	lizard Lizard2
}

func (d *Dragon2) Fly() {
	d.bird.Fly()
}
func (d *Dragon2) Crawl() {
	d.lizard.Crawl()
}

func NewDragon() *Dragon2 {
	return &Dragon2{Bird2{}, Lizard2{}}
}

// Other Example

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle with radius %f", c.Radius)
}
func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}
func (s *Square) Resize(factor float32) {
	s.Side *= factor
}

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

type TransParentShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransParentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}

/**
type ColoredSquare struct {
	Square
	Color string
}
*/

func Start() {
	//d := Dragon{}
	//d.Bird.Age = 10
	//d.Bird.Fly()
	//d.Lizard.Age = 5
	//d.Lizard.Crawl()
	//d.Fly()
	//d.Crawl()

	circle := Circle{2}
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())
	c := redCircle.Shape.(*Circle)
	c.Resize(22)

	fmt.Println(redCircle.Render())
	rhsCircle := TransParentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())
}
