package models

import "time"

type User struct {
	//ID
	ID int64 `json:"id"`

	//Name
	Name string `json:"name"`

	//Birth
	Birth time.Time `json:"birth"`

	//Active
	Active bool `json:"active"`

	//Status code
	Status_code int `json:"status_code"`

	//Created_at
	Created_at time.Time `json:"created_at"`

	//Modified_at
	Modified_at time.Time `json:"modified_at"`
}
