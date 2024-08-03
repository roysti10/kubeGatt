package errors

import (
	"fmt"
	"os"
)

type TermError struct {
	Msg string
}

func (e TermError) Error() string {
	return fmt.Sprintf("TermError: %s", e.Msg)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "A error occured:%v\n", err)
		os.Exit(1)
	}

}
