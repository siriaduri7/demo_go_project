/* ##################################################################################################################

SYNTAX : DECLARING AND INITIALIZING A SLICE using make

slice := make([]<data_type, length, capacity )

Generally, make function is used to create an empty slice.

length based on lets say [:5] = 4 is length
capacity = 5

##################################################################################################################
A slice is defined as a continuous segment of an underlying array, slices provide access to parts of an array in 
sequential order. 
They are more flexible and more powerful than arrays since arrays had limitations of being fixed size.
variable types (elements can be added or removed)
###########A slice has three major components.
Pointer, length, and capacity.
pointer : The pointer is used to point to the first element of the array that is accessible through that slice.
It is not necessary that the pointed element is the first element of the array.
but pointers are just variables that hold memory addresses.
When you change the value of an element in slice, it gets affected in the underlying array as well.

####################################################################################################################
SYNTAX:

<slice_name := []<data_type>{<values> }

####################################################################################################################

######################## DECLARING AND INITIALIZING A SLICE #########################################
package main
import "fmt"

func main() {
	slice := []int{10,20,30}
	fmt.Println(slice)
}
o/p [10 20 30]
##########################################################################################################################################################

 [ 10 20 30 97 54 60 ]
 array[0 :3] // 10 20 30
 array [1:6] // 20 30 97 54 60 */

 /* package main
 import "fmt"

 func main () {
	arr := [10]int {10, 20, 30, 40, 50 ,60 ,70 , 80, 90, 100}
	slice := arr[1:8]
	fmt.Println(cap(arr))
	fmt.Println(cap(slice))
	fmt.Println(slice)

	sub_slice := slice[0:3] // subslice is done from abive slice o/p
	fmt.Println(sub_slice)

 }
 //o/p [20 30 40 50 60 70 80]
 // subslice o/p [20 30 40]

/*
package main
import "fmt"

func main() {
	slice := make([]int, 5 ,8) //slice := make([]<data_type, length, capacity )
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}

o/p :
[0 0 0 0 0]
5
8

############################ SLICE AND INDEX NUMBERING ########################## */
/*package main
import "fmt"

func main() {
	arr := [5]int{ 10, 20, 30, 40 , 50}
	slice := arr [:3]
	fmt.Println(arr)
	fmt.Println(slice)

	slice[1] =9000
	fmt.Println("after modification")
	fmt.Println(arr)
	fmt.Println(slice)

}*/
/* o/p
[10 20 30 40 50]
[10 20 30]
after modification
[10 9000 30 40 50]
[10 9000 30] */

/*
##################################### APPENDING TO A SLICE #############################

package main
import "fmt"
func main() {
	arr := [4]int{10,20,30,40}
	slice := arr[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 900, -90, 50)

	slice := arr[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	
}
o/p
[20 30]
2
3
[20 30 900 -90 50]
5
6

################################ appending to aslice #################
syntax
slice = append(slice, anotherslice...)

#########################################################

package main
import "fmt"

func main() {
	arr_1 := [5]int{10,20,30,40,50}
	slice := arr_1[:2]

	arr_2 := [5]int{5,15,25,35,45}
	slice_2 := arr_2[:2]
	new_slice = append(slice, slice_2...)
	fmt.Println(new_slice)

}

o/p
[10 20 5 15]

################################### DELETING FROM A SLICE #############################

package main
import "fmt"
func main () {
	arr := [5]int{10,20,30,40,50}
	i := 2 // refering index 2 = 20
	fmt.Println(arr)
	slice_1 = arr[:i]    // index 2 (-1) => [10 20]
	slice_2 =arr[i+1:]   // index 2+1 = 3 [40 50]
	new_slice = append(slice_1,slice_2...)
	fmt.Println(new_slice)
}

o/p
[10 20 30 40 50]
[10 20 40 50]


####################################### COPYING FROM A SLICE ################################

SYNTAX :

func copy(dst,src []type) int

num :=copy(dest_slice, src_slice)
************************************************************************************
*/
package main
import "fmt"

func main() { /// practise again 
	
	src_slice := []int{10,20,30,40,50}
	dest_slice := make([]int, 3)
	num := copy(dest_slice, src_slice)
	fmt.Println(dest_slice)
	fmt.Println("number of elements: ", num)

}
/*
o/p
[10 20 30]
number of elements:  3