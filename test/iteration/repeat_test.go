package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("By default 5 times serie expected", func(t *testing.T) {
		result := Repeat(0, "s")
		expected := "sssss"
		if result != expected {
			t.Errorf("expected %q but got %q", expected, result)
		}
	})
	t.Run("If provided, x times series expected", func(t *testing.T) {
		result := Repeat(10, "z")
		expected := "zzzzzzzzzz"
		if result != expected {
			t.Errorf("expected %q but got %q", expected, result)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(5, "s")
	}
}
