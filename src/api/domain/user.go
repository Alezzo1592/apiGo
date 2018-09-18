package domain

import "time"

type User struct{
	id int64 `json:"id"`
	name string `json: "name"`
	lastName string `json:"last_name"`
	age int64 `json:"age"`
	mail string `json:"mail"`
	registrationDay time.Time `json:"registration_day"`
}