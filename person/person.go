package person


type Person struct{
	Name string
	Age int
}

func (p Person) Details() string {
	return "My name is " + p.Name + ". I'm " + string(p.Age) + " years old"
}