//https://leetcode.com/problems/trapping-rain-water/

package main

import (
	// "bufio"
	"fmt"
	"log"
	// "os"
	// "strconv"
)

func ERROR(e error){
	if e!=nil {
		log.Fatalln("error : ",e)
	}
}

func main(){
	var li []int 

	// b,_,err := bufio.NewReader(os.Stdin).ReadLine()
	// ERROR(err)

	// s:= string(b)
	// ss,err := strconv.Atoi(s)
	var a int 
	for i:=0 ;i<3;i++{
		fmt.Scan(&a)
		li= append(li,a)
	}
	fmt.Println(li)
	fmt.Printf("%T",li)
	 for _,i:= range li {
		 
	 }
	


}

// func trap(height []int) int {
    
// }