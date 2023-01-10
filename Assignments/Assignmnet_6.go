

package main
import "fmt"
type person struct {
	first_name string
	last_name string
	favorite_country []string
}

type vehicle struct{
	doors string
	color string
}

type truck struct { 
	fourWheel bool
	truck_vehicle vehicle // vehicle embedded in truck like this
}

type sedan struct {
	luxary bool
	sedan_vehicle vehicle
}

type square struct {
	length int
}

type circle struct {
	radius int
}

type SHAPE inteface {
	Area() float64
}


//create a type SHAPE that defines an interface as anything that has the AREA method



func main() {
	fmt.Println("1st question")

// p: person {"john, "snow" , {}}
	first := person {
		first_name : "John",
		last_name : "snow",
		favorite_country : []string{"USA", " Australia", " Italy"},
	}
	fmt.Println("First person struct", first )
	fmt.Println("person first name", first.first_name)
	fmt.Println("Looping over the favourite Country")

	second := person {
		first_name : "Tyrion",
		last_name : "Lannister",
		favorite_country : []string{ "Australia", "turkey", "Dubai"},

	}
	
	first.firstperson()
	fmt.Println("second person struct", second )
	fmt.Println("person first name", second.first_name)
	fmt.Println(second.last_name)
	fmt.Println("Looping over the favourite Country")
	second.Secondperson()

	fmt.Println ("second question")
	//map_name := make(map[key_datatype]value_data_type)

	a := make(map[string]person)
	a[first.last_name] = first
	a[second.last_name] = second
	fmt.Println(a) 
	for i,j := range a {
		fmt.Println("person struct", a[i])
		//fmt.Println("person first name", a[i].first)
		//// fmt.Println(a[i].firstname, a[i].lastname)
		fmt.Println("looping over the favourite country")
		for x :=range j.favorite_country {
			fmt.Println(x,j.favorite_country[x])
		}
	}

	fmt.Println("3rd question")
	ford := truck{
		fourWheel : false,
		truck_vehicle : vehicle{ "twodoor", "grey"},

	}

	toyota := sedan {
		luxary : true,
		sedan_vehicle : vehicle { "fourdoor", "black"},
	}
	fmt.Println(ford)
	//fmt.Println(ford.vehilce.twowheel)
	//fmt.Println(ford.twodoor)
	fmt.Println(toyota)
	fmt.Println(toyota.sedan_vehicle)
	//fmt.Println(toyota.sedan_vehicle.color)

	fmt.Println( "4th question")
	sr := SQUARE{3}
    cr := CIRCLE{2}
    fmt.Print("SQUARE AREA : ")
    INFO(sr)
    fmt.Print("CIRCLE AREA : ")
    INFO(cr)
	

}

func (first *person) firstperson() {
	
	for i, s := range first.favorite_country {
		fmt.Println(i,s)
	}
	
	
}
func (second *person) Secondperson() {
	
	for i, s := range second.favorite_country {
		fmt.Println(i,s)
	}
}

	func (sr sqaure) AREA() float32 {
		return int(sr.length * sr.width)
	}

	func (cr circle) AREA() float32 {
		return PI * float32(cr.radius * cr.radius)
	}
	func INFO(Sh SHAPE) float64 {
		fmt.Println((sh.AREA()))
		return sh.AREA()
	}


	
/* output:
1st question
First person struct {John snow [USA  Australia  Italy]}
person first name John
Looping over the favourite Country
0 USA
1  Australia
2  Italy
second person struct {Tyrion Lannister [Australia turkey Dubai]}
person first name Tyrion
Lannister
Looping over the favourite Country
0 Australia
1 turkey
2 Dubai
second question
map[Lannister:{Tyrion Lannister [Australia turkey Dubai]} snow:{John snow [USA  Australia  Italy]}]
person struct {John snow [USA  Australia  Italy]}
looping over the favourite country
0 USA
1  Australia
2  Italy
person struct {Tyrion Lannister [Australia turkey Dubai]}
looping over the favourite country
0 Australia
1 turkey
2 Dubai
3rd question
{false {twodoor grey}}
{true {fourdoor black}}
{fourdoor black}  */
type SQUARE struct {
    length int
}
type CIRCLE struct {
    radius int
}
type SHAPE interface {
    AREA() float32
}
func Hello() string {
    return "Hello"
}
func INFO(shape SHAPE) float32 {
    fmt.Println(shape.AREA())
    return shape.AREA()
}
func (square SQUARE) AREA() float32 {
    return float32(square.length * square.length)
}
func (circle CIRCLE) AREA() float32 {
    return PI * float32(circle.radius*circle.radius)
}
var mapvariable map[string]person = make(map[string]person)
const PI float32 = 3.14
func main() {
    Hello()
    person1 := person{"John", "Snow", []string{"USA", "Austrilia", "Italy"}}
    person2 := person{"Tyrion", "Lannister", []string{"Austria", "Turkey", "Dubai"}}
    //1
    fmt.Println("First person struct ", person1)
    fmt.Println("Person First Name ", person1.first_name)
    fmt.Println("Looping over the favourite Country")
    for i, v := range person1.fav_country {
        fmt.Println(i, " ", v)
    }
    fmt.Println("Second person struct ", person2)
    fmt.Println("Person First Name ", person2.first_name)
    fmt.Println("Looping over the favourite Country")
    for i, v := range person2.fav_country {
        fmt.Println(i, " ", v)
    }
    fmt.Println("------------------------------------")
    //2
    mapvariable[person1.last_name] = person1
    mapvariable[person2.last_name] = person2
    for _, v := range mapvariable {
        fmt.Println(v)
    }
    fmt.Println("-------------------------------------")
    //3
    truck1 := truck{vehicle{2, "Orange"}, true}
    sedan1 := sedan{vehicle{4, "Black"}, true}
    fmt.Println("Truck1 Details : ", truck1)
    fmt.Println("Sedan Details : ", sedan1)
    fmt.Println("Truck1 is of color : ", truck1.vehicle.color)
    fmt.Println("Sedan1 number of doors : ", sedan1.vehicle.doors)
    fmt.Println("-------------------------------------")
    //4
    square := SQUARE{3}
    circle := CIRCLE{2}
    fmt.Print("SQUARE AREA : ")
    INFO(square)
    fmt.Print("CIRCLE AREA : ")
    INFO(circle)
    fmt.Println("-------------------------------------")
    //5
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter the String:")
    inp1, _ := reader.ReadString('\n')
    inp11 := strings.TrimSpace(inp1)
    fmt.Println("input is :", inp11)
    input := strings.Split(inp11, " ")
    //input := strings.SplitAfter(inp11, " ")
    fmt.Println(input)
    var m1 map[string]int = make(map[string]int)
    for _, v := range input {
        //key := strings.TrimSpace(v)
        //fmt.Println(i, " ", key, " before", m1[key])
        if m1[v] == 0 {
            m1[v] = 1
        } else {
            m1[v] = m1[v] + 1
        }
        //fmt.Println(i, " ", key, " after", m1[key])
    }
    fmt.Println("Ocuurances of strings:", m1 
