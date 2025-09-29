package main

import(
	
	"time"
)
func sendMessage(message string,ch chan string,duration time.Duration){
     time.Sleep(duration)
	 ch<-message
}
// func main(){
// 	ch:=make(chan string)
// 	go sendMessage("Hello from Goroutine",ch,1*time.Second)
// 	go sendMessage("Hello from Main",ch,2*time.Second)
// 	go sendMessage("Hello from Another Goroutine",ch,3*time.Second)
  
// 	for i:=1;i<=3;i++{
// 		msg:=<-ch
// 		fmt.Println(msg)
// 	}

// }