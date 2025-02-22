package main

import (
	"csv/constant"
	"fmt"
)

func main() {
	student := constant.Stu{}
	student.Name = "Tom"
	student.Age = 18

	fmt.Println(student)
}
