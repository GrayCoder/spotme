# spotme
Small go command line tool to request the cheapest spot instance available

The tool expects the following environment variables:
- AWS_DEFAULT_REGION
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY

usage:

-c, --cloud-config <file name>
-a, --ami <ami id>
-k, --key <key name>
-s, --security-group <security group id>
-i, --instance-types <list of instance types>
-t, --test

aws ec2 describe-spot-price-history --instance-types m1.xlarge --availability-zone us-east-1b --start-time 2016-03-06T09:00:00 --end-time 2016-03-06T09:30:00 | jq '.'

possible commands:
spotme --instance-types=m3.xlarge,m4.xlarge                                                                                                                                                                                                
ex.                                        
  m3.xlarge                                
    us-east-1a,0.03                        
    us-east-1b,0.03                        
    us-east-1c,0.02                        
    us-east-1d,0.03                        
    us-east-1e,0.04                        
  m4.xlarge                                
    us-east-1a,0.13                        
    us-east-1b,0.13                        
    us-east-1c,0.12                        
    us-east-1d,0.13                        
    us-east-1e,0.14                        
                                           
spotme --instance-types=m3.xlarge,m4.xlarge --low=true
  m3.xlarge,us-east-1c,0.02                          
                                           
                                           
spotme list --instance-types=m3.xlarge,m4.xlarge --low=true | spotme request --cloud-config=somefile --ami=ami-345345 --security-group=sg-472f7923
