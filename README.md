# SubnetIpRange
## Overview
Subnet IP range calculation.

We may need to known the IP range of a subnet in our program sometimes. The subnet may like `172.16.1.19/20`, and the IP shoud range from `172.16.1.19` to `172.16.15.255`. To conveniently invoke function, I try to write a package to provide calculation of subnet IP range.

## Testing
You can test this package by `range_test.go` like following:
```
go test range_test.go
```


