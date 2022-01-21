package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ERROR(e error) {
	if e != nil {
		log.Fatalln("error : ", e)
	}
}
func main() {
	fmt.Println("### - Best Time to Buy and Sell Stock - https://leetcode.com/problems/best-time-to-buy-and-sell-stock/ ")
	fmt.Println("Give a space seperated array : ")
	by, _, err := bufio.NewReader(os.Stdin).ReadLine()
	ERROR(err)
	s := string(by)
	sli := strings.Split(s, " ")
	var li []int
	for _, j := range sli {
		x, err := strconv.Atoi(j)
		ERROR(err)
		li = append(li, x)
	}

	fmt.Println("your string : ", li, "\n Enter your terget value : ")
	
	maxProfit(li)
}


func maxProfit(prices []int) int {
    i:=prices[0]
	max:=0
	for _,j:= range prices{
		if j-i >max {
			max=j-i
		}
		if j<i {
			i=j
		}
	}
	fmt.Println(max)
	return max
}