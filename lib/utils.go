package awscaller

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/sts"
)

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
