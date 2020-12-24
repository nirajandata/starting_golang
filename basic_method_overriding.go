package main
import "fmt"

type Human struct{
 name string
 reg_number int
}
type Student struct {
 Human
 school string
} 
func (h *Human) Hey() (string,string){
 return "the number is ",h.name
} 
func (e *Student) Hey() (string,string){
 return "oe k xa ",e.name
}
func main(){
 nir:=Human{"nirajan",434344}
 stud:=Student{Human{"nvraj",54534},"pokhariya"}
 fmt.Println(nir.Hey())
 fmt.Println(stud.Hey())
}
