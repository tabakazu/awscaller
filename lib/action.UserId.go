package awscaller

import (
	"fmt"

	"github.com/urfave/cli"
)

func actionUserID() cli.Command {
	return cli.Command{
		Name:  "userid",
		Usage: "Shows a aws api caller user id",
		Action: func(c *cli.Context) error {
			callerID := getCallerIdentity()
			fmt.Printf("- UserId : %v\n", *callerID.UserId)
			return nil
		},
	}
}
