package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("repeat 5 times by default", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat 2 times when specified", func(t *testing.T) {
		repeated := RepeatTimes("a", 2)
		expected := "aa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
