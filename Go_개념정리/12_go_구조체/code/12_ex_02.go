package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}
type VipUser struct {
	UserInfo User
	VipLevel int
}
type VipUser2 struct {
	User
	VipLevel int
}

func main() {
	var user1 User = User{
		Name: "홍길동",
		Age:  20,
	}
	var user2 User = User{
		"김유신",
		30,
	}
	var vipUser1 VipUser
	vipUser1 = VipUser{
		UserInfo: user1,
		VipLevel: 1,
	}
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(vipUser1)
}
