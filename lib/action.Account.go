package awscaller

import (
	"fmt"

	"github.com/urfave/cli"
)

func actionAccount() cli.Command {
	return cli.Command{
		Name:  "account",
		Usage: "Shows a aws api caller account number",
		Action: func(c *cli.Context) error {
			callerID := getCallerIdentity()
			fmt.Printf("- Account : %v\n", *callerID.Account)
			return nil
		},
	}
}
