package awscaller

import (
	"fmt"
	"regexp"

	"github.com/urfave/cli"
)

func actionAttachedPolicies() cli.Command {
	return cli.Command{
		Name:  "policies",
		Usage: "Shows a aws api caller attached policies",
		Action: func(c *cli.Context) error {
			callerID := getCallerIdentity()
			userName := regexp.MustCompile(`^(\S.*)\/`).ReplaceAllString(*callerID.Arn, "")
			listPolicies := listAttachedUserPolicies(userName)
			fmt.Println("- AttachedPolicies :")
			for _, policy := range listPolicies.AttachedPolicies {
				fmt.Printf("\t- %v\n", *policy.PolicyName)
			}
			return nil
		},
	}
}
