package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Config struct {
	Ami   string
	Az    string
	Cc    string
	ExAz  string
	It    string
	Key   string
	R     string
	Sg    string
	Test  bool
	Until string
}

func getCurrentPrice(conn *ec2.EC2, it string, az string) string {
	req := ec2.DescribeSpotPriceHistoryInput{
		AvailabilityZone:    aws.String(az),
		DryRun:              aws.Bool(config.Test),
		StartTime:           aws.Time(time.Now().Add(time.Duration(-2) * time.Hour)),
		EndTime:             aws.Time(time.Now()),
		InstanceTypes:       aws.StringSlice(strings.Split(it, ",")),
		ProductDescriptions: aws.StringSlice([]string{"Linux/UNIX (Amazon VPC)"}),
	}
	aresp, err := conn.DescribeSpotPriceHistory(&req)
	if err != nil {
		panic(err)
	}
	// list is sorted by time so we just take the first element
	return *aresp.SpotPriceHistory[0].SpotPrice
}

var config Config

func init() {
	region := os.Getenv("AWS_DEFAULT_REGION")
	if region == "" {
		region = "us-east-1"
	}
	flag.StringVar(&config.Ami, "ami", "", "the ami id to use for the spot request")
	flag.StringVar(&config.Az, "availability-zones", "", "the availability zones to search (defaults to all)")
	flag.StringVar(&config.Cc, "cloud-config", "", "the filename (with relative path) of the cloud-config to use")
	flag.StringVar(&config.ExAz, "excluded-availability-zones", "", "the availability zones to exclude")
	flag.StringVar(&config.It, "instance-types", "", "the list of instance types (comma separated) that are acceptable to use as hardware")
	flag.StringVar(&config.Key, "key", "", "the name of the key to use")
	flag.StringVar(&config.R, "region", region, "the region to run against")
	flag.StringVar(&config.Sg, "security-groups", "", "the list of security groups (comma separated) to attach to the instance")
	flag.BoolVar(&config.Test, "test", false, "if true, will generate a sample request and print it")
	flag.StringVar(&config.Until, "until", "", "the length of time you want the spot request to last for")
}

// Here we'd like to parse the instance types, find all their current prices in all the AZs, and then pick the cheapest
func main() {
	flag.Parse()
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String(config.R)})
	aresp, err := svc.DescribeAvailabilityZones(&ec2.DescribeAvailabilityZonesInput{DryRun: aws.Bool(config.Test)})
	if err != nil {
		panic(err)
	}
	prices := make([]float64, 0)
	for _, v := range aresp.AvailabilityZones {
		p := getCurrentPrice(svc, "m3.large", *v.ZoneName)
		pf, err := strconv.ParseFloat(p, 64)
		if err != nil {
			panic(err)
		}
		prices = append(prices, pf)
	}
	fmt.Println(prices)
}
