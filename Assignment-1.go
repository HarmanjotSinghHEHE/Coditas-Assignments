package main

import "fmt"

type Person struct { 																			 //now creating a structure named person
	Index  int
	Name   string
	Age    int
	Gender string
	Email  string
	Ph_No  int64
}
//some method reciever are called as pointers where i need to change or modify the original slice content. (p *person)
//where as for processes like introduce or check vote person object is directly called. (p person)

type Persons []Person 																			 // A slice to store multiple persons

func (p *Persons) AddNewPerson(name string, age int, gender string, email string, phNo int64) {

	// Conditions to handle edge cases where name or age or Phone no is invalid
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

	ind := len(*p) + 1 																			 // +1 to len of *p to start with index 1
	newPerson := Person{
		Index:  ind,
		Name:   name,
		Age:    age,
		Gender: gender,
		Email:  email,
		Ph_No:  phNo,
	}
	*p = append(*p, newPerson)
	fmt.Println("Person added")
}

func (p Persons) Introduce() {
	if len(p) == 0 {
		fmt.Println("No Entries To Show")
		return
	}
	for _, person := range p {
		fmt.Printf(
			"Index: %d || Name: %s || Age: %d || Gender: %s || Email: %s || Phone Number: %d\n",
			person.Index, 																		 // Index
			person.Name,   																		 // field stores the name data
			person.Age,    																		 // storing the age
			person.Gender, 																		 // storing gender
			person.Email,  																		 // stores email addresses
			person.Ph_No,  																		 // stores phone numbers
		)
	}
}

func (p *Persons) UpdateName() { 																// Method to update a person's name
	var nm string 																				// to take the name from the user, whose name to be changed
	var nn string 																				// to store new name and update it in the slice
	found := false
	fmt.Println("\nEnter the Name you want to update")
	fmt.Scanln(&nm)
	fmt.Println("\nEnter the new Name")
	fmt.Scanln(&nn)

	for i, person := range *p {
		if person.Name == nm {
			fmt.Println("Person found and name will be updated")
			(*p)[i].Name = nn
			found = true
		}
	}
	if !found {
		fmt.Println("\n404 Not Found")
	}
}

func (p *Persons) UpdateAge() { 																// Method to update a person's age
	var nn string  																				// to take the name from the user, whose age you want to change
	var newAge int 																				// to store the new age to update it in the slice
	found := false																				//flag to check if person is found
	fmt.Println("\nEnter the person's name for who you want to update the Age")
	fmt.Scanln(&nn)
	fmt.Println("\nEnter the person's new Age")
	fmt.Scanln(&newAge)

	for i, person := range *p {
		if person.Name == nn {
			(*p)[i].Age = newAge
			found = true
		}
	}
	if !found {
		fmt.Println("\n404 Not Found")
	}
}

func (p Persons) CheckVote() { 																	// Method to check voting eligibility
	var nn string  																				// to input the name for which you want to check if age > 18
	found := false 																				// flag if person not found
	fmt.Println("\nEnter the person's name for who you want to check if he is eligible")
	fmt.Scanln(&nn)
	for _, person := range p {
		if person.Name == nn {
			found = true
			if person.Age >= 18 {
				fmt.Println("\nEligible for Voting")
			} else {
				fmt.Println("\nNot Eligible")
			}
			break 																				// Breaking the loop after checking
		}
	}
	if !found {
		fmt.Println("\n404 Name Not Found")
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
		fmt.Scanln(&key) 																		// for user input of the switch

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

		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln()
	}
}
