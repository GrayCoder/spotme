package main

import (
	"flag"
	"fmt"
	"strings"
	//	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Config struct {
	Ami   string
	Az    string
	Cc    string
	It    string
	Key   string
	Sg    string
	Test  bool
	Until string
}

//func getLatestPrice(az string, it string, ) {
//
//}

//	req1 := ec2.DescribeSpotPriceHistoryInput{
//		AvailabilityZone:    aws.String("us-east-1a"),
//		DryRun:              aws.Bool(config.Test),
//		StartTime:           aws.Time(time.Now().Add(time.Duration(-2) * time.Hour)),
//		EndTime:             aws.Time(time.Now()),
//		InstanceTypes:       aws.StringSlice(strings.Split(config.It, ",")),
//		ProductDescriptions: aws.StringSlice([]string{"Linux/UNIX (Amazon VPC)"}),
//	}
//	sresp1, err := svc.DescribeSpotPriceHistory(&req1)
//	if err != nil {
//		panic(err)
//	}
//	for x, y := range sresp1.SpotPriceHistory {
//		fmt.Println(x)
//		fmt.Println(*y.SpotPrice)
//	}

var config Config

func init() {
	flag.StringVar(&config.Ami, "ami", "", "the ami id to use for the spot request")
	flag.StringVar(&config.Az, "availability-zones", "", "the availability zones to search (defaults to all)")
	flag.StringVar(&config.Cc, "cloud-config", "", "the filename (with relative path) of the cloud-config to use")
	flag.StringVar(&config.It, "instance-types", "", "the list of instance types (comma separated) that are acceptable to use as hardware")
	flag.StringVar(&config.Key, "key", "", "the name of the key to use")
	flag.StringVar(&config.Sg, "security-groups", "", "the list of security groups (comma separated) to attach to the instance")
	flag.BoolVar(&config.Test, "test", true, "if true, will generate a sample request and print it")
	flag.StringVar(&config.Until, "until", "", "the length of time you want the spot request to last for")
}

func main() {
	flag.Parse()
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	fmt.Println("Instances: ", strings.Split(config.It, ","))

	aresp, err := svc.DescribeAvailabilityZones(&ec2.DescribeAvailabilityZonesInput{DryRun: aws.Bool(config.Test)})
	if err != nil {
		panic(err)
	}
	fmt.Println(aresp)
	// Here we'd like to parse the instance types, find all their current prices in all the AZs, and then pick the cheapest
}
