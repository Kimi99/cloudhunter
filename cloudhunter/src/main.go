package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func main() {
	fmt.Println("[!] Enter AWS account region: ")

	var region string
	_, err := fmt.Scanln(&region)

	regions := []string { 
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-northeast-3",
		"ap-south-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"ca-central-1",
		"eu-central-1",
		"eu-north-1",
		"eu-west-1",
		"eu-west-2",
		"eu-west-3",
		"sa-east-1",
		"us-east-1",
		"us-east-2",
		"us-west-1",
    	"us-west-2" }

	if err != nil  {
		log.Fatal(err)
	} else if !slices.Contains(regions, region) {
		fmt.Println("[-] Region not valid.")
	}
	
	// Initialize a session in desired region that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if err != nil {
		log.Fatal(err)
	}

	// Create a IAM service client.
	svc := iam.New(sess)

	result, err := svc.ListUsers(&iam.ListUsersInput{
		MaxItems: aws.Int64(10),
	})

	if err != nil {
		fmt.Println("[-] Error", err)
		return
	}

	for _, user := range result.Users {
		if user == nil {
			continue
		}

		fmt.Printf("[+] Found user! Username: %s; Created Date: %v\n", *user.UserName, user.CreateDate)
	}
}