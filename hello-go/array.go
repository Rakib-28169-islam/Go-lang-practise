package main

import "fmt"

type intArr []int
func arrayFunction(){
		// Array
	var arr = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slice
	slice := []string{"apple", "banana", "cherry"}
	slice = append(slice, "date")
	fmt.Println("Slice:", slice)

	// Map
	m := make(map[string]int)
	m["Alice"] = 25
	m["Bob"] = 30
	fmt.Println("Map:", m)

	arrName := []int{1,2,3,4,5,6,7,8,9,10}
	sliceArr:=arrName[0:len(arrName)/2]
	fmt.Println("Int Slice:", sliceArr)
	newArr:=arrName[:]
	fmt.Println("New Array:", newArr)


	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

    
    part1 := original[2:5]     // [2, 3, 4]
    part2 := original[:3]      // [0, 1, 2] 
    part3 := original[7:]      // [7, 8, 9] 
    part4 := original[:]       // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9] 

    fmt.Println("Original:", original)
    fmt.Println("part1 [2:5]:", part1)
    fmt.Println("part2 [:3]:", part2)
    fmt.Println("part3 [7:]:", part3)
    fmt.Println("part4 [:]:", part4)

	// Modifying the original slice
	original[0] = -1000
	fmt.Println("\nAfter modifying original:")
     fmt.Println("Original:", original)
    fmt.Println("part1 [2:5]:", part1)
    fmt.Println("part2 [:3]:", part2)
    fmt.Println("part3 [7:]:", part3)
    fmt.Println("part4 [:]:", part4)


	//slice copy without reference
	fmt.Println("\nSlice Copy without reference:")
	sliceA:=[]int{1,2,3,4,5}
	sliceB:=make([]int,len(sliceA))
	copy(sliceB,sliceA)
	fmt.Println("Slice A:",sliceA)
	fmt.Println("Slice B:",sliceB)
	sliceA[0]=-100
	fmt.Println("After modifying Slice A:")
	fmt.Println("Slice A:",sliceA)
	fmt.Println("Slice B:",sliceB)
}

func twoDimensionalArray(){

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D Array (Matrix):")
	for i,row :=range matrix{
		for j,val :=range row{
			fmt.Printf("matrix[%d][%d] = %d\n",i,j,val)
		}
	}
}

func (arr intArr) sum()int{
	sum:=0
	for _,val:=range arr{
		sum+=val
	}
	return sum
}
func (arr *intArr) append(values ...int){
	*arr = append(*arr,values...)
}
func (arr intArr) get(index int)(int,error){
	if index<0 || index>=len(arr){
		return 0,fmt.Errorf("array index out of bound")
	}
	return arr[index],nil
}
func (arr *intArr) pop()(int,error){
	if len(*arr)==0{
		return 0,fmt.Errorf("array is empty")
	}
	val:=(*arr)[len(*arr)-1]
	*arr=(*arr)[:len(*arr)-1]
	return val,nil
}
func (arr *intArr) sort(){
	for i:=0;i<len(*arr)-1;i++{
		for j:=0;j<len(*arr)-i-1;j++{
			if (*arr)[j]>(*arr)[j+1]{
				(*arr)[j],(*arr)[j+1]=(*arr)[j+1],(*arr)[j]
			}
		}
	}
}

func testMethodsWithReceiver(){
	arr := intArr{1, 2, 3, 4, 5,-1,2,6}//intArr is a custom type defined in array.go
	fmt.Println("Initial Array:", arr)
	
	// Using the method to calculate the sum
	sum := arr.sum()
	fmt.Println("Sum of elements:", sum)

	// Using the method to append elements
	arr.append(6, 7, 8)
	fmt.Println("After appending elements:", arr)

	// Using the method to get an element
	val, err := arr.get(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Element at index 2:", val)
	}

	// Using the method to pop an element
	poppedVal, err := arr.pop()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Popped element:", poppedVal)
		fmt.Println("Array after popping:", arr)
	}
	
	// Using the method to sort the array
	arr.sort()
	fmt.Println("Sorted Array:", arr)


//expected output:

// 	After appending elements: [1 2 3 4 5 -1 2 6 6 7 8]
// Element at index 2: 3
// Popped element: 8
// Array after popping: [1 2 3 4 5 -1 2 6 6 7]
// Sorted Array: [-1 1 2 2 3 4 5 6 6 7]
}