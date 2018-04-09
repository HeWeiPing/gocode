package tempconv_test

import "fmt"
import "testing"
import "tempconv"

func TestCToF(t *testing.T) {
	fmt.Printf("%d C->F=%g\n", 100, tempconv.CToF(100))
}
