package models

import (
    "github.com/robfig/revel"
)

type User struct {
    Password                string
    Email                   string
}

// validate login details
func (user *User) Validate(v *revel.Validation) {
    v.Required(user.Email == "e@d.a").Message("Incorrect!")
    v.Email(user.Email)

    v.Required(user.Password == "plm").Message("Incorrect!")
    v.MinSize(user.Password, 1)
}