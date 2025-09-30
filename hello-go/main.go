package main

 import "fmt"

func process(items ...interface{}) {
	for _, item := range items {
		if v, ok := item.(int); ok {
			fmt.Printf("Integer: %d\n", v)
		}else if v, ok := item.(string); ok {
			fmt.Printf("String: %s\n", v)
		}else if v,ok := item.([]int);ok{
		  fmt.Printf("Process Array:\n")
		  for _,num := range v{
			fmt.Printf("%d\n",num)
		  }
		}else if v,ok := item.([]string);ok{
	     fmt.Printf("Process String Array:\n")
		 for _,str := range v{
			fmt.Printf("%s\n",str)
		 }
		}else {
			fmt.Printf("type: %T  %v\n", item, item)
		}
	}
}
func main() {
	// arrayFunction()
	// twoDimensionalArray()
	// displayStudentGroupWise()
    // testMethodsWithReceiver()
	process(42, "hello", 3.14, true, []int{1, 2, 3}, map[string]int{"a": 1, "b": 2})
}
