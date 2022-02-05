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
	
	x:=maxProduct(li)
	fmt.Println(x)
}

func maxProduct(nums []int) int {
	currentMax:= nums[0]
	currentmin := nums[0]
	finalMax:= nums[0] 
	MAX:= func (a int,b int)int{
		if a>b {
			return a
		}
		return b
	}
	MIN:= func (a int,b int)int{
		if a<b {
			return a
		}
		return b
	}
	for i:=1;i<len(nums);i++{
		temp:= currentMax

		currentMax =MAX(MAX(currentMax * nums[i],currentmin*nums[i]),nums[i])
		currentmin= MIN(MIN(temp*nums[i],currentmin*nums[i]),nums[i])
		finalMax=MAX(currentMax,finalMax)
	}
	return finalMax
}

