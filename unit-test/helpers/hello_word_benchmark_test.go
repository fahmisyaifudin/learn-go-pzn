package helpers

import "testing"

func BenchmarkHelloWorld(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fahmi")
	}
}

func BenchmarkHelloWorldSub(b *testing.B)  {
	b.Run("Fahmi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fahmi")
		}
	})
	b.Run("Syaifudin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Syaifudin")
		}
	})
}