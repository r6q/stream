package stream_test

import (
	"fmt"
	"testing"

	"github.com/r6q/stream"
	"github.com/stretchr/testify/assert"
)

func ExampleStream_Map() {
	res := stream.Of(1, 2, 3).
		Map(func(it int) int {
			return it * it
		}).
		Collect()
	fmt.Println(res)

	// Output: [1 4 9]
}

func TestFilter(t *testing.T) {
	actual := stream.Of("quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog").
		Filter(func(it string) bool {
			if len(it) > 3 {
				return true
			}
			return false
		}).
		Collect()

	assert.Equal(t, actual, []string{"quick", "brown", "jumps", "over", "lazy"})
}
