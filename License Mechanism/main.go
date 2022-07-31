package main

import "fmt"

//The Proxy pattern is a structural pattern whose purpose is to provide a surrogate or intermediary of an object to control its access.
//Usually wraps an object to hide some of its characteristics.

type Provider interface {
	Provide()
}

type Software struct{}

func (c *Software) Provide() {
	fmt.Println("Not reached to limit")
}

type User struct {
	Num0fUsers int
}

type SoftwareProxy struct {
	soft Software
	user *User
}

func (c *SoftwareProxy) Provide() {
	if c.user.Num0fUsers <= 32 {
		c.soft.Provide()
	} else {
		fmt.Println("Concurrent number of users reached the limit for this software")
	}
}

func NewSoftwareProxy(user *User) *SoftwareProxy {
	return &SoftwareProxy{Software{}, user}
}

func main() {
	soft := NewSoftwareProxy(&User{67})
	soft.Provide()
}
