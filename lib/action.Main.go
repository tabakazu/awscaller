package awscaller

import (
	"fmt"
	"regexp"

	"github.com/urfave/cli"
)

func actionMain(c *cli.Context) {
	callerID := getCallerIdentity()
	userName := regexp.MustCompile(`^(\S.*)\/`).ReplaceAllString(*callerID.Arn, "")
	listPolicies := listAttachedUserPolicies(userName)
	fmt.Printf("- Account  : %v\n", *callerID.Account)
	fmt.Printf("- UserId   : %v\n", *callerID.UserId)
	fmt.Printf("- UserName : %v\n", userName)
	fmt.Println("- AttachedPolicies :")
	for _, policy := range listPolicies.AttachedPolicies {
		fmt.Printf("    - %v\n", *policy.PolicyName)
	}
}
