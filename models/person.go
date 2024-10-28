package models

import ()

type Person struct {
	ID uint		     `json:"person_id"`
	Name string      `json:"name"`
	Email string     `json:"email"`
	Age uint         `json:"age"`
}

type Persons []Person