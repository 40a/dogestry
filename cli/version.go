package cli

import (
	"fmt"
)

const Version string = "2.1.6"

func PrintVersion() error {
	_, err := fmt.Printf("Dogestry %s\n", Version)
	return err
}

func (cli *DogestryCli) CmdVersion(args ...string) error {
	return PrintVersion()
}
