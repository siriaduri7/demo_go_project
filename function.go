/* SYNTAX
func <function_name>(<i/p patameters>) <return type> {
	//body of function
}
ex:
func AddNumbers(a int, b int) int {
/*
}
/*
################## CALLING FUNCTION #########################
<function_name>(<argument(s)>)
AddNumbers (2,3)
SumOfNumbers := AddNumbers(2, 3)


#################### PARAMETERS VD ARGUMENTS #######################
func addNumbers(a int, b int) int {
	sum := a+b
	return sum
}

sunc main() {
	SumOfNumbers := addNumbers(2,3)
	fmt.Println(SumOfNumbers)

}
#################################################################################### (important points)
//Function parameters are the names that are listed  in the function's definition, //a int, b int (fn parameters)
******************************************************
//while function arguments are the real values that we pass to the function. // (2,3) (fn arguments)
####################################################################################
The statement containing the return values is also known as the terminating statement.
package main
import "fmt"

func operation(a int, b int) (int, int) { //if u declare as (sum int diff int ) at return values place than dnt use shorthand 
	sum := a+b //if intialized in retuen place dnt use shorthand 
	diff :=a - b // diff = a-b
	return sum,diff
}
func main () {
	sum,difference := operation(20,10)
	fmt.Println(sum, difference)
	 
}
*/
/*
############################ VARIADIC FUNCTIONS ##################################
Variadic functions are functions that accept variable number of arguments.
In Go, it is possible to pass a varying number of arguments  of the same data type as referenced in the function signature.
To declare a variadic function, the type of the final parameter should be preceded by an ellipsis,

syntax:

func <func_name>(param-1 type, param-2 type, para3 ...type) <return_type>

eg:

func sumNumbers(numbers ...int) int

func sumNumbers(tr string, number ...int)

//the variadic function, the type of the last parameter is preceded by an ellipsis,that is three dots. It indicate that the 
function can be called at any number of parameters of this type. */
/*
package main
import "fmt"

func sumofNumbers(numbers ...int) int {
	sum :=0
	for _, value :=range numbers{
		sum +=value
	}
	return sum 	

}
func main() {
	fmt.Println(sumofNumbers())
	fmt.Println(sumofNumbers(10,20))
	fmt.Println(sumofNumbers(30,40))
	fmt.Println(sumofNumbers(50,60,70,80,90))
} */

// #u gave multiple subjects so to print one after other u shud use for loop

package main
import "fmt"

func printDetails(student string, subjects ...string) {
	fmt.Println("hey " , student, ", here are ur subjects -")
	for _,sub := range subjects { // used blank identifier
		fmt.Println("%s ,", sub)

	}

}

func main() {
	printDetails("rana", "biology", "math", "science")
}

/* ######################### blank identfies '_' ###############################
 blank identifier is the single underscore operator. It is used to ignore the values that are returned by functions.
ex:
package main
import "fmt"
func f() (int, int ){
	return 42, 53
}
func main() {
	a,b := f()
	fmt.Println(a,b)
}
o/p
42,53

################### what if weanna print only single value ##################
package main
import "fmt"
func f() (int, int) {
	return 42, 53
}
func main() {
	v,_ := f()
	fmt.Println(v)
}
o/p:
42 // want to return only one value than we use blank identifier
