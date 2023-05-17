package composit

import (
	"fmt"
	"strings"
)

/**
A mechanism for treating individual objects and
compositions of objects in a uniform manner.
*/

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')
	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{"Circle", color, nil}
}

func NewSquare(color string) *GraphicObject {
	return &GraphicObject{"Square", color, nil}
}

func Start() {
	drawing := GraphicObject{"My Drawing", "", nil}

	drawing.Children = append(drawing.Children, *NewCircle("Red"))
	drawing.Children = append(drawing.Children, *NewSquare("Yellow"))

	group := GraphicObject{"Group 1", "", nil}
	group.Children = append(group.Children, *NewCircle("Blue"))
	drawing.Children = append(drawing.Children, group)

	fmt.Println(drawing.String())
}

type NeuronI interface {
	Iter() []*Neuron
}

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range n.Neurons {
		result = append(result, &n.Neurons[i])
	}
	return result
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

func Connect(left, right NeuronI) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}

func Start2() {
	neuron1, neuron2 := &Neuron{}, &Neuron{}
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron1)
	Connect(layer1, layer2)
}

type Component interface {
	Print() string
	Add(Component)
}
type Employee struct {
	Name, Position string
	Sub            []Component
}

func (e *Employee) Print() string {
	return fmt.Sprintf("Name: %s, Position: %s", e.Name, e.Position)
}
func (e *Employee) Add(c Component) {
	e.Sub = append(e.Sub, c)
}

type Department struct {
	Name       string
	Components []Component
}

func (d *Department) Print() string {
	return fmt.Sprintf("Department: %s", d.Name)
}
func (d *Department) Add(c Component) {
	d.Components = append(d.Components, c)
}

type SongComponent interface {
}

type Song struct {
	Name, Singer string
}
type PlayList struct {
	Name       string
	Components []SongComponent
}
