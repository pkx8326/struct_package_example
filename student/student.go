package student

import "fmt"

type Student struct {
	firstname string
	lastname  string
	score     float64
}

func Getdata() (studentdata *Student) { //This function is a struct method
	var studentfirstname, studentlastname string
	var studentscore float64

	fmt.Print("Please input the student's first name :")
	fmt.Scan(&studentfirstname)
	fmt.Print("Please input the student's last name :")
	fmt.Scan(&studentlastname)
	fmt.Print("Please input the student's score :")
	fmt.Scan(&studentscore)

	studentdata = &Student{ //This line is the struct's pointer
		firstname: studentfirstname,
		lastname:  studentlastname,
		score:     studentscore,
	}
	return
}

func (studentdata *Student) Displaydata() {
	fmt.Print("\nThe student's first name is ", (*studentdata).firstname)
	fmt.Print("\nThe student's last name is ", (*studentdata).lastname)
	fmt.Print("\nThe student's score is ", (*studentdata).score)
}

func (data *Student) Cleardata() { //This functions gets a pointer as an argument and return a de-referenced pointer
	fmt.Print("\n\n---Clearing Data---")
	(*data).firstname = "" //de-referenced pointer (value)
	(*data).lastname = ""  //de-referencing a value in a struct like this is not mandatory
	(*data).score = 0
	fmt.Print("\n")
}
