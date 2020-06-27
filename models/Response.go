package models

// Response Struct
type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// ResponseWithToken struct
type ResponseWithToken struct {
	Response Response `json:"response"`
	Token    string   `json:"token"`
}

// ResponseWithCount struct
type ResponseWithCount struct {
	Count    int      `json:"count"`
	Response Response `json:"response"`
}

// Date Struct
type Date struct {
	Day   int    `json:"day"`
	Month string `json:"month"`
	Year  int    `json:"year"`
}

//ResponseWithTokenAndImage ...
type ResponseWithTokenAndImage struct {
	Response Response `json:"response"`
	Token    string   `json:"token"`
	Picture  []byte   `json:"picture"`
}
