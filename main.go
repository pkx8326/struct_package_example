package main

import (
	"fmt"

	"example.com/struct_package_example/student"
)

func main() {
	//var studentdata Student
	var studentdata *student.Student //decleare a variable "studentdata" as the struct built in the student package
	var nxoptn string
	for {

		c := menu()
		if c == "1" {
			studentdata = student.Getdata() //getdata() returns a pointer, we need to de-reference it to get its value
			nxoptn = nextoptn()
		} else if c == "2" {
			studentdata.Displaydata()
			nxoptn = nextoptn()
		} else if c == "3" {
			studentdata.Cleardata()
			nxoptn = nextoptn()
		} else {
			nxoptn = "n"
		}
		if nxoptn == "N" || nxoptn == "n" {
			fmt.Println("Goodbye")
			break
		} else {
			continue
		}
	}
}

////////////////////////////////////////////////

func menu() (c string) {
	for {
		fmt.Println("Choose the following actions :")
		fmt.Println("1. Input data")
		fmt.Println("2. Display data")
		fmt.Println("3. Clear data")
		fmt.Println("4. Exit")
		fmt.Print("Your choice :")
		fmt.Scan(&c)

		if c == "1" || c == "2" || c == "3" || c == "4" {
			return
		} else {
			fmt.Println("Invalid choice")
			continue
		}
	}
}

func nextoptn() (nxoptn string) {
	for {
		fmt.Print("\nNext operation? [Y/N] :")
		fmt.Scan(&nxoptn)
		if nxoptn == "Y" || nxoptn == "y" || nxoptn == "N" || nxoptn == "n" {
			return
		} else {
			fmt.Println("Invalid choice.")
			continue
		}
	}
}
