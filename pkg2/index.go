package pkg2

import "fmt"

// Greet returns a greeting message
var SDK = "SAMPLE_PKG"

func Greet(name string) string {
	return fmt.Sprintf("Greetings from pkg2, %s!", name)
}
