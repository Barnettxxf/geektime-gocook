package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func fn() error {
	e1 := errors.New("error")
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")
	return errors.Wrap(e3, "outer")
}
func main() {
	// type strakTracker interface {
	// 	StackTrace() errors.StackTrace
	// }

	// err, ok := errors.Cause(fn()).(strakTracker)
	// if !ok {
	// 	panic("oops, err does not implement stackTracker")
	// }

	// st := err.StackTrace()
	// fmt.Printf("%+v", st)
	fmt.Printf("%+v", fn())
}
