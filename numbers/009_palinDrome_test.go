package numbers

import (
	"fmt"
	"testing"
)

func TestPalinDrome(t *testing.T) {
	x := 1421242344234241
	res := palinDrome(x)
	fmt.Printf("%d:%v", x, res)
}
