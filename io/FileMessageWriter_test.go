package io

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {

	fmt.Printf("Hello!\n")
	fmt.Printf("Bugger!\n")
	fmt.Print("Wow!\n")
	// var buf bytes.Buffer
	// tracer := New(&buf)
	// if tracer == nil {
	// 	t.Error("Return from New should not be nil")
	// } else {
	// 	tracer.Trace("Hello trace package.")
	// 	if buf.String() != "Hello trace package.\n" {
	// 		t.Errorf("Trace should not write '%s'.", buf.String())
	// 	}
	// }

}
