package cli

import (
	"github.com/alecthomas/kong"
	termError "github.com/roysti10/termCI/internal/errors"
	run "github.com/roysti10/termCI/internal/run"
	cli "github.com/roysti10/termCI/internal/structs"
)

func Run() error {
	ctx := kong.Parse(&cli.CLI)
	switch ctx.Command() {
	case "run":
		return run.Execute()
	default:
		return termError.TermError{Msg: ctx.Command()}
	}
}
