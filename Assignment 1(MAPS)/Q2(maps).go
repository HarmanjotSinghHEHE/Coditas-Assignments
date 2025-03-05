package main

import "fmt"

type employee struct {
	index      int
	name       string
	age        int
	salary     int
	department string
}

type department struct {
	name      string
	employees map[int]employee
}

type departments map[string]department

func (d *departments) addNewDepartment(name string) {
	if name == "" {
		fmt.Println("Department name cannot be empty")
		return
	}

	if *d == nil {
		*d = make(map[string]department)
	}

	if _, exists := (*d)[name]; exists {
		fmt.Println("Department already exists")
		return
	}

	newDept := department{
		name:      name,
		employees: make(map[int]employee),
	}
	(*d)[name] = newDept
	fmt.Println("Department added successfully")
}

func (d *departments) addEmployee(name string, age int, salary int, deptName string) {
	if name == "" {
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

	if *d == nil {
		fmt.Println("No departments exist. Please add a department first")
		return
	}

	if dept, exists := (*d)[deptName]; exists {
		index := len(dept.employees) + 1
		newEmployee := employee{
			index:      index,
			name:       name,
			age:        age,
			salary:     salary,
			department: deptName,
		}
		dept.employees[index] = newEmployee
		(*d)[deptName] = dept
		fmt.Println("Employee added successfully")
	} else {
		fmt.Println("Department not found")
	}
}

func (d departments) displayInfo() {
	if len(d) == 0 {
		fmt.Println("No departments to show")
		return
	}

	for _, dept := range d {
		fmt.Printf("\nDepartment: %s\n", dept.name)
		if len(dept.employees) == 0 {
			fmt.Println("No employees in this department")
			continue
		}
		for _, emp := range dept.employees {
			fmt.Printf("Index: %d || Name: %s || Age: %d || Salary: %d || Department: %s\n",
				emp.index, emp.name, emp.age, emp.salary, emp.department)
		}
	}
}

func (d *departments) updateSalary() {
	var empName string
	var newSalary int

	fmt.Println("\nEnter the employee name to update salary")
	fmt.Scanln(&empName)
	fmt.Println("\nEnter the new salary")
	fmt.Scanln(&newSalary)

	if newSalary <= 0 {
		fmt.Println("Salary must be greater than zero")
		return
	}

	found := false
	for deptName, dept := range *d {
		for index, emp := range dept.employees {
			if emp.name == empName {
				emp.salary = newSalary
				dept.employees[index] = emp
				(*d)[deptName] = dept
				fmt.Println("Salary updated successfully")
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		fmt.Println("Employee not found")
	}
}

func (d *departments) giveRaise() {
	var empName string
	var percentage float64

	fmt.Println("\nEnter the employee name for raise")
	fmt.Scanln(&empName)
	fmt.Println("\nEnter raise percentage")
	fmt.Scanln(&percentage)

	if percentage < 0 {
		fmt.Println("Percentage cannot be negative")
		return
	}

	found := false
	for deptName, dept := range *d {
		for index, emp := range dept.employees {
			if emp.name == empName {
				newSalary := float64(emp.salary) * (1 + percentage/100)
				emp.salary = int(newSalary)
				dept.employees[index] = emp
				(*d)[deptName] = dept
				fmt.Printf("Gave %s a %.2f%% raise. New salary: %d\n",
					empName, percentage, emp.salary)
				found = true
				break
			}
		}
		if found {
			break
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

	depts := departments{}

	for {
		fmt.Println("*****************************************************************")
		fmt.Println("Welcome to employee management system")
		fmt.Println("=======> 1. Add new department")
		fmt.Println("==========> 2. Add employee")
		fmt.Println("==============> 3. Display all info")
		fmt.Println("=================> 4. Update salary")
		fmt.Println("====================> 5. Give raise")
		fmt.Println("=======================> 6. Exit")
		fmt.Println("*****************************************************************")
		fmt.Print("Enter your choice (1-6): ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\nEnter department name")
			fmt.Scanln(&deptName)
			depts.addNewDepartment(deptName)

		case 2:
			fmt.Println("\nEnter employee name")
			fmt.Scanln(&name)
			fmt.Println("\nEnter age")
			fmt.Scanln(&age)
			fmt.Println("\nEnter salary")
			fmt.Scanln(&salary)
			fmt.Println("\nEnter department name")
			fmt.Scanln(&deptName)
			depts.addEmployee(name, age, salary, deptName)

		case 3:
			depts.displayInfo()

		case 4:
			depts.updateSalary()

		case 5:
			depts.giveRaise()

		case 6:
			fmt.Println("Exiting")
			return

		default:
			fmt.Println("Invalid choice. Please enter a valid option (1-6).")
		}

		fmt.Println("\nPress Enter to continue....")
		fmt.Scanln()
	}
}
