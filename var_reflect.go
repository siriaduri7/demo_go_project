//########## using reflect.TypeOf() #############
//First, we would need to import the packages. FMT for formatting and printing our output and 
//reflect package for using the type of function.
package main
import (
	"fmt"
	"reflect"
)
func main() {
	var grades int =43
	var message string ="hello"
	fmt.Printf("variable grades: %v is of type %v \n", grades, reflect.TypeOf(grades))
	fmt.Printf("variable message : %v is of type %v \n", message, reflect.TypeOf(message))

	/*fmt.Printf("Type : %v \n",reflect.TypeOf(1000))
	fmt.Printf("Type : %v \n",reflect.TypeOf("sri"))
	fmt.Printf("Type : %v \n",reflect.TypeOf(46.0))
	fmt.Printf("Type : %v\n", reflect.TypeOf(true))*/
	// ###############use type of method with variables#############

}


/*################################ FIND TYPE OF VARIABLE  ########################
how can you find the data type of that variable?////
%T FORMAT SPECIFIER (prints datatype of variable)
*) reflect.TypeOf function from the reflect package. 
ex:
*/
/*package main
import ("fmt")
func main() {
	var grades int =43
	var message string = "hello world"
	var isCheck bool =true
	var amount float32 =5644.54
	
	fmt.Printf("variable grades : %v is of type %T \n" , grades, grades)
	fmt.Printf("variable message : %v is if type of %T \n", message, message)
	fmt.Printf("variable isCheck: %v is of type of %T \n", isCheck, isCheck)
	fmt.Printf("variable amount: %v is of type of %T \n", amount, amount)
}

o/p : variable gardes :43 is of type int
*/