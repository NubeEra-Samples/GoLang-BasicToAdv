package main

import "fmt"

type Conf struct {
	IpAddress string
	TypeOfIp  string
	TimeStamp string
	State     string
	Paid      bool
	cost      string
}

func (c *Conf) UpdateIp(ip string) {
	count := 0
	for _,val:= range ip {
		if val == '.' {
			count++
		}
	}
	if count != 3 {
		return
	} else {
		c.IpAddress = ip
	}
}

func (c *Conf) updateTypeOfIp(ipType string){
	if(ipType!="Private"&&ipType!="Public"){
		return
	}
	c.TypeOfIp = ipType

}

func (c* Conf) updateState(newState string){
	if(newState!="Running"&&newState!="Shutdown"&&newState!="Restart"){
		return
	}
	c.State = newState
}

func (c* Conf) updateTime(){
	
}

func (c* Conf) updatePaid(newP bool){
	c.Paid = newP
}

func (c* Conf) updateCost(newCost string){
	c.cost = newCost;
}

func main() {
	var c Conf
	c.UpdateIp("1.2.4.5")
	c.updateCost("125$ per Hour")
	c.updatePaid(true)
	c.updateState("Running")
	c.updateTypeOfIp("Private")
	fmt.Print(c)
}