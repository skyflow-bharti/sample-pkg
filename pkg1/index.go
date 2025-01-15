package pkg1

import (
	"fmt"
	"github.com/skyflow-bharti/sample-pkg/pkg2"
)

// Hello returns a greeting message
func Hello(name string) string {
	fmt.Println("here is my SDK Name", pkg2.SDK)
	return fmt.Sprintf("Hello from pkg1, %s!", name)
}
