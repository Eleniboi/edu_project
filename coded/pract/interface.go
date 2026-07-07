package main


import(
	"fmt"
)


type Animal interface{

	Speak() string
}

func makeAnimalSpeak(a Animal){
	fmt.Println(a.Speak())
}

func (g Goat) Speak() string{

	return g.Name+"Meeeehhh"
}
func (c Cow) Speak() string{

	return c.Name+"Mooooooh"
}

func (m Monkey) Speak() string{

	return m.Name+"kikikiki"
}

type Goat struct{
	Name string
}
type Monkey struct{
	Name string
}

type Cow struct{

	Name string
}


func main() {

	goat := Goat{Name: "meeahh"}
	monkey := Monkey{Name: "kikiki"}
	cow := Cow{Name: "mooooh"}

	makeAnimalSpeak(goat)
	makeAnimalSpeak(monkey)
	makeAnimalSpeak(cow)
}