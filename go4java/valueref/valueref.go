// Package valueref defines functions and methods passing pointers and values
// to demonstrate the call by value mechanism in Go.
package valueref

// MyInteger contains an int Value.
type MyInteger struct {
	Value int
}

type MyInt int

// DoubleMyInteger doubles the Value in the given MyInteger.
// This won't have any side effect, since the parameter is a copy of the
// original MyInteger.
func DoubleMyInteger(v MyInteger) { v.Value *= 2 }

// DoulbeMyIntegerPtr double the Value in the MyInteger pointed by the given
// pointer.
// This will actually modify the Value in the original MyInteger, since we're
// passing a pointer.
func DoubleMyIntegerPtr(v *MyInteger) { v.Value *= 2 } // or (*v).Value

// DoubleInt doubles the value of the given int.
// This won't have any side effect, since the parameter is a copy of the
// original int.
func DoubleInt(v int) { v *= 2 }

// DoubleIntPtr doubles the value of the int pointed by the given pointer.
// This will actually modify the value of the original int, since we're
// passing a pointer.
func DoubleIntPtr(v *int) { *v *= 2 }

// Double doubles the Value in the MyInteger.
// The original MyInteger is not modified, since the receiver is a copy.
func (i MyInteger) Double() { i.Value *= 2 }

// DoublePtr doubles the Value in the MyInteger.
// The original MyInteger is modified, since the receiver is a pointer.
func (i *MyInteger) DoublePtr() { i.Value *= 2 }

// Double double the value of the given MyInt.
// The original MyInt is not modified, since the receiver is a copy.
func (i MyInt) Double() { i *= 2 }

// DoublePtr doubles the Value of the MyInt.
// The original MyInt is modified, since the receiver is a pointer.
func (i *MyInt) DoublePtr() { (*i) *= 2 }
