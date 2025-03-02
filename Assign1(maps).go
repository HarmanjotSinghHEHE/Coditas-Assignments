package main

import (
	"fmt"
	"regexp"		//FOR CHECKING IF THE EMAIL IS VALID 
)

type Person struct {
	Index  int // Unique identifier as map key
	Name   string
	Age    int
	Gender string
	Email  string
	Ph_No  int64
}

type Persons map[int]Person 	//WE WILL USE INDEX AS THE KEY

func (p *Persons) AddNewPerson(name string, age int, gender string, email string, phNo int64) {
	// DEFINING EMAIL PATTERN
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// Input validation
	if name == "" {
		fmt.Println("Name Can Not be Empty\n")
		return
	}
	if age <= 0 {
		fmt.Println("Age Cannot be less than Zero")
		return
	}
	if phNo <= 0 {
		fmt.Println("Enter a Valid Phone No ")
		return
	}
	// FOR CHECKING IF THE EMAIL ENTERED IS CORRECT OR NOT
	if !emailRegex.MatchString(email) {
		fmt.Println("Invalid Email Format")
		return
	}

	if *p == nil {
		*p = make(map[int]Person)
	}

	index := len(*p) + 1 // THIS GENERATES NEXT RATE

	newPerson := Person{
		Index:  index,
		Name:   name,
		Age:    age,
		Gender: gender,
		Email:  email,
		Ph_No:  phNo,
	}

	(*p)[index] = newPerson
	fmt.Printf("Person added with Index: %d\n", index)
}

func (p Persons) Introduce() {
	if len(p) == 0 {
		fmt.Println("No Entries To Show")
		return
	}

	for _, person := range p {
		fmt.Printf(
			"Index: %d || Name: %s || Age: %d || Gender: %s || Email: %s || Phone Number: %d\n",
			person.Index,
			person.Name,
			person.Age,
			person.Gender,
			person.Email,
			person.Ph_No,
		)
	}
}

func (p *Persons) UpdateName() {
	var index int
	var newName string
	fmt.Println("\nEnter the Index of the person whose name you want to update")
	fmt.Scanln(&index)
	fmt.Println("\nEnter the new Name")
	fmt.Scanln(&newName)

	if person, exists := (*p)[index]; exists {
		if newName == "" {
			fmt.Println("Name Can Not be Empty")
			return
		}
		person.Name = newName
		(*p)[index] = person
		fmt.Println("Name updated successfully")
	} else {
		fmt.Println("\n404 Not Found")
	}
}

func (p *Persons) UpdateAge() {
	var index int
	var newAge int
	fmt.Println("\nEnter the Index of the person whose Age you want to update")
	fmt.Scanln(&index)
	fmt.Println("\nEnter the person's new Age")
	fmt.Scanln(&newAge)

	if person, exists := (*p)[index]; exists {
		if newAge <= 0 {
			fmt.Println("Age Cannot be less than Zero")
			return
		}

		person.Age = newAge
		(*p)[index] = person
		fmt.Println("Age updated successfully")
	} else {
		fmt.Println("\n404 Not Found")
	}
}

func (p Persons) CheckVote() {
	var index int
	fmt.Println("\nEnter the Index of the person to check voting eligibility")
	fmt.Scanln(&index)

	if person, exists := p[index]; exists {
		if person.Age >= 18 {
			fmt.Println("\nEligible for Voting")
		} else {
			fmt.Println("\nNot Eligible")
		}
	} else {
		fmt.Println("\n404 Not Found")
	}
}

func main() {
	var (
		Name   string
		Age    int
		Gender string
		Email  string
		Ph_No  int64
		key    int
	)

	persons := Persons{}

	for {
		fmt.Println("*****************************************************************Welcome*****************************************************************")
		fmt.Println("=======> 1. Add a Person")
		fmt.Println("==========> 2. Introduce/Display")
		fmt.Println("==============> 3. Update Name")
		fmt.Println("=================> 4. Update Age")
		fmt.Println("====================> 5. Check Voting Eligibility")
		fmt.Println("=======================> 6. Exit")
		fmt.Println("************************************************************Enter your choice index number*****************************************************************")
		fmt.Scanln(&key)

		switch key {
		case 1:
			fmt.Println("\nEnter the Name")
			fmt.Scanln(&Name)
			fmt.Println("\nEnter the Age")
			fmt.Scanln(&Age)
			fmt.Println("\nEnter the Gender as MALE OR FEMALE")
			fmt.Scanln(&Gender)
			fmt.Println("\nEnter the Email")
			fmt.Scanln(&Email)
			fmt.Println("\nEnter the Phone Number")
			fmt.Scanln(&Ph_No)
			persons.AddNewPerson(Name, Age, Gender, Email, Ph_No)
		case 2:
			persons.Introduce()
		case 3:
			persons.UpdateName()
		case 4:
			persons.UpdateAge()
		case 5:
			persons.CheckVote()
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option (1-6).")
		}

		fmt.Println("\nPress Enter to continue")
		fmt.Scanln()
	}
}
