package main

import "fmt"

type Employee struct { 														//structure to store employee details
	Index      int
	Name       string
	Age        int
	Salary     int
	Department string
}

type Department struct {												 	//structure to store dept name and employee Name
	Name      string
	Employees []Employee
}

type Departments []Department												//SLICE TO STORE MULTIPLE DEPARTMENTS

func (d *Departments) AddNewDepartment(name string) {
	if name == "" {															//Check for edge cases where ndeoartment name is empty
		fmt.Println("Department name cannot be empty")
		return
	}

	newDept := Department{
		Name:      name,
		Employees: []Employee{},
	}
	*d = append(*d, newDept)
	fmt.Println("Department added successfully")
}

func (d *Departments) AddEmployee(name string, age int, salary int, deptName string) {
	if name == "" {															// Edge cases to check invalid data entered for name , ge or salary
		fmt.Println("Name cannot be empty")
		return
	}
	if age <= 0 {
		fmt.Println("Age cannot be less than zero")
		return
	}
	if salary <= 0 {
		fmt.Println("Salary must be greater than zero")
		return
	}

	found := false
	for i, dept := range *d {
		if dept.Name == deptName {											//this line checks whether the department entered by user exists in department slice or not
			index := len(dept.Employees) + 1
			newEmployee := Employee{
				Index:      index,
				Name:       name,
				Age:        age,
				Salary:     salary,
				Department: deptName,
			}
			(*d)[i].Employees = append((*d)[i].Employees, newEmployee)
			fmt.Println("Employee added successfully")
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Department not found")
	}
}

func (d Departments) DisplayInfo() {
	if len(d) == 0 {														//IN CASE NO DATA IS PRESENT
		fmt.Println("No departments to show")
		return
	}

	for _, dept := range d {
		fmt.Printf("\nDepartment: %s\n", dept.Name)
		if len(dept.Employees) == 0 {
			fmt.Println("No employees in this department")
			continue
		}
		for _, emp := range dept.Employees {
			fmt.Printf("Index: %d || Name: %s || Age: %d || Salary: %d || Department: %s\n",
				emp.Index, emp.Name, emp.Age, emp.Salary, emp.Department)
		}
	}
}

func (d *Departments) UpdateSalary() {									//TO UPDATE SALARY WITHOUT PERCENTAGE
	var empName string
	var newSalary int

	fmt.Println("\nEnter the employee name to update salary")
	fmt.Scanln(&empName)
	fmt.Println("\nEnter the new salary")
	fmt.Scanln(&newSalary)

	found := false
	for i, dept := range *d {
		for j, emp := range dept.Employees {
			if emp.Name == empName {
				(*d)[i].Employees[j].Salary = newSalary
				fmt.Println("Salary updated successfully")
				found = true
				break
			}
		}
	}
	if !found {
		fmt.Println("Employee not found")
	}
}

func (d *Departments) GiveRaise() {										// TO GIVE SALARY HIKE BY CERTAIN PERCENT
	var empName string
	var percentage float64

	fmt.Println("\nEnter the employee name for raise :")
	fmt.Scanln(&empName)
	fmt.Println("\nEnter raise percentage :")
	fmt.Scanln(&percentage)

	found := false
	for i, dept := range *d {
		for j, emp := range dept.Employees {
			if emp.Name == empName {
				newSalary := float64(emp.Salary) * (1 + percentage/100)
				(*d)[i].Employees[j].Salary = int(newSalary)
				fmt.Printf("Gave %s a %.2f%% raise. New salary: %d\n",
					empName, percentage, (*d)[i].Employees[j].Salary)
				found = true
				break
			}
		}
	}
	if !found {
		fmt.Println("Employee not found")
	}
}

func main() {
	var (
		name     string
		age      int
		salary   int
		deptName string
		choice   int
	)

	depts := Departments{}

	for {
		fmt.Println("*****************************************************************")
		fmt.Println("Welcome to Employee Management System")
		fmt.Println("=======> 1. Add New Department")
		fmt.Println("==========> 2. Add Employee")
		fmt.Println("==============> 3. Display All Info")
		fmt.Println("=================> 4. Update Salary")
		fmt.Println("====================> 5. Give Raise")
		fmt.Println("=======================> 6. Exit")
		fmt.Println("*****************************************************************")
		fmt.Print("Enter your choice (1-6): ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\nEnter department name")
			fmt.Scanln(&deptName)
			depts.AddNewDepartment(deptName)

		case 2:
			fmt.Println("\nEnter employee name")
			fmt.Scanln(&name)
			fmt.Println("\nEnter age")
			fmt.Scanln(&age)
			fmt.Println("\nEnter salary")
			fmt.Scanln(&salary)
			fmt.Println("\nEnter department name")
			fmt.Scanln(&deptName)
			depts.AddEmployee(name, age, salary, deptName)

		case 3:
			depts.DisplayInfo()

		case 4:
			depts.UpdateSalary()

		case 5:
			depts.GiveRaise()

		case 6:
			fmt.Println("Exiting")
			return

		default:
			fmt.Println("Invalid choice. Please enter a valid option (1-6).")
		}

		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln()											//FOR CLI TO LOOK TIDY 
	}
}
