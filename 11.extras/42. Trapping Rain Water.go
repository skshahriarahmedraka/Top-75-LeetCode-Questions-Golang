//https://leetcode.com/problems/trapping-rain-water/

package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	// "os"
	// "strconv"
)

func ERROR(e error){
	if e!=nil {
		log.Fatalln("error : ",e)
	}
}

func main(){
	var n int 

	b,_,err:= bufio.NewReader(os.Stdin).ReadLine()

	ERROR(err)

	s:= string(b)
	ss:=strings.Split(s," ")
	var li []int 

	for _,i:= range ss {
		b,err:=strconv.Atoi(i)
		ERROR(err)
		li=append(li,b)
	}
	


}

// func trap(height []int) int {
    
// }