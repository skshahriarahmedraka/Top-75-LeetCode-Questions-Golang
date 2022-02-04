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
	
	x:=productExceptSelf(li)
	fmt.Println(x)
}

func productExceptSelf(nums []int) []int {
	ans:= make([]int,len(nums))
	for i:=1 ;i<=len(nums)-1;i++ {
		if i== 1 {
			ans[i]=nums[i-1]
		}else {
			ans[i]=ans[i-1]*nums[i-1]
		}
	}
	
	left:= nums[len(nums)-1]
	for i:=len(nums)-2;i>=0;i--{
		if (i==0) {
			ans[i]=left
		} else{
			ans[i]*=left
			left*=nums[i]
		}
	}
	return ans  
    
}