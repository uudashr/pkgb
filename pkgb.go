package pkgb

import (
	"fmt"

	"github.com/uudashr/pkga"
)

// DoSomething just do something.
func DoSomething() string {
	s := pkga.SayHello()
	return fmt.Sprintf("Do something with a =>> %s", s)
}
