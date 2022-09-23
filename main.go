package main

import (
	"fmt"
	s "github.com/silabig1294/Learninggo/salary"
	"github.com/silabig1294/Learninggo/addition"
	k "github.com/silabig1294/Learninggo/multiply"
)

type multiply int

func (m multiply) tentimes() int {
	return int(m * 10)
}

func main(){
	//method and object
	var num int
	fmt.Print("Enter any positive integer: ")
	fmt.Scanln(&num)
	mul := multiply(num)
	fmt.Println("Ten times of a given number is: ",mul.tentimes())
	fmt.Println("<=================================>")
	// finish lesson 1

	//test method and object salary
	sal1 := s.Basic(9000.0).Total()
	fmt.Printf("Salary cal for First Employee:\n%.2f\n",sal1)
	fmt.Println(sal1.Hra().Total())
	fmt.Println(sal1.Hra().Total().Salary().Total())
	fmt.Println("000000000000000000000000000000000000")
	sal2 := s.Basic(8000.0).Total()
	fmt.Printf("Salary cal for Second Employee:\n%.2f\n",sal2)
	fmt.Println(sal2.Salary().Total())
	fmt.Println("000000000000000000000000000000000000")
	sal3 := s.Total(9500)
	fmt.Printf("Salary cal for Third Employee:\n%.2f\n",sal3)
	fmt.Println(sal3.Salary().Total())
	fmt.Println("000000000000000000000000000000000000")
	sal4 := s.Salary(9500)
	fmt.Printf("Salary cal for Third Employee:\n%.2f\n",sal4)
	fmt.Println(sal4.Total())
	fmt.Println("000000000000000000000000000000000000")


	// Overloading Method
	var b addition.Multiply = 15
	b.Twice()
	fmt.Println(b)

	var a addition.Addition = 15
	a.Twice()
	fmt.Println(a)

	fmt.Println("**************************************")
	//
    fmt.Println("Call by value")
	var mul1 k.Multiply
	mul1.Twice(10)
	fmt.Println(mul1.Display())

	fmt.Println("Call by pointer")	
	mul2 := new(k.Multiply) // mul2 is a pointer
	mul2.Twice(10)
	fmt.Println(mul2.Display())

	fmt.Println("**************************************")





}