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
func init() {
	var yusuf Employee;

	fmt.Println()
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

func structEmbedding() {
	type Point struct {
		X, Y int
	}

	type Circle struct {
		Point
		Radius int
	}

	type Wheel struct {
		Circle
		Spokes int
	}

	var w Wheel;
	w.Y = 2
	w.Radius = 3;
}