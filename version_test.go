package gotime

import "testing"

func TestVersion(t *testing.T) {
	t.Log(Version)
}

func BenchmarkVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(Version)
	}
}
