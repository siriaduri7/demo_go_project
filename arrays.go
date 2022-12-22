/*  An array a collection of similar data elements that are stored at contiguous memory locations.
 we know that an array is a collection of data elements of the same data type.
syntax : in arrays CAPACITY = LENGTH
      var <array name> [<size of the array>] <data type>
ex :  var  grades [5] int
var fruits [3] string
############# ARRAY INITIALIZATION ######################
var grades [3]int = [3]int{10,20,30}

###################### ARRAY SHORTHAND DECLARATION ####################
 
grades := [3]int{10, 20, 30}

################################## ELLIPSES #################################
We have a name of the array, the shorthand operator, these three dots over here in Golang are turned as ellipses.
grades = [...]int{10, 20, 30 }
we need not specify the size or the length of our array. In this case, the compiler will implicitly calculate the length of
 the array based on the number of elements that we've specified here.


ex:

package main
import "fmt"
func main() {
	var grades [5] int
	var fruits [3] string
	fmt.Println(fruits)
}
o/p [0 0 0 0 0]
[   ]
This is because we haven't initialized the array yet.Since the array is of integer data type, we get zero, because zero is 
the zero value of integers.
//Now over here it might not be visible, but the second array actually consists of three empty strings.

############# ARRAY INITIALIZATION ######################

var grades [3]int = [3]int{10,20,30}
//The number of values should always be less than or equal to the size of the array. here it shud be less or equal to [3]

################We can initialize an array also using the shorthand declaration. The syntax is simple.###########
grades := [3]int{10, 20, 30} */
/*
package main
import "fmt"
func main() {
	var grades [3]int = [3]int{10, 20 , 30}
	

	 fruits := [2]string{"apple", "grapes"}
	

	 names := [...]string{"sri", "remo","sumo"}
	 //fmt.Println(grades)
	 //fmt.Println(fruits)
	 //fmt.Println(names)
	 fmt.Printf("grades values %d, fruits value %v, names values %v", grades, fruits, names)
} */
 /* ########################### len() ##########################
 Now, the length of the array refers to the number of elements stored in the array.
 The len function over here takes an array as an input and returns us the size of the array.
 packages main
 import "fmt"
 func main() {
	var animals [2]string = [2]string{"cat", "dog"}
	fmt.Println(len(animals)
 }
 o/p : 2
 In addition to this, an array is also numbered and these numbers are called array index.

 grades [90 86 76 42 85]
 index  [0   1  2  3  4]
 grades[3] => 42 ; grades [0] => 90
 #############################################
 package main
 import "fmt"
 func main() {
	var fruits [4]string = [4]string{"apple","grapes","orange","pine"}

	fmt.Println(fruits[3])
 }
 o/p : pine 

 ################ We first printed it, then we're going to change the value of the element at index one, and#################

 (*******************************) first printed => change the value ********************************************
  
 package main
 import "fmt"
 func main() {
	var grades [5]int = [5]int{60, 70, 80, 90, 100}
	fmt.Println(grades)
	grades[1] = 50
	fmt.Println(grades)
 }

 ################# FOR-loop through an array ###########################

 // for i = 0; i < len(grades); i++
 fmt.Println(grades[i]) */
 /* package main
 import "fmt"
 func main() {
	var grades [4]int =[4]int {80,40,30,90}
	for i :=0 ; i <len(grades); i++ {
		fmt.Println(grades[i]) // focus here we used grades and i in println
	}
 }
 
 /* ######################### LOOPING THROUGH AN ARRAY ########################
 // for index, element :=range grades { //initialize index,element variables, that stores return values and we have a range expression, which consists of
	fmt.Println(index, "=>", element) //range keyword and grades is array name
	
//The range keyword is mainly used in for loops, in order to iterate over all the elements of an array, slice, or map. */

package main
import "fmt"
func main () {
 var grades [5]int =[5]int {30,50,70,60,85}
 for index, element := range grades {
	fmt.Println(index, "=>" , element)
 }
 
}
/* o/p:
0 => 30
1 => 50
2 => 70
3 => 60
4 => 85
/* ########################## MULTIDIMENSIONAL ARRAY #####################
A multi-dimensional array is an array that has more than one dimension. It's like an array of arrays or basically an array 
that has multiple levels.
The simplest multi-dimensional array is the 2-d array or two-dimensional array.
index    0  1
(0)     [2, 4]
(1)		[4,16]
(2)		[8,64]

package main
import "fmt"
func main() {
	arr := [3][2]int{{2,4}, {4,16}, {8,64}} //[3][2] here [3] refres rows [2] refers columns
	fmt.Println(arr[2][1])
}
o/p: 64

