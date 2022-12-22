/* //syntax : switch expression {
case value_1:
	//execute when expression equals to value_1
case value_2:
	//execute when expression equals to value_2
default:
	//execute when no match is found
}
################## SWITCH ##############
// This switch statement tests the value of an expression and compares it with multiple cases. Once the case match is found, a
block of statements associated with that particular case is executed.
Each case in a block of a switch has a different name, or a number, or a value, which is referred to as an identifier. 
package main
import ="fmt"
func main () {
	var i int =100 // (2) ex : var i int =800
	switch i {
	case 10:
		fmt.Println("i is 10")
	case 100,200 :
		fmt.Println("i is either 100 or 200")
	default:
		fmt.Println("i is neither 0,100 or 200")
	}
}
o/p : i is either 100 or 200
(2)i is neither 0,100 or 200
############################################### fall-through keyword ##################################################
We also have a very special keyword for a switch-case statement and that is called the fall-through keyword.
// The fall-through keyword is used in switch-case, to force the execution flow to fall through the successive case block. */
package main
import "fmt"
func main() {
	var i int =10
	switch i {
	case -5:
		fmt.Println("-5")
	case 10:
		fmt.Println("10") // so fallthrough means a case nd kinda case till fallthrough keyword lenianta varku it will print
		fallthrough // it continuous  until it finds a case, which does not consist of the fallthrough keyword.
	case 20:
		fmt.Println("20")
		fallthrough
	default:
		fmt.Println("default")
	}
}

/* ############# switch case giving expressions in case value #############
// 
In this case, we do not need to mention an expression after switch. We can simply write our conditions inside the case blocks 
itself.
According to our interpretation, below both the cases evaluated are true, This is because Go uses an implicit break statement 
for each case. Over here, a case block with the condition a+b==30 was executed, because it was the first one in the order.

package main
import "fmt"
func main() {
	var a,b int = 10,20
	switch {
	case a+b == 30:
		fmt.Println("equal to 30")
	case a+b <=30:
		fmt.Println("less than or equal to 30")
	default:
		fmt.Println("greater than 30")
	}
}