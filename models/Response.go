package models

import "mime/multipart"

// Response Struct
type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// ResponseWithToken struct
type ResponseWithToken struct {
	Response Response       `json:"response"`
	Token    string         `json:"token"`
	Picture  multipart.File `json:"picture"`
}

// ResponseWithCount struct
type ResponseWithCount struct {
	Response Response `json:"response"`
	Count    int      `json:"count"`
}

// Date Struct
type Date struct {
	Day   int    `json:"day"`
	Month string `json:"month"`
	Year  int    `json:"year"`
}
