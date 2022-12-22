//*************************### how we can print variable and string together.**************#######
// check variable declaration below ##################
package main
import ("fmt")
func main(){
	var name string = "Koddy"
	//var user string = "harry"
	var i int = 79
	//syntax : fmr.printf("template string %s", object args(s))
	//fmt.printf("hey, %v! you have scored %d/100 in chemistry", name, i) . // %v take string ; %d takes int

	fmt.Printf("welcome to %v, and he is %d", name, i)

	//o/p hey, koddy! you have scored 79/100 in chemistry


	//fmt.Println("welcome to \n", name, ",\n ", user)

	//(or) fmt.Print(name, "\n")   //"\n" means new line
	  //    fmt.Print(user) 
	  
	 // (or)fmt.Println(name)
		//fmt.Println(user)
	
//First, we have a string called Welcome to, Then we have a (comma) to add more to the string.Now, we have added a name variable.
//Again, a comma, again, a string, the string consists of ((comma and whitespace)).Again, a (comma) to add more.Then finally, 
//a user variable.


}

/* ###################### we can declare a variable in many ways #####################
1) one is declaring and assigning it a value on the same line.
ex : var user string = "harry"
2)  it's equally valid to declare a variable and then assign its value later.
ex : var user string
      user = "harry"
3) Variables of the same type can be declared and even assigned values on the same lines, such as this.
  var s,t string = "foo" , "bar"
  fmt.println(s)
  fmt.println(t)
4) While if the variables are of different data types, we can declare them like this.
  var (
	s string = "foo"
	i int = 5)

  ########## SHORT VARIABLE DECLARATION ######
 5) In addition to the previous shorthand methods, we have one more shorthand version. It's known as short variable declaration.
 ex :)   s := "hello" // s is variable name 
:= short variable assignment operator It indicates that the short variable declaration is being used.This means that there's
 no need to use the var keyword or declare the variable data type.
 It also means that the value to the right of this operator should be assigned to the variable.Then we have the hello world string [?]