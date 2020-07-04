package validate

import (
	"fmt"
	"testing"
)

type User struct {
	Id    int    `validate:"number,min=1,max=1000"`
	Name  string `validate:"string,min=2,max=10"`
	Bio   string `validate:"string"`
	Email string
}

func TestValidate(t *testing.T) {
	user := User{
		Id:    30,
		Name:  "121212",
		Bio:   "11",
		Email: "foobar",
	}

	err := ValidateStruct(user)
	fmt.Printf("%s\n", err.Error())
}
