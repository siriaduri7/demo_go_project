// syntax if condition {
	//execute when condition is true // } else {
		//executes when condition is false }
//f this condition is true,
//Note, for the syntax in if-else block, the else statement should start from the same line where the if ends.
//statements inside the body of if are executed.If the condition is false,the statements inside the body of if are not executed.
/* ############## IF STATEMENT ################################
package main
import "fmt"
func main() {
	var a string ="happy"
	if a == "happy"{
		fmt.println(a)
	}
}
###################### IF_ELSE STATEMENT ################################## */
/* 
package main
import "fmt"
func main () {
	var fruit string ="orange" // here it took reference as fruit in var and in if block also
	if fruit == "apple"
	fmt.println( "fruit is apple" )
} else {
	fmt.println("fruit is not apple")
}

###################### MULTIPLE ELSE IF CONDITION ###################
SYNTAX
 if condition_1 {
	//execute when condition_1 is true
 } else if condition_2 {
	//execute when consition_1 is false, and condition_2 is true //
 } else if condition_3 {
	//execute when condition_1 and 2 are false and Condition_3 is true //
 } else {
	// when none of the baove conditions are true//
 } */

 package main
 import "fmt"
 func main() {
	fruit := "grapes"
	if fruit == " apple" {
		fmt.Println(" i love apple")
	} else if fruit == " oranges" {
		fmt.Println (" oranges are not apples")
	} else {
		fmt.Println("no appetite")
	}
 }