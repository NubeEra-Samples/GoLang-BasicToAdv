package main

import (
	"time"
)

type IpConfig struct {
	IpAdd     IpAddress
	typeIp    string
	timestamp time.Time
	state     string
	Paid      string
	cost      int64
}
type IpAddress struct {
	a string
	b string
	c string
	d string
}

func main() {

}
