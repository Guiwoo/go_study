package proxy

import "fmt"

/**
Motivation
- Calling foo.Bar()
- foo is in the same process as Bar()
- What if you wanna put all Foo-related operations into a separate process ?

Proxy will Handle this Situations !
- Same Interface, entirely different behavior
"A type that functions as an interface to a particular resource.
That resource may be remote, expensive to construct, or may require logging or some other added functionality ."
*/

type Driven interface {
	Drive()
}

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 19 {
		c.car.Drive()
	} else {
		fmt.Println("Driver too young")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, driver}
}

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitMap(filename string) *Bitmap {
	fmt.Println("Loading image from", filename)
	return &Bitmap{filename: filename}
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image", b.filename)
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func (l *LazyBitmap) Draw() {
	if l.bitmap == nil {
		l.bitmap = NewBitMap(l.filename)
	}
	l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func Start() {
	car := NewCarProxy(&Driver{12})
	car.car.Drive()
	car.Drive()

	fmt.Println()

	_ = NewBitMap("demo.png")
	//DrawImage(bmp)
}

/**
Proxy vs Decorator
- [Proxy] provides an identical interface; decorator provides an enhanced interface
- [Decorator] typically aggregates (or has reference to) what it is decorating; proxy doesn't have to
- [Proxy] might not even be working with a materialized object


Proxy has the same interface as the underlying object
To create a proxy, simply replicate the existing interface of an object
Add relevant functionality to the redefined methods
Different proxies (communication, logging, caching, etc.) have completely different behaviors
*/
