package main

import "fmt"

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	var house1 House
	house1.Address = "서울시 서초구 서초동"
	house1.Size = 50
	house1.Price = 100
	house1.Category = "아파트"

	fmt.Println(house1)
	fmt.Printf("%v\n", house1)

	fmt.Printf("")
}
