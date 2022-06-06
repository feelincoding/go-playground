package main

import "fmt"

type Data struct {
	value int
	data  [2]int
}

func ChangeData(args Data) {
	args.value = 999
	args.data[0] = 999
}
func ChangeData2(args Data) {
	args.value = 999
	args.data[0] = 999
}
func ChangeData3(args *Data) {
	args.value = 999
	args.data[0] = 999
}
func ChangeData4(args *Data) {
	args.value = 100
	args.data[0] = 100
}
func main() {
	var data Data
	ChangeData(data)
	fmt.Println(data)
	var p *Data = &data
	ChangeData2(*p)
	fmt.Println(data)
	ChangeData3(p)
	fmt.Println(data)
	ChangeData4(&data)
	fmt.Println(data)

	fmt.Println("===========================")
	var p2 = new(Data)
	p2 = &data
	ChangeData2(*p2)
	fmt.Println(data)
}
