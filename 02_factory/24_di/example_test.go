package di_test

import (
	"fmt"

	di "github.com/CausingBrick/design-pattens/02_factory/24_di"
)

// Dependence relationship  A -> B -> C

// A
type A struct {
	B   *B
	Num int
}

// NewA NewA
func NewA(b *B) *A {
	return &A{
		B:   b,
		Num: 1,
	}
}

// B B
type B struct {
	C   *C
	Num int
}

// NewB NewB
func NewB(c *C) *B {
	return &B{
		C:   c,
		Num: 2,
	}
}

// C C
type C struct {
	Num int
}

// NewC NewC
func NewC() *C {
	return &C{
		Num: 3,
	}
}

func Example() {
	c := di.New()

	panicErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}
	// Inject constructor
	panicErr(c.Provide(NewA))
	panicErr(c.Provide(NewB))
	panicErr(c.Provide(NewC))

	err := c.Invoke(func(a *A) {
		fmt.Printf("%d %d %d", a.Num, a.B.Num, a.B.C.Num)
	})
	panicErr(err)

	// Output: 1 2 3
}
