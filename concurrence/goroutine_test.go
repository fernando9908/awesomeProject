package concurrence

import (
	"os"
	"testing"
)

func TestHelloGoRountine(t *testing.T) {
	HelloGoRountine()
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func BenchmarkHelloGoRountine(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HelloGoRountine()
	}
}

func BenchmarkHelloGoRountineParallel(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			HelloGoRountine()
		}
	})
}
