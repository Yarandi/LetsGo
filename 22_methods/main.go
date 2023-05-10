package main

import "fmt"

func main(){
	hamed := User{"Hamed", "oldemail@g.c", true, 35}
	hamed.GetStatus()
	hamed.GetAge()
	hamed.NewMail()
	//we see its not override to original struct data
	println(hamed.Email)
}

type User struct{
	Name string 
	Email string
	Status bool
	Age int
}

//method definition to getting user status
func (u User) GetStatus() {
	fmt.Println(u.Status)
}

//method definition to getting user age
func (u User) GetAge() {
	fmt.Println(u.Age)
}

//method definition to manipulate email
//the point is this new mail is a copy not nanipulated the original one
func (u User) NewMail() {
	u.Email = "new@gmail.com"
	fmt.Println(u.Email)
}