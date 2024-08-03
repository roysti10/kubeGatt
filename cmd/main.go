package main

import (
	term "github.com/roysti10/termCI/cmd/cli"
	termError "github.com/roysti10/termCI/internal/errors"
)

func main() {
	err := term.Run()
	termError.CheckError(err)
}
