package main
import (
	"fmt"
	// "time"
	"sync"
)

func printNumbers(msg string,wg *sync.WaitGroup){
	defer wg.Done()
	for i:=1;i<=5;i++{
		fmt.Println(msg,":",i)
		
	}
	// time.Sleep(1*time.Second)

}
func printMultiplicationTable(num int,wg *sync.WaitGroup){
		defer wg.Done()

	for i:=1;i<=10;i++{
		fmt.Printf("%d x %d = %d\n",num,i,num*i)
		
	}
	// time.Sleep(1*time.Second)
}
func printEvenOdd(num int,wg *sync.WaitGroup){
	defer wg.Done()
	if num%2==0{
		fmt.Printf("%d is Even\n",num)
	} else {
		fmt.Printf("%d is Odd\n",num)
	}
	// time.Sleep(1*time.Second)
} 
// func main(){
// 	// printNumbers("Direct")
//     var wg sync.WaitGroup
	
// 	for i:=1;i<=3;i++{
// 		wg.Add(3)
// 		go printNumbers("Goroutine", &wg)
// 		go printMultiplicationTable(i, &wg)
// 		go printEvenOdd(i, &wg)
// 		wg.Wait()
		
// 	}
  
// 	fmt.Print("\n\n")
// 	wg.Wait()


// }