package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
func ERROR(e error){
	if e != nil{
		log.Fatalln("error : ",e )
	}
}
func main() {
	fmt.Println("### Two Sum : https://leetcode.com/problems/two-sum/ ")
	fmt.Println("Give a space seperated array : ")
	by,_,err:= bufio.NewReader(os.Stdin).ReadLine()
	ERROR(err)
	s:= string(by)
	sli:= strings.Split(s," ")
	var li []int 
	for _,j:= range sli{
		x,err:=strconv.Atoi(j)
		ERROR(err)
		li = append(li,x)
	}

	fmt.Println("your string : ", li,"\n Enter your terget value : ")
	var t int 
	fmt.Scanln(&t)
	twoSum(li,t)
}

func twoSum(nums []int, target int) []int {
	m :=make(map[int]bool) 
	m2 :=make(map[int]int) 
	for i,j := range nums {
		m[j]=true 
		m2[j]=i
	}
	var li []int
	for i,j := range nums{		
		if m[target-j] && len(li)<2 && (target-j)!=j {
			li=append(li,i)
			li=append(li,m2[target-j])
		}
	}
	fmt.Println(li)
	return nums
}