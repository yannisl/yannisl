package main
  
import (
"fmt"
"strconv" 
)
 
type E interface{}

type myOwnFunction func(words string, param E) E
  
func goodbye(aFunc myOwnFunction, param E) {
    fmt.Println(aFunc ("\\Goodbye", param))
}
  
func hello(aFunc myOwnFunction, param E) {
    fmt.Println(aFunc ("\\Hello", param))
}
  

  
func main() {
    bold := func(words string, param E) E {
        return "\\textbf{" + words +"}"
    	}
    emph := func(words string, param E) E {
    	
    	return "\\emph{" + words + param.(string) + "}"
    }	


    hello(bold, "z")
    goodbye(bold, "y")
  	hello(emph, strconv.Itoa(1))  
} 