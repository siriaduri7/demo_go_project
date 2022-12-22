// ########### ARTHEMATIC OPERATOR###############
package main
import "fmt"
func main() {
	var a,b int = 2, 4
	fmt.Println( a+b)
}

	/*//float
	var a, b float64 =79.02, 75.66
	fmt.println("%.2f", a-b)

	// division : The division operator returns the quotient, when left operand is divided by the right operand.

	var a,b =24,2
	fmt.Println(a/b)
	o/p 12

	// modules :  the modulus operator returns the remainder when left operand is divided by the right operand.

 var a,b =24,7
	fmt.Println(a%b)
	o/p 3
	When we divide 24 by 7, we get a remainder as 3 

	################# INCREMENT(++)/ UNARY OPERATOR ##############
	Unary operators are operators that act upon a single operand to produce a new value.
This operator increments the value of the operand by one.
 
var i int =1
i++
fmt.Println(i)

###########################LOGICAL OPERATOR ################

Logical operators are used to determine the logic between variables or values.
Some common logical comparisons hat are used in boolean statements include, are two variables both true?
Does either of the two expressions evaluate to true?,
logical and (&&)
logical or(||)
logical NOT(!)

var x int =10
fmt.Println((x < 100)&& (x<200)) t t (t)
fmt.Println((x<30)) && (x<0))  f f (f)
o/p
true
false

logical or(||)
var x int =10
fmt.Println((x < 0) || (x<200)) // t f (t)
fmt.Println((x<0)) || (x>200))  // f f (f)
o/p
true
false

NOT(!) : its unary  operator
var x,y int = 10,20
fmt.println(!(x>y))
fmt.println(!(true))
fmt.println(!(false))

o/p true
false true
######################## ASSIGNMENT OPERATOR###############

= assign
+= add and assign
-= subtract and assign
*= multiply and assign
/= divide and assign quotient
%= sivide and assign modulus


