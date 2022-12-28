/*Create an ARRAY which holds 5 VALUES of TYPE int
	● assign VALUES to each index position.
	● Using format printing ○ print out the TYPE of the array */

	package main
	import "fmt"
	import "reflect"
	func main() {
		var arr [5]int =[5]int{40,50,60,70,80}
		fmt.Println(arr)
		//fmt.Printf("arr = %T\n", arr)
		fmt.Printf("Type : %v \n",reflect.TypeOf(arr))

		
		arr[0] =100
		arr[1] =101
		arr[2] = 102
		arr[3] = 103
		arr[4] =104
		fmt.Println(arr)
		for index, element := range arr {
			fmt.Println(index, "=>" , element)
		}
/*  Create a SLICE of TYPE int
● assign 10 VALUES
● Using format printing ○ print out the TYPE of the slice */
		// slice :=make([]int, 10,10)
		slice :=[]int{10,20,30,40,50,60,70,80,90,100}
		fmt.Printf("type of slice: %v \n",reflect.TypeOf(slice))
/*3.     Using the code from the previous example,
use SLICING to create the following new slices which are then printed:
	● [42 43 44 45 46]
	● [47 48 49 50 51]
	● [44 45 46 47 48]
	● [43 44 45 46 47]*/

	    slice[0] =42
		slice[1] = 43
		slice[2] =44
		slice[3] =45
		slice[4] = 46
		slice[5] =47
		slice[6]=48
		slice[7]=49
		slice[8]=50
		slice[9] =51
		
		fmt.Println(slice)
		slice1 := slice[0:5]
		fmt.Println(slice1)
		slice2 := slice[5:]
		fmt.Println(slice2)
		slice3 := slice[2:7] 
		fmt.Println(slice3)
		slice4 := slice[1:6]
		fmt.Println(slice4)

		x:=[]int{42,43,44,45,46,47,48,49,50,51}
		a := []int{52}
		b :=append(x, a...)
		fmt.Println(b)
		c := []int{53,54,55}
		d :=append(b,c...)
		fmt.Println(d)
		y:=[]int{56,57,58,59,60}
		z:=append(x,y...)
		fmt.Println(z)

		



		

         

	}


	
