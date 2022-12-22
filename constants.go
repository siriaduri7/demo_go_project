// #### stntax : const <const name> <data type> = <value>
//Constants as the names state, they are variables, whose values once initialized cannot be modified.
/* we have 2 types of constants
1)typed constants

2)untyped constants (here there will be no data type refer exampel )
Constants are untyped, unless they're explicitly given a data type and declaration and they allow for much more flexibility.
An example of how you can declare and initialize an untyped constant is this:
ex: const age =12
const h_name, h_age = "herm", 12

 typed constants: these are constants where you explicitly specify the data type in the declaration, but the flexibility that
  comes with untyped constants is lost here.To create a constant of string data type,
  //simplted types cons will have datatype in syntax
  ex: const name string = "harry"
  const age int =12

  package main
  import ("fmt")
  func main() {
	const i = 25
	conts name = "harry"
	const is_muggle = false
	fmt.printf("%v: %T", i,i)
	fmt.printf("%v: %T", name, name)
	fmt.printf("%v: %T", is_muggle, is_muggle)
  }
  ################EX OF CONST USAGE PROGRAM######## (IMPORTANT #######################################)
  ###############################################################################################
  /* package main
  import "fmt"
  const PI float 64 =3.14 // global constant
  func main() {
	var radius float64 =5.0
	var area float64
	area = PI*radius*radius

  }
  /*
  ############ ONCE INITIALIZED CANNOT CHANGE THE VALUE #######################
  important point to note here, that you cannot declare constant and not initialize it a value and maybe try to assign it a value later on.
You can't do something like this.If you would do this, we'll get an error and the error is going to say that, "Missing value in constant declaration."
  */


  package main
  import  "fmt"
func main(){
	
	const name = "sri"
	// const name
	//name ="ram"
	//const name :="sri" //(above ex in comments)
	fmt.Printf("%v : %T\n", name) // cannot assign new name becoz its a constant cannot chnage
}
/* The short and variable declaration does not work for constants. The compiler will give you the following error

if we try to use the short hand variable assignment operator.It will say, "Unexpected :=" because it was expecting the simple
 equal to or the simple assignment operator. (above ex in comments)*/