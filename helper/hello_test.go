package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TableTestingType struct {
	name, request, expected string
}

func TestMain(m *testing.M) {
	//biasanya untuk implementasi pemanggilan connection database
	fmt.Println("Before Running Unit Test")
	m.Run()
	fmt.Println("After Running Unit Test")
}

func TestTableTest(t *testing.T) {

	arrTable := []TableTestingType{
		{"Reza", "Reza", "Hello Reza"},
		{"Amel", "Amel", "Hello Amel"},
		{"Rizky", "Rizky", "Hello Rizky"},
		{"Uus", "Uus", "Hello Uus"},
		{"Wina", "Wina", "Hello Wina"},
	}

	for _, data := range arrTable {
		t.Run(data.name, func(t *testing.T) {
			result := HelloWorld(data.request)
			require.Equal(t, data.expected, result)
		})
	}

}

func TestSubTest(t *testing.T) {
	t.Run("Reza", func(t *testing.T) {
		result := HelloWorld("Reza")
		require.Equal(t, "Hello Reza", result)
	})
	t.Run("Amel", func(t *testing.T) {
		result := HelloWorld("Amel")
		require.Equal(t, "Hello Amel", result)
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can't run on Linux")
	}
	result := HelloWorld("Reza")
	require.Equal(t, "Hello Reza", result)
}

func TestHelloWorldWithReguire(t *testing.T) {
	result := HelloWorld("Reza")
	require.Equal(t, "Hello Reza", result)
}

func TestHelloWorldWithAssert(t *testing.T) {
	result := HelloWorld("Reza")
	assert.Equal(t, "Hello Reza", result)
}

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
