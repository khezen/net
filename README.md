# net

## net/ip

```golang
// SubnetIPV4 - takes an ip v4 as string("192.168.220.254") and a subnet mask as int (17)
// and returns a the corresponding subnet as string ("192.168.92.0/17")
func SubnetIPV4(ipv4 string, mask int) (subnet string, err error)
```