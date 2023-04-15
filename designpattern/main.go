package main

import "designpattern/bridge"

/**
Complicated object ts aren't designed from scratch
- They reiterate exisitng desings

An existing desing is a Prototype
We make a copy of the prototype and customize it
Requires 'deep copy' support
We make the cloning convenient via a Factory
*/

// Deep copyting

func main() {
	bridge.Start()
}
