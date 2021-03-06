package main

import "fmt"

type greeter interface{
	greet(string) string
}
type rus struct {}
type eng struct {}

func (r *rus) greet(name string)string{
	return fmt.Sprintf("Привет, %s" , name)
}
func (e *eng) greet(name string)string{
	return fmt.Sprintf("Hello, %s" , name)
}

func sayHello(g greeter,name string){
	fmt.Println(g.greet(name))

}

func main(){
	sayHello(&rus{} , "Петя")
	sayHello(&eng{}, "Petr")
}