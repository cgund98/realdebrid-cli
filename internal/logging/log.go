package logging

import (
	"fmt"
	"os"
)

// Print a message and exit the program
func Fatalf(msg string, a ...any) {
	fmt.Fprintf(os.Stderr, msg+"\n", a...)
	os.Exit(1)
}
