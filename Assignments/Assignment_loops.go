 //Using the for loop print 1 to 100 numbers
package main
import "fmt"
func main() {
	for i :=1 ; i <=100 ; i++{
	fmt.Println(i)
	}

/*Create a for loop using this syntax
     for condition { }
print out the odd numbers in 1 to 50. */


	
	
	fmt.Println("odd numbers in 1 to 50 :")
	var oddNum int = 50
	
	for i :=1; i<=oddNum; i++{
		if(i%2 != 0) {
		fmt.Println(i)
		}
	}
/*Create a for loop using this syntax
    for { }
print out the even numbers in 1 to 50. */
fmt.Println("even numbers in 1 to 50")
	var eveNum int = 50
	
	for a := 1; a<= eveNum ; a++ {
		if (a%2 ==0) {
		fmt.Println(a)
	}
}

/*fmt.Println("Fourth problem:")
	for i := 50; i < 106; i++ {
		fmt.Println(i, i/6)
	}
*/
fmt.Println("The quotient for each number between 50 and 105")
for p :=50 ; p < 106 ; p++ {
	fmt.Println(p,p/6)
	}





//Print out the quotient for each number between 50 and 105 when it is divided by 6.

//. Read the user input. If the input is equal to Golang tutorial print welcome else print end

var first string 
fmt.Println("Enter the input: ")
fmt.Scanln(&first)
if(first == "Golangtutorial"){
fmt.Println("print welcome")
} else {
fmt.Println("print end")
}
/*fmt.Println("Fifth problem:")
	y := "Golang tutorial"
	fmt.Println("enter string")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	x := scanner.Text()

	if x == y {
		fmt.Println("Welcome")
	} else {
		fmt.Println("end")
	} */


fmt.Println("6th program solution :")
for i := 1 ; i <=80 ; i++ {
	if i %2 ==0 && i %4 ==0 {
		fmt.Println("golang tutorial")
	} else if i%2 ==0 {
		fmt.Println("golang")
	} else if i%4 ==0 {
		fmt.Println("tutorial")

	} else {
		fmt.Println(i)
	}
}
}
	

//
/*var number = int(1)
for number <= 80{
	scratch := ""
	if number%2 == 0{
		scratch += "Golang"

	}
	if number%4 == 0{
		scratch += "tutorial"
	}
	if (number%2 ==0 && number%4 ==0 ){
		scratch += "Golangtutorial"

	}
	if scratch != "" {
		fmt.Println("%s \n", scratch)
	} else {
		fmt.Println("%d \n", number)
	}
	number++
}
} */



/*
func main() {
    str1 := "Abc"
    str2 := "Abc"
    if strings.Compare(str1, str2) == 0 {
        fmt.Println("The two strings are equal.")
    } else {
        fmt.Println("The two strings are not equal.")
    }
} 

6. Write a program to print the numbers from 1 to 80.
But, for multiples of two print Golang instead of the number and for the multiples of four print tutorial.
For numbers which are multiples of both two and four print Golang tutorial. */


/*var number = int(1)

func main() {
	for number <= 100 {
		scratch := ""
		if number%3 == 0 {
			scratch += "Fizz"
		}
		if number%5 == 0 {
			scratch += "Buzz"
		}
		if scratch != "" {
			fmt.Printf("%s\n", scratch)
		} else {
			fmt.Printf("%d\n", number)
		}
		number++
	}
}
*/

    
