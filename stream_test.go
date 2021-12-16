package stream_test

import (
	"fmt"

	"github.com/r6q/stream"
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
