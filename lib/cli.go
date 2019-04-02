package awscaller

import (
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/urfave/cli"
)

const appName string = "awscaller"
const appUsage string = "display aws api caller infomation"
const appVersion string = "0.0.2"

// CliApp return *cli.App
func CliApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion
	app.Action = mainCmd
	return app
}

func mainCmd(c *cli.Context) {
	callerID := getCallerIdentity()
	fmt.Printf("- Account  :\t%v\n", *callerID.Account)
	fmt.Printf("- UserId   :\t%v\n", *callerID.UserId)
	userName := regexp.MustCompile(`^(\S.*)\/`).ReplaceAllString(*callerID.Arn, "")
	fmt.Printf("- UserName :\t%v\n", userName)

	listPolicies := listAttachedUserPolicies(userName)
	fmt.Println("- AttachedPolicies :")
	for _, policy := range listPolicies.AttachedPolicies {
		fmt.Printf("\t- %v\n", *policy.PolicyName)
	}
}

func listAttachedUserPolicies(userName string) *iam.ListAttachedUserPoliciesOutput {
	svc := iam.New(session.New())
	input := &iam.ListAttachedUserPoliciesInput{
		UserName: aws.String(userName),
	}
	result, err := svc.ListAttachedUserPolicies(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return nil
	}
	return result
}

func getCallerIdentity() *sts.GetCallerIdentityOutput {
	svc := sts.New(session.New())
	input := &sts.GetCallerIdentityInput{}
	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return nil
	}
	return result
}
