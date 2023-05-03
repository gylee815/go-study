package main

import (
	"fmt"
)

type PasswordError struct {
	Len         int
	RequiredLen int
}

func (err PasswordError) Error() string {
	return "len of password you set is too short"
}

func CheckPwLen(pw string) error {
	if len(pw) < 8 {
		return PasswordError{len(pw), 8}
	} else {
		return nil
	}
}

func SignUp(id string, pw string) (bool, error) {
	err := CheckPwLen(pw)
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			return false, fmt.Errorf("Error msg: %w, Requried len: %d, Current len: %d", errInfo, errInfo.RequiredLen, errInfo.Len)
		} else {
			return false, nil
		}
	} else {
		return true, nil
	}
}

func main() {
	if rst, err := SignUp("test", "test"); rst {
		fmt.Println("Sign up succeed")
	} else {
		fmt.Println(err)
	}
}
