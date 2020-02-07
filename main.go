// Programmer name: Vincent Gabb
// Date completed:  2/6/2020
// Description: Variable Practice

package main

import "fmt" 

func main() {  

  var food = "Chicken Alfredo"
  
  var name string

  var age int

  fmt.Println("Enter your name here:")

  fmt.Scanln(&name)

  fmt.Println("Welcome to the Jungle" , name)

  fmt.Println("Enter your age here:")

  fmt.Scanln(&age)

  fmt.Println("When I was" , age , "I loved to eat" , food , "all the time.")
}
// end.