package main

import (
	"testing"
)

func BenchmarkHitValue(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, err := hitValue()
		if err != nil {
			b.Errorf("failed to get str: %s", err)
		}
	}

}

func TestHitValue(t *testing.T) {
	_, err := hitValue()
	if err != nil {
		t.Errorf("failed to get str: %s", err)
	}
}

func BenchmarkHitPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := hitPointer()
		if err != nil {
			b.Errorf("faild to get str: %s", err)
		}
	}
}
