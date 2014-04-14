package valueref

import "fmt"

func ExampleMyInteger() {
	v := MyInteger{2}
	fmt.Println("1: v.Value is", v.Value) // 1: v.Value is 2

	DoubleMyInteger(v)
	fmt.Println("2: v.Value is", v.Value) // 2: v.Value is 2

	DoubleMyIntegerPtr(&v)
	fmt.Println("3: v.Value is", v.Value) // 3: v.Value is 4
	// Output:
	// 1: v.Value is 2
	// 2: v.Value is 2
	// 3: v.Value is 4
}

func ExampleDoubleMyIntegerPtr() {
	i := 2
	fmt.Println("1: i is", i) // 1: i is 2

	DoubleInt(i)
	fmt.Println("2: i is", i) // 1: i is 2

	DoubleIntPtr(&i)
	fmt.Println("3: i is", i) // 1: i is 4
	// Output:
	// 1: i is 2
	// 2: i is 2
	// 3: i is 4
}

func ExampleMyIntegerDouble() { // OMIT
	v := MyInteger{2}
	v.Double()
	fmt.Println(v.Value) // 2
	// Output:
	// 2
}

func ExampleMyIntegerDoublePtr() { // OMIT
	p := &MyInteger{2}
	p.DoublePtr()
	fmt.Println(p.Value) // 4
	// Output:
	// 4
}

func ExampleByIntDouble() { // OMIT
	v := MyInt(2)
	v.Double()
	fmt.Println(v) // 2
	// Output:
	// 2
}

func ExampleByIntDoublePtr() { // OMIT
	v := MyInt(2)
	p := &v
	p.DoublePtr()
	fmt.Println(v) // 4
	// Output:
	// 4
}
