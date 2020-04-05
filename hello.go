package hello

import (
	"fmt"

	"rsc.io/quote"
)

// Hello returns a greeting.
func Hello() string {
	return (quote.Hello())
}

// SomeOneSays returns some wise phrase.
func SomeOneSays(name string) string {
	return (fmt.Sprintf("%s says: %s", name, quote.Hello()))
}
