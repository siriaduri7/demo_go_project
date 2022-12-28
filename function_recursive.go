/*Recursion is a concept where a function calls itself by direct or indirect means.
The function keeps calling itself until it reaches a base case. Recursion is generally used to solve a problem where the solution
 is dependent on the smaller instances of the same problem.
 */
/*/////////////////////////////////////////////////
 package main
 import "fmt"
 func factorial(n int)int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1) // f(5) = 5* f(4*3*2*1) = 120
}
func main(){
	n :=5
	result := factorial(n)
	fmt.Println("factorial of", n, "is :", result)



 } */
 //o/p factorial of 5 is : 120
 /*

 recursive function
 		factorial (5)
return 5 * factorial(4) = 120
     return 4 * factorial (3) = 24
	    return 3 * factorial(2) = 6
 			return 2 * factorial(1) = 2
 					1

					This is what we return to our main function and we get 120 as the output.

It's the same function, but the execution instance is different.

######################### Anonymous Function ################################
An anonymous function is a function that is declared without any named identifier to refer to it. // no fn name

They can accept inputs and return outputs just as standard functions do. They're generally used for containing functionality
that need not be named and is possibly for a very short-term use.
package main
import "fmt"
func main() {
	x := func(l int, b int)int {
		return l*b

	}
	fmt.Printf("%T \n" , x)
	fmt.Println(x(20,30))
}
o/p

func(int, int) int
600

########################### HIGH ORDER FUNCTIONS #############################
A high order function is a function that receives a function as an argument or returns a function as an output.
why use high order function ?
composition :
We can create smaller functions that take care of certain piece of logic,
and then we can compose more complex functions by using these different smaller functions.
Reduces bugs :
It also makes the code very easy to read and understand.
*/ // ???????????????????????????????? DOUBT HOF ################## CHECK ##############################

package main
import "fmt"

func calcArea(r float64) float64 {
	return 3.14 * r * r // pi*r*r
}

func calcPerimeter(r float64) float64 {
	return 2 * 3.14 *r
}

func calcDiameter(r float64 ) float64 {
	return 2 * r
}

func main() {
	var query int
	var radius float64

//Query basically refers to the query number over here which is one, two, or three for area, perimeter, and diameter as we just learned.
fmt.Println("enter the radius of the circle: ")
fmt.Scanf("%f", &radius)
fmt.Println("enter \n 1-area \n 2- perimeter \n 3- diameter: ")
fmt.Scanf("%d", &query)

if query ==1 {
	fmt.Println("Result: ",calcArea(radius))
} else if quer ==2 {
	fmt.Println("Result: ",calcPerimeter(radius))
} else if query == 3 {
	fmt.Println("Result: ", calcDiameter(radius))
} else {
	fmt.Println("invalid query")
}

}
/* o/p :
enter the radius of teh circle :9.1
enter
1- area
2 - perimeter
3 -diameter 
result : 260.0234
thank you 
*/ //################## using HIGH ORDER FUNCTION ########################
/*func main() {
func printResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Println("Result: ", result)
	fmt.Println("Thank you")
	
}

func getFunction(query int) func(r float64) float64 {

	query_to_func := map[int]func(r float64) float64 {
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return query_to_func[query]
}
}
*/
/*
$################################# DEFER STATEMENT #############################

package main
import "fmt"

func printName(str string) {
	fmt.Println( str)
}
func printRollNum(rno int){
	fmt.Println(rno)
}
func printAddress(adr string) {
	fmt.Println (adr)
}
func main() {
	printName("joe")
	defer printRollNum("77")
	printAdress("street32")
}

o/p
joe
street32
77 // rno has defer func so it will be executed after its surrounding functions call printed at last

 