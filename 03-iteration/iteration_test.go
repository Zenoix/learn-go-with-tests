package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 12)
	}
}

func ExampleRepeat() {
	repeated := Repeat("hi", 3)
	fmt.Println(repeated)
	// Output: hihihi
}
