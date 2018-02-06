package iso4217_test

import (
	"fmt"

	"github.com/rmg/iso4217"
)

func ExampleByCode() {
	name, minor := iso4217.ByCode(0)
	fmt.Printf("name: '%s', minor: %d\n", name, minor)
	// Output: name: '', minor: 0
}

func ExampleByCode_cad() {
	name, minor := iso4217.ByCode(124)
	fmt.Printf("name: '%s', minor: %d\n", name, minor)
	// Output: name: 'CAD', minor: 2
}

func ExampleByCode_xxx() {
	name, minor := iso4217.ByCode(999)
	fmt.Printf("name: '%s', minor: %d\n", name, minor)
	// Output: name: 'XXX', minor: 0
}

func ExampleByCode_unknown() {
	name, minor := iso4217.ByCode(1234591)
	fmt.Printf("name: '%s', minor: %d\n", name, minor)
	// Output: name: '', minor: 0
}

func ExampleByName() {
	code, minor := iso4217.ByName("")
	fmt.Printf("code: %d, minor: %d\n", code, minor)
	// Output: code: 0, minor: 0
}

func ExampleByName_cad() {
	code, minor := iso4217.ByName("CAD")
	fmt.Printf("code: %d, minor: %d\n", code, minor)
	// Output: code: 124, minor: 2
}

func ExampleByName_xxx() {
	code, minor := iso4217.ByName("XXX")
	fmt.Printf("code: %d, minor: %d\n", code, minor)
	// Output: code: 999, minor: 0
}

func ExampleByName_unknown() {
	code, minor := iso4217.ByName("NOT_REAL")
	fmt.Printf("code: %d, minor: %d\n", code, minor)
	// Output: code: 0, minor: 0
}
