package dto

import (
	"fmt"
	"regexp"
)

type ForgetPasswordRequestDto struct {
	Email string `json:"email"`
}

type CheckForgetPasswordRequestDto struct {
	ForgetPasswordRequestDto
	Code string `json:"code"`
}

type SubmitForgetPasswordRequestDto struct {
	CheckForgetPasswordRequestDto
	Password string `json:"password"`
}

type CheckOTPSubAccountRequestDto struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type SubmitRegisterSubAccountRequestDto struct {
	CheckOTPSubAccountRequestDto
	Username string `json:"username"`
	Password string `json:"password"`
}

func (d *ForgetPasswordRequestDto) Validate() error {
	if d.Email == "" {
		return fmt.Errorf("email required")
	}

	return nil
}

func (d *CheckForgetPasswordRequestDto) Validate() error {
	if d.Email == "" {
		return fmt.Errorf("email required")
	}

	if d.Code == "" {
		return fmt.Errorf("code required")
	}

	return nil
}

func (d *SubmitForgetPasswordRequestDto) Validate() error {
	if d.Email == "" {
		return fmt.Errorf("email required")
	}

	if d.Code == "" {
		return fmt.Errorf("code required")
	}

	if d.Password == "" {
		return fmt.Errorf("password required")
	}

	return nil
}

func (d *CheckOTPSubAccountRequestDto) Validate() error {
	if d.Email == "" {
		return fmt.Errorf("email required")
	}

	if d.Code == "" {
		return fmt.Errorf("code required")
	}

	return nil
}

func (d *SubmitRegisterSubAccountRequestDto) Validate() error {
	if d.Email == "" {
		return fmt.Errorf("email required")
	}

	if d.Code == "" {
		return fmt.Errorf("code required")
	}

	if d.Username == "" {
		return fmt.Errorf("username required")
	}

	if d.Password == "" {
		return fmt.Errorf("password required")
	}
	if d.Password != "" {

		secure := true
		tests := []string{".{7,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}
		for _, test := range tests {
			t, _ := regexp.MatchString(test, d.Password)
			if !t {
				secure = false
				break
			}
		}
		if !secure {
			//return fmt.Errorf(`Password must at least 8 characters including at least 3 of the following 4 types of characters: a lower-case letter, an upper-case letter, a number, a special character (such as !@#$%^&*)`)
		}

	}

	return nil
}
