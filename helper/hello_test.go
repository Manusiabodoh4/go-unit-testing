package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("Reza")
	if result != "Hello Reza" {
		t.Fail()
		fmt.Println("Function tidak mengembalikan Hello Reza")
	}
}

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("Reza")
	if result != "Hello Reza" {
		t.FailNow()
		fmt.Println("Function tidak mengembalikan Hello Reza")
	}
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Reza")
	if result != "Hello Reza" {
		t.Error("Result must be Hello Reza")
		fmt.Println("Function tidak mengembalikan Hello Reza")
	}
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Reza")
	if result != "Hello Reza" {
		t.Fatal("Result must be Hello Reza")
		fmt.Println("Function tidak mengembalikan Hello Reza")
	}
}
