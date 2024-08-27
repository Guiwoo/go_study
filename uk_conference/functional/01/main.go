package main

import (
	"fmt"
	"os"
	"strconv"
)

type Integer interface {
	int | int8 | int16 | int32 | int64
}

type Cache[T Integer] map[T]T

func (c Cache[T]) Dump() {
	for k, v := range c {
		fmt.Printf("c[%v],value:%v\n", k, v)
	}
}

type Function func(any) any
type Transformer func(Function) Function

func recurF(f any) Function {
	return f.(func(any) any)(f).(func(any) any)
}
func YRecurF(g Function) Function {
	return recurF(
		func(f any) any {
			return g(func(x any) any {
				return recurF(f)(x)
			})
		},
	)
}
func TransFormRecur(f Function) Function {
	return f(f).(Function)
}
func YTransFormer(g Transformer) Function {
	return TransFormRecur(
		func(f any) any {
			return g(func(x any) any {
				return TransFormRecur(f.(Function))(x)
			})
		})
}

func main() {
	combinator()
}

func combinator() {
	f := Y(func(g any) any {
		return func(n any) (r any) {
			if n, ok := n.(int); ok {
				switch {
				case n < 0:
					panic(n)
				case n < 2:
					return 1
				}
				return n * g.(func(any) any)(n-1).(int)
			}
			panic(n)
		}
	})
	os.Exit(f(5).(int))
}

// Y => f.(->x.f(x x))(->x.f(x x)
func Y(g func(any) any) func(any) any {
	return func(f any) func(any) any {
		return f.(func(any) any)(f).(func(any) any)
	}(func(f any) any {
		return g(func(x any) any {
			return f.(func(any) any)(f).(func(any) any)
		})
	})
}

func cache() {
	f := cacheFactorial[int]()
	PrintErrors := IfPanics(PrintErrorMessage)
	Each(os.Args[1:], func(v string) {
		PrintErrors(
			ValidInteger(v, func(i int) {
				fmt.Printf("%v!: %v\n", i, f(i))
			}),
		)
	})
}
func cacheFactorial[T Integer]() (f func(T) T) {
	c := make(Cache[T])
	return func(n T) (r T) {
		if r = c[n]; r == 0 {
			switch {
			case n < 0:
				panic(n)
			case n == 0:
				r = 1
			default:
				r = n * f(n-1)
			}
			c[n] = r
			c.Dump()
		}
		return
	}
}

func recur() {
	printErrors := IfPanics(PrintErrorMessage)
	Each(os.Args[1:], func(v string) {
		printErrors(ValidInteger(v, func(i int) {
			fmt.Printf("%v! : %v\n", i, recurFactorial(i))
		}),
		)
	})
}

func Each[T any](s []T, f func(T)) {
	if len(s) > 0 {
		f(s[0])
		Each(s[1:], f)
	}
}

func recurFactorial[T Integer](n T) (r T) {
	switch {
	case n < 0:
		panic(n)
	case n == 0:
		r = 1
	default:
		r = n * recurFactorial(n-1)
	}
	return
}

func currying() {
	printErrors := IfPanics(PrintErrorMessage)
	for _, v := range os.Args[1:] {
		printErrors(
			ValidInteger(v, func(i int) {
				fmt.Printf("%v!: %v\n", i, Factorial(i))
			}))
	}
}

func ValidInteger[T Integer](v string, f func(i T)) func() {
	return func() {
		if x, e := strconv.Atoi(v); e == nil {
			f(T(x))
		} else {
			panic(e)
		}
	}
}

func IfPanics(e func()) func(func()) {
	return func(f func()) {
		defer e()
		f()
	}
}

func PrintErrorMessage() {
	if x := recover(); x != nil {
		fmt.Printf("no defined value for %v\n", x)
	}
}

func validFunc() {
	var errors int
	computeFactorial := ForValidValues(
		func(i int) {
			fmt.Printf("%v! : %v\n", i, Factorial(i))
		},
		func() {
			if x := recover(); x != nil {
				fmt.Printf("no defined value for %+v\n", x)
				errors++
			}
		},
	)

	for _, v := range os.Args[1:] {
		computeFactorial(v)
	}
	os.Exit(errors)
}

func ForValidValues[T Integer](f func(T), e func()) func(string) {
	return func(v string) {
		defer e()

		if x, e := strconv.Atoi(v); e == nil {
			f(T(x))
		} else {
			panic(v)
		}
	}
}

func PanicForHandleByFunc() {
	for _, v := range os.Args[1:] {
		func() {
			defer func() {
				if x := recover(); x != nil {
					fmt.Println("no factorial...")
				}
			}()
			if x, e := strconv.Atoi(v); e == nil {
				fmt.Printf("%v! : %v\n", x, Factorial(x))
			} else {
				panic(x)
			}
		}()
	}
}

func Factorial[T Integer](n T) (r T) {
	if n < 0 {
		panic(n)
	}
	for r = 1; n > 0; n-- {
		r *= n
	}
	return
}
