package main

import "fmt"

//when to use ponters
// 1. updating state
type User struct {
	name   string
	email  string
	number int32
}

func (u User) Email() string {
	return u.email
}
func Email(user User) string {
	return user.email
}

func (u User) updateEmail(email string) string {
	u.email = email
	return email
}

func (u *User) updateEmailAgain(email string) string {
	u.email = email
	return email
}

func changeValue(str *string) {
	*str = "changed"
}
func changevalue2(str string) {
	str = "changed"
}

func main() {
	fmt.Println("pointers .... ")
	user := User{
		email: "fooobar@g.com",
	}
	user.updateEmail("aggg@bar.com")
	user.updateEmailAgain("yoo@bitch.com")
	fmt.Println(user.Email())

	toChange := "hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

}
