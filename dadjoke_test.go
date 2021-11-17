package dadjoke_test

import (
	"fmt"
	"testing"

	dadjoke "github.com/anders-14/godadjoke"
)

func TestFetch(t *testing.T) {
	fmt.Println(dadjoke.Fetch())
}
