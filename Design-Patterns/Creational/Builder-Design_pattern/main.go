// Create a builder design pattern
// We have a student class with 10 attributes.
// Problem : To initialize the student object we will be needing
// the sequence of the attributes and // to check the data is correct or not
package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
)

type Student struct {
	name         string
	age          int
	sex          string
	grad_year    int
	address      string
	phone_number string
	email        string
	batch        int
}

type StudentBuilder struct {
	name         string
	age          int
	sex          string
	grad_year    int
	address      string
	phone_number string
	email        string
	batch        int
}

func NewStudnetBuilder() *StudentBuilder {
	return &StudentBuilder{}
}

func (b *StudentBuilder) setName(name string) *StudentBuilder {
	b.name = name
	return b
}

func (b *StudentBuilder) setAge(age int) *StudentBuilder {
	b.age = age
	return b
}

func (b *StudentBuilder) setSex(sex string) *StudentBuilder {
	b.sex = sex
	return b
}

func (b *StudentBuilder) setGradYear(year int) *StudentBuilder {
	b.grad_year = year
	return b
}
func (b *StudentBuilder) setAddress(address string) *StudentBuilder {
	b.address = address
	return b
}
func (b *StudentBuilder) setPhone_number(phone string) *StudentBuilder {
	b.phone_number = phone
	return b
}

func (b *StudentBuilder) setEmail(email string) *StudentBuilder {
	b.email = email
	return b
}
func (b *StudentBuilder) setBatch(batch int) *StudentBuilder {
	b.batch = batch
	return b
}

func (b *StudentBuilder) validate() error {
	if b.age < 18 {
		return errors.New("age should be greater than 18")
	}
	if b.grad_year > 2022 {
		return errors.New("under Qualified")
	}
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(b.email) {
		return errors.New("invalid email format")
	}
	if len(b.phone_number) != 10 {
		return errors.New("phone number must have 10 digits")
	}
	return nil
}
func (b *StudentBuilder) Build() (*Student, error) {
	err := b.validate()
	if err != nil {
		return nil, err
	}
	student := &Student{
		name:         b.name,
		age:          b.age,
		sex:          b.sex,
		email:        b.email,
		address:      b.address,
		grad_year:    b.grad_year,
		phone_number: b.phone_number,
		batch:        b.batch,
	}
	return student, nil

}

func main() {
	st, err := NewStudnetBuilder().setName("Nikhil").setAge(19).setEmail("abc.1@co.com").setGradYear(2022).setBatch(123).setSex("Male").setAddress("adnjndfj").setPhone_number("1234567890").Build()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Address %p\n" , st)
}


// Summary:
// - Created a Student struct to represent a student with various attributes.
// - Created a StudentBuilder struct with similar attributes to help in building a Student object.
// - Added setter methods (setName, setAge, etc.) in StudentBuilder to set values for each attribute.
// - Added a validate method in StudentBuilder to check validation rules before creating the object (e.g., age must be >18).
// - Created a Build method in StudentBuilder that calls validate and, if successful, creates and returns a Student object.
