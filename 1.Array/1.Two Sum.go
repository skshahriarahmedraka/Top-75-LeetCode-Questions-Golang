package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "sort"
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
	fmt.Println("ans: ",twoSum(li,t))
}

// func twoSum(nums []int, target int) []int {
// 	sort.Ints(nums)
// 	for i,n:= range nums {
// 		if n <target{
// 			x:=sort.SearchInts(nums,target-n)
// 			fmt.Println("now :",i,x,n)
// 			if x!= len(nums)-1 || (x==len(nums)-1 && nums[x]==target-n) {
// 				return []int{i,x}
// 			}
// 		} 
// 	}
// 	return nums
// }
func twoSum(nums []int, target int) []int {
	m:= make(map[int]int ,5)

	for i,j:= range nums {
		m[j]=i
	}
	for i,j := range nums {
		if j<target {
			x:=target -j 
			if b,err := m[x];err{
				return []int{i,b}
			}else{
				continue 
			}
		}
	}

	return nums
}