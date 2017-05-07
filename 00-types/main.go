package main

import (
	"fmt"
)

type awsService struct {
	serviceName string
	regionsAvailable []string 
}

func main() {
	
	s3 := awsService{
		serviceName: "Simple Storage Service (S3)",
		regionsAvailable: []string{"eu-west-1", "us-east-1"},
	}

	fmt.Println("* Service detials *")
	fmt.Println("Name:", s3.serviceName)
	fmt.Println("Regions available:", s3.regionsAvailable)
}