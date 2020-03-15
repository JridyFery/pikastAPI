package models



// Response Struct
type Response struct {
	Code	int 		`json:"code"`
	Status	string 		`json:"status"`
	Data	interface{}	`json:"data"`
}

// UserResponse Struct
type UserResponse struct {
	ID					uint				`json:"id"`
	Name				string 				`json:"name"`
	Email				string 				`json:"email"`
	Roles				[]string 			`json:"roles"`
	DateOfBirth 		Date				`json:"dateOfBirth"`
	Country 			string				`json:"countryOfResidence"`
}

// ResponseWithToken struct
type ResponseWithToken struct {
	Response	Response	`json:"response"`
	Token		string		`json:"token"`
}

// ResponseWithCount struct
type ResponseWithCount struct {
	Response	Response		`json:"response"`
	Count		int 			`json:"count"`
}

// Date Struct
type Date struct {
	Day 	int		`json:"day"`
	Month 	string 	`json:"month"`
	Year 	int		`json:"year"`
}