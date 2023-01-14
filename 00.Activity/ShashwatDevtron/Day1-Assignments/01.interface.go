package main

import (
	"fmt"
	"time"
)

type USD int32

type Config struct {
    IPAddress string
	TypeOfIP string
	State string
	TimeStamp time.Time
	Paid string
	Cost USD
}
func verifyIP(ip string) bool {
	count := 0
	for _,c := range ip {
		if c == '.' {
			count += 1
		}
	}
	if count != 3 {
		return false
	}
	return true
}
func verifyType(t string) bool {
	if t == "Public" || t == "Private"{
		return true
	}
	return false
}
func verifyState(t string) bool {
	if t == "Running" || t == "Stopped" || t == "Exited"{
		return true
	}
	return false
}
func verifyPaid(t string) bool {
	if t == "Yes" || t == "No" || t == "yes" || t == "no" {
		return true
	}
	return false
}

func (c *Config) verify() (bool,string) {
	if verifyIP(c.IPAddress) && verifyType(c.TypeOfIP) && verifyState(c.State) && verifyPaid(c.Paid) {
		return true, ""
	}
	return false, "Varification Failed"
}

func serverConfig(config *Config) {
	c,err := config.verify()
	if err != "" {
		fmt.Println(err)
	}
	fmt.Println(c)
}

func main() {
	s := Config {
		IPAddress : "127.0.0.1",
		TypeOfIP : "Public",
		State : "Running",
		TimeStamp : time.Now(),
		Paid : "Yes",
		Cost : 125,
	}
    serverConfig(&s)
	fmt.Println(s)
}