package models

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	description string `json:"description"`
}

var Students = make(map[int]Student)

func init() {
	Students[1] = Student{Id: 1, Name: "Lakshay Garg", Age: 20, Email: "Lakshay@yopmail.com", description: "Hi from lakshay"}
	Students[2] = Student{Id: 2, Name: "Vaibhav", Age: 22, Email: "Vaibhav@yopmail.com", description: "Hi from vaibhav"}
}
