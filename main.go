package main 

import(
 "fmt" 
) 

func main(){
 
	add := Add(5, 10)	
 	fmt.Println("Addition: ", add) 

}


func Add(a, b int)int{
	sum := 0

	sum = a+b

	return sum
}