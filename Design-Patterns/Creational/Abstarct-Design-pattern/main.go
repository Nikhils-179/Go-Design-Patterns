package main

import "fmt"

func main() {
	var factory GUIFactory
	fmt.Println("Give os type")
	var input string
	fmt.Scanln(&input)
	if input == "windows"{
		factory = WindowsFactory{}
	}else if input == "mac"{
		factory = MacosFactory{}
	}

	button := factory.CreateButton()
	Checkbox := factory.CreateCheckbox()

	button.Paint()
	Checkbox.Paint()
}

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type Button interface {
	Paint()
}

type Checkbox interface {
	Paint()
}

type WindowsButton struct{} // implemets Button

func (wb WindowsButton) Paint() {
	fmt.Println("Paiting a windows button")
}

type MacOSButton struct{} //implements Button

func(mb MacOSButton) Paint(){
	fmt.Println("Painitng a MacOS Button")
}

type WindowsCheckBox struct{} // implememts Windowscheckbox

func (wc WindowsCheckBox) Paint(){
	fmt.Println("Paiitng a windows checkbox")
}

type MacOSCheckBox struct{} //implememts MacOs checkbox

func(mb MacOSCheckBox) Paint(){
	fmt.Println("Painitinng a macos checkbox")
}

type WindowsFactory struct{} //implements GUI factory

func (wf WindowsFactory) CreateButton() Button{
	return WindowsButton{} //encapsulationn happenig here Widows factory method CreateButto called upon WindowsButton
}

func (wf WindowsFactory) CreateCheckbox() Checkbox{
	return WindowsCheckBox{}
}

type MacosFactory struct{} // implememts MacOsFactory

func (mf MacosFactory) CreateButton() Button{
	return MacOSButton{}
}

func (mf MacosFactory) CreateCheckbox() Checkbox{
	return MacOSCheckBox{}
}
