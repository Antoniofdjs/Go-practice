package main

import (
	"errors"
	"fmt"
	"os"
)


func talk(p person) {
	p.speak()
}

func (u User) speak() {
	fmt.Println("Hello from User!")
}
// func (u student) speak() {
// 	fmt.Println("Hello from Student!")
// }
func div(num1 int, num2 int)(int, error){
	
	if num2==0{
		err:= errors.New("Cannot devide by 0")
		return 0, err
	}
	return num1/num2, nil
}
type person interface{
	speak()
}
type User struct{
	name string
	age int
}

type student struct{
	name string
}


func main (){
	
	
	result, err :=div(10, 3)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	
	fmt.Println(result)
	usuario:= User{
		name:"Jona",
		age:24,
	}
	studiante:= student{
		name: "Pepe",
	}
	fmt.Printf("NAME: %v AGE: %v\n", usuario.name, usuario.age)
	talk(usuario)
	talk(studiante)

		
	}
