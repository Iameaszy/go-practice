package main

import (
	"fmt"
	"time"
)


type Employee struct { 
	ID int 
	Name, Address, Position string 
	DoB time.Time 
	Salary int 
	ManagerID int 
}
var employees []Employee;
func main() {
	var yusuf Employee;

	yusuf.Name = "Adeniyi Yusuf"
	yusuf.Position = "Software Engineer"
	employees = append(employees, yusuf);
	getEmployee(employees, 0).Name = "Olasunkanmi"
	fmt.Println("get employee", 	getEmployee(employees, 0).Name)
	fmt.Printf("yusuf %#v, employess: %#v", yusuf, employees);
}


func getEmployee(employees []Employee, id int) *Employee{
	var employee Employee
	for _, value := range employees{
		if value.ID == id {
			employee = value
		} 
	}
	return &employee;
}
