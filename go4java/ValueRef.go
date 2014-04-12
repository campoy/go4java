package main

import "fmt"

type Integer struct {
	Value int
}

func DoubleInteger(v Integer) { v.Value *= 2 }
func DoubleInt(v int)         { v *= 2 }

func DoubleIntegerPtr(v *Integer) { v.Value *= 2 } // or (*v).Value
func DoubleIntPtr(v *int)         { *v *= 2 }

func main() {
	v := Integer{2}      // v.Value is 2
	DoubleInteger(v)     // v.Value is still 2
	DoubleIntegerPtr(&v) // v.Value is now 4
	fmt.Println(v)

	i := 2           // i is 2
	DoubleInt(i)     // i is still 2
	DoubleIntPtr(&i) // i is now 4
	fmt.Println(i)
}
