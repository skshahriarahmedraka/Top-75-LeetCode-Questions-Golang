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
	
	x:=maxSubArray(li)
	fmt.Println(x)
}

func maxSubArray(nums []int) int {
    max:=nums[0]
	temp:= 0

	for _,i:= range nums {
		temp += i 
		if temp >max {
			max=temp 
		}
		if temp <0 {
			temp=0
		}
	}
	return max 
}