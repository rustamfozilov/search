package search

import (
	"context"
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {

	ctx := context.Background()

	f := []string {
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
	}

	all := All(ctx, "6", f)
	fmt.Println(all)
}