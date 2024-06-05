package main

import ("fmt"
		"errors")

type vehicle interface {
	Drive() string
}

type car struct{}

func (c *car) Drive() string {
	return "I am driving car"
}

type bike struct{}

func(b *bike) Drive()string{
	return "I am driving bike"
}

type Plane struct{}

func(p *Plane) Drive()string{
	return "I am driving plane"
}

func vehiclefactory(x string) (vehicle ,error){
	fmt.Println("Vehcicle type is : " , x)
	switch x {
	case "car" : return &car{} , nil
	case "bike" : return &bike{} , nil
	case "plane" : return &Plane{} , nil
	default : return nil , errors.New("unknown type")
	}
}
func main() {
	var input string
	fmt.Println("Write the vehicle type")
	fmt.Scanln(&input)

	result , err := vehiclefactory(input) 
	if err!= nil || result == nil {
		panic("Error occured doesnot exist")
		
	}else{
		fmt.Println(result.Drive())
	}

}