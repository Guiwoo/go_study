package main

type Function func(any) any
type Transformer func(Function) Function

type Recursive func(Recursive) Function

func (r Recursive) Apply(f Transformer) Function {
	return f(r(r))
}

func Y(f Transformer) Function {
	g := func(r Recursive) Function {
		return func(x any) any {
			return r.Apply(f)(x)
		}
	}
	return g(g)
}

func main() {

}
