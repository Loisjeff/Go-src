package  main

import "fmt"

const name = "local"

const VERSION = 1.1

func  prlntName(){
	fmt.Println(name)
}
func init()  {
	fmt.Println("init")

}
