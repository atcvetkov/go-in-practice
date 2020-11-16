package main

/*
$ flag_cli -h / --h / -help / --help

$ flag_cli
Hello World!
$ flag_cli –s –name Buttercup
Hola Buttercup!
$ flag_cli –-spanish –name Buttercup
Hola Buttercup!
*/

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish Language.")
}

func main() {
	flag.Parse()

	testString := "TEST"
	fmt.Println(testString) // TEST

	testStringPointer := &testString
	fmt.Println(testStringPointer) // 0xc000010250

	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
