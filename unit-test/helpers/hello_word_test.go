package helpers

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M)  {
	fmt.Println("Before test")
	m.Run()
	fmt.Println("After test")
}


func TestHelloWorld(t *testing.T)  {
	hello := HelloWorld("Fahmi")
	if hello != "Hello Fahmi" {
		t.Fail()
		t.Error("Error and continue test....")
	}
	fmt.Println("This line will executed")
}

func TestSayHello(t *testing.T)  {
	hi := SayHello("Fahmi")
	if hi != "Hi Fahmi" {
		t.FailNow()	
		t.Fatal("Error and exit test ...")
	}
	fmt.Println("This line will not executed")
}

func TestHelloWorldAssert(t *testing.T)  {
	hello := HelloWorld("Fahmi")
	assert.Equal(t, "Hello Fahmi", hello)
}

func TestSkip(t *testing.T)  {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	hello := HelloWorld("Fahmi")
	assert.Equal(t, "Hello Fahmi", hello)
}

func TestSubTest(t *testing.T)  {
	t.Run("Fahmi", func(t *testing.T) {
		hello := HelloWorld("Fahmi")
		assert.Equal(t, "Hello Fahmi", hello)
	})
	t.Run("Syaifudin", func(t *testing.T) {
		hello := HelloWorld("Syaifudin")
		assert.Equal(t, "Hello Syaifudin", hello)
	})
}