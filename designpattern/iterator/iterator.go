package iterator

import "fmt"

/**
Iteration (traversal) is a core functionality of various data structures
An iterator is a type that facilitates the traversal
	- keeps a pointer to the current element
	- Knows how to move to a different element
Go allows iteration with range
	- built-in support in many objects(arrays,slice,etc(map)
	- Can be supported in out own structs

Iterator An object that facilitates the traversal of a data structure
*/

type Person struct {
	FirstName, MiddleName, LastName string
}

func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()
	return out
}

/**
Iterator Un idiomatic way , it used to in C++
*/

type PersonNameIterator struct {
	person  *Person
	current int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{person, -1}
}
func (p *PersonNameIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}
func (p *PersonNameIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	}
	panic("We should not be here!")
}

/**
Binary Tree Traverse
*/

type Node struct {
	value               int
	left, right, parent *Node
}

func NewNode(value int, left, right *Node) *Node {
	n := &Node{value: value, left: left, right: right}
	left.parent = n
	right.parent = n
	return n
}
func NewTerminalNewNode(value int) *Node {
	return &Node{value: value}
}

type InOrderIterator struct {
	Current       *Node
	root          *Node
	returnedStart bool
}

func NewInOrderIterator(root *Node) *InOrderIterator {
	i := &InOrderIterator{root: root, Current: root, returnedStart: false}
	for i.Current.left != nil {
		i.Current = i.Current.left
	}
	return i
}

func (i *InOrderIterator) Reset() {
	i.Current = i.root
	i.returnedStart = false
}

func (i *InOrderIterator) MoveNext() bool {
	if i.Current == nil {
		return false
	}
	if !i.returnedStart {
		i.returnedStart = true
		return true
	}
	if i.Current.right != nil {
		i.Current = i.Current.right
		for i.Current.left != nil {
			i.Current = i.Current.left
		}
		return true
	} else {
		p := i.Current.parent
		for p != nil && i.Current == p.right {
			i.Current = p
			p = p.parent
		}
		i.Current = p
		return i.Current != nil
	}
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{root: root}
}
func (b *BinaryTree) InOrder() *InOrderIterator {
	return NewInOrderIterator(b.root)
}

func Start() {
	p := &Person{"Alexander", "Graham", "Bell"}
	for _, name := range p.Names() {
		fmt.Println(name)
	}

	fmt.Println()

	for name := range p.NamesGenerator() {
		fmt.Println(name)
	}

	fmt.Println()

	for pp := NewPersonNameIterator(p); pp.MoveNext(); {
		fmt.Println(pp.Value())
	}

	fmt.Println("=====")
	root := NewNode(1, NewTerminalNewNode(2), NewTerminalNewNode(3))

	it := NewInOrderIterator(root)
	for it.MoveNext() {
		fmt.Printf("%d,", it.Current.value)
	}
}

/**
An iterator specifies how you can Traverse an object
Moves along the iterated collection, indication when last element has been reached
Not idiomatic in Go
*/
