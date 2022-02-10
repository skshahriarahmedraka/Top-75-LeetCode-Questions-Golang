//https://leetcode.com/problems/median-of-two-sorted-arrays/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)
func ERROR(e error) {
	if e!= nil {
		log.Fatalln("error : ",e)
	}
}
func  main () {
	var (
		 li1 []int
		 li2 []int  
	)

	b, _, err := bufio.NewReader(os.Stdin).ReadLine()
	ERROR(err)
	s:= string(b)
	ss:= strings.Split(s," ")
	for _,i := range(ss){
		x,err:= strconv.Atoi(i)
		ERROR(err)
		li1 = append(li1, x)
	}
	fmt.Println("first list : ",li1)


	b, _, err = bufio.NewReader(os.Stdin).ReadLine()
	ERROR(err)
	s= string(b)
	ss= strings.Split(s," ")
	for _,i := range(ss){
		x,err:= strconv.Atoi(i)
		ERROR(err)
		li2 = append(li2, x)
	}
	fmt.Println("sceond list : ",li2)
	findMedianSortedArrays(li1,li2)

	// li2=append(li1,li2...)

	// fmt.Println("list :",li2)
	// sort.Ints(li2)

	// fmt.Println(li2)
	//  var ans float64
	// //  var ans2 int
	// if len(li2)%2 ==0 {
	// 	ans= float64(li2[int(len(li2)/2) -1]+li2[int(len(li2)/2)])/2
	// }else {
	// 	ans= float64(li2[int(len(li2)/2) -1])
	// }
	// fmt.Println("ans : ",ans)
	

}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    
	nums2=append(nums1,nums2...)

	sort.Ints(nums2)
	le:= len(nums2)
	 var ans float64
	if le%2 ==0 {
		ans= float64(nums2[int(le/2) -1]+nums2[int(le/2)])/2
	}else {
		ans= float64(nums2[int(le/2) ])
	}

	return ans 

}