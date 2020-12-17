package main
import "fmt"

type Person struct{
 name string 
 age int
}

func odder (p1,p2 Person) (string , int){
if p1.age>p2.age {
 return p1.name , p1.age-p2.age
}

  return p2.name, p2.age-p1.age
 
}

func main(){

nir :=Person{name :"nirey", age:18}
sand :=Person{name:"sandesh",age:19}
old_name,age_gap := odder(nir,sand)
fmt.Printf(" %s is older by %d years",old_name,age_gap)
}
