/*/ package contains one or more go sources
// go_lang_projects// is package
packages 2 typesutility
exrvutable
2 condition
the name of the pkg shud be main
2 that it must contain the main() function

lower case can visibl privatee only in particular packagevar is caps 
than it can be exposted it a pbulic variable

// export (1st variable=off)
go mode init
go run main.go
 mod tidy */

package main
import "fmt"

func main() {

	var x int = 200
	var y int = 2
	var c int


	c = x + y
	fmt.Printf("Line 1 - Value of c is %d\n", c )
	c = x - y
	fmt.Printf("Line 2 - Value of c is %d\n", c )
	c = x * y
	fmt.Printf("Line 3 - Value of c is %d\n", c )
	c = x / y 
	fmt.Printf("Line 4 - Value of c is %d\n", c )


}
