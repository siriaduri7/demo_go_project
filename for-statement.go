//A loop is a sequence of instructions that is continually repeated until a certain condition is reached.
/* syntax : 
for initialization; condition; post {
	//statements
}
/This initialization statement is optional and it executes before the for loop starts. It is mostly used in simple statements 
like variable declarations or increment or assignment statement.
//The condition statement holds a Boolean expression, which is evaluated at the starting of each iteration of the loop.
If the value of the conditional statement is true, then the loop will execute.
//post : After the post statement, the condition statement evaluates again. If the value of the conditional statement is false, then the loop ends.
The post statement in a previous example was incrementing the value of i, which happened at the end of every loop.

################# EXAMPLE ############### */
/*package main
import "fmt"
func main() {
	for i :=1 ; i <=5 ; i++ {
	fmt.Println(i*i)
	}
}
*/
package main
import "fmt"
func main(){
	i := 1 //The initialization statement is not included in the loop.
	for i <=5 {
		fmt.Println(i*i)
		i+=1
	}
}

/* ################### INFINITE LOOP ###############
Now, if we skip the condition , we get an infinite loop. An infinite loop is a looping construct that does not terminate the 
loop and executes the loop forever.

package main
import "fmt"
func main() {
	sum := 0
	for {
		sum ++ // repeated forvever
	}
	fmt.Println(sum) // never reached
}

############################# BREAK STATEMENT ###################
The break statement ends the loop immediately when it is encountered.
this is because in the third iteration, we encountered the break statement, which made us completely exit the entire for loop structure.

package main
import "fmt"
func main() {
	for i :=1 ; i <= 5; i++{
		if i ==3 {
		break
		}
		fmt.Println(i)
	}
	
}

####################### CONTINUE STATEMENT ################
The continuous statement skips the current iteration of loop and continues with the next iteration.
Again, in this case, instead of break, we've used continue.
#When i became three and if condition was encountered and the continue statement was executed, it simply skipped the fmt.Println
statement for i=3

package main
import "fmt"
func main() {
	for i := 1; i <=5 ; i++ {
		if ==3 {
			continue       // it will skip 3 to print o/p 1245
		}
		fmt.Println(i)
	}
}

