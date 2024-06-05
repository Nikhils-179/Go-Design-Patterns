package main

import "fmt"

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	return nil
}

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) setWindowType() {
	n.windowType = "Wooden Window"
}

func (n *NormalBuilder) setDoorType() {
	n.doorType = "Wooden Door"
}

func (n *NormalBuilder) setNumFloor() {
	n.floor = 2
}

func (n *NormalBuilder) getHouse() House {
	return House{
		doorType:   n.doorType,
		windowType: n.windowType,
		floor:      n.floor,
	}
}

type House struct {
	windowType string
	doorType   string
	floor      int
}

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func main() {
	normalbuilder := getBuilder("normal")
	director := newDirector(normalbuilder)
	director.setBuilder(normalbuilder)
	normalHouse := director.buildHouse()
	fmt.Println("Normal house door type : ",normalHouse.doorType)
	fmt.Println("Normal House windw type : " ,normalHouse.windowType)
	fmt.Println("Normal house number floor ", normalHouse.floor)
}


// Product: The complex object that is being constructed. : House here
// Builder Interface: Defines the steps required to build the product. : Ibuilder
// Concrete Builder: Implements the steps defined in the Builder interface. NormalbUilder implements IBuilder
// Director: Constructs the product using the Builder interface. Director orchestrates constructionn process .It uses a builder t construct object step by step
// Client Code: Uses the Director to get the final product.