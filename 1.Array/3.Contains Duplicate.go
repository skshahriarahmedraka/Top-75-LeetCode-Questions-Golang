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
	
	x:=containsDuplicate(li)
	fmt.Println(x)
}

func containsDuplicate(nums []int) bool {
    m:= make(map[int]bool)
	for _,i := range nums {
		if _,err:= m[i] ; !err {
			m[i]=true
			continue
		}else {
			return true 
		}
	}
	return false
}