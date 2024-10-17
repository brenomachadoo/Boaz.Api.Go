package exceptionsUtil

import "fmt"

type ExceptionGeneric struct {
	err string //error description
}

func New(text string) error {
	return &ExceptionGeneric{text}
}

func (e *ExceptionGeneric) Error() string {
	fmt.Print(e.err)
	return e.err
}
