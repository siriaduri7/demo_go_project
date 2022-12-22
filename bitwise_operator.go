/* 12 =00001100 (In Binary)                   
25 =00011001 (In Binary)	

0 0 0 0 1 1 0 0 							
0 0 0 1 1 0 0 1
____________________
 0 0 0 0 1 0 0 0 = 8 (In Decimal) */
package main
import "fmt"
func main() {
	var a, b int  = 12, 25
	z := a & b
	fmt.Println(z)
}

/* ############# bitwise OR############# f f (f)
0 0 0 0 1 1 0 0 							
0 0 0 1 1 0 0 1  
____________________
0 0 0 1 1 1 0 1 = 29  (In Decimal) 

package main
import "fmt"
func main() {
	var a, b int = 12, 25
	z := a | b
	fmt.Println(z)
}

####################### bitwise XOR ################ 
The result of XOR is 1 if the two bits are opposite. (0) (1) 1
0 0 0 0 1 1 0 0 							
0 0 0 1 1 0 0 1  
____________________						
0 0 0 1 0 1 0 1 = 21 (In Decimal)

package main
import ="fmt"
func main() {
	var a, b int = 12, 25
	z := x (control pina symbol) y
	fmt.Println(z)
}

############### LEFT SHIFT (<<)#####################
It shifts all bits towards left by a certain number of specified bits.
The bit positions that have been vacated by the left shift operator are filled with (0).
212 = 11010100
212 << 1

  1 1 0 1 0 1 0 0
1 1 0 1 0 1 0 0 (0) = 424 (IN DECIMAL) // It shifts all bits towards left by a certain number of specified bits.

package main
import "fmt"
func main() {
	var x int = 212
	z := x << 1
	fmt.Println(z)
}
o/p 424

###################### RIGHT SHIFT (>>) #############
It shifts all bits towards right by a certain number of specified bits.
If we have in our example the number 212, and we want to right shift it by two bits, our output would look something like this.
We've basically shifted all the bits by two bits over here,
212 = 11010100
212 >> 2

  1 1 0 1 0 1 0 0
(00) 1 1 0 1 0 1 53 (IN DCEIMAL) // We've basically shifted all the bits by two bits over here, and all the excess bits to 
the right are discarded.
package main
import "fmt"
func main() {
	var x int = 212
	z := x >> 2
	fmt.Println(z)
}
o/p 53