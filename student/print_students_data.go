package student

import (
	"fmt"
)

func PrintStudentsData(s Student) {
	fmt.Println("Your student's name is: ", s.Name)
	fmt.Println("Your student's platform is: ", s.Platform)
	fmt.Println("Your student's current lesson is: ", s.CurrentLesson)
	fmt.Println("Your student's classes left is: ", s.ClassesLeft)
}
