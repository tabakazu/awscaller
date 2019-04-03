package awscaller

import (
	"fmt"
	"regexp"

	"github.com/urfave/cli"
)

func actionUserName() cli.Command {
	return cli.Command{
		Name:  "username",
		Usage: "Shows a aws api caller user name",
		Action: func(c *cli.Context) error {
			callerID := getCallerIdentity()
			userName := regexp.MustCompile(`^(\S.*)\/`).ReplaceAllString(*callerID.Arn, "")
			fmt.Printf("- UserName : %v\n", userName)
			return nil
		},
	}
}
