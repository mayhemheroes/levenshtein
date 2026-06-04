package fuzz

import "testing"

func FuzzComputeDistance(f *testing.F) {
	f.Add("hello\nworld")
	f.Fuzz(func(t *testing.T, data string) {
		_ = Fuzz([]byte(data))
	})
}
