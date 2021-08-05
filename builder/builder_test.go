package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	t.Run("builder1", func(t *testing.T) {
		testBuilder(t, new(ConcreateBuilder1), "1")
	})
	t.Run("builder2", func(t *testing.T) {
		testBuilder(t, new(ConcreateBuilder2), "2")
	})
}

func testBuilder(t *testing.T, b builder, want string) {
	director := new(Director)
	director.builder = b
	result := director.Construct()
	if want != result {
		t.Errorf("want: %s, result: %s", want, result)
	}
}
