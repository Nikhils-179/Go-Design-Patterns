package main

import (
	"fmt"
	"sync"
)

type Singletonn struct{}

var innstance *Singletonn
//lazyway
func Getinnstance() *Singletonn {
	if innstance == nil {
		innstance = &Singletonn{}
	}
	return innstance
}

//Eager way , instance created at variable declaration
var instance2 = &Singletonn{}

func Getinnstance2() *Singletonn{
	return instance2
}


var instance3 *Singletonn

var once sync.Once

func Getinnstance3() *Singletonn{
	once.Do(func (){
		instance3 = &Singletonn{}
	})
	return instance3
}



func main() {
	singleton1 := Getinnstance()
	singletonn2 := Getinnstance()

	fmt.Println(singleton1 == singletonn2) //lazy way

	singleton11 := Getinnstance2()
	singleton22 := Getinnstance2()

	fmt.Println(singleton11 == singleton22) //eager way

	singleton111 := Getinnstance3()
	singleton222 := Getinnstance3()

	fmt.Println(singleton111 == singleton222) //Thread safe way

}