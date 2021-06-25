package data

// test data
var users = []User{
	{Name: "Peter Jones",
		Email:    "peter@gmail.com",
		Password: "Peter_pass",
	},
	{
		Name:     "John Smith",
		Email:    "john@gmail.com",
		Password: "john_pass",
	},
	{
		Name:     "Magareth Doe",
		Email:    "magareth@gmail.com",
		Password: "aka1",
	},
	{
		Name:     "John Doe",
		Email:    "jdoe@gmail.com",
		Password: "aka2",
	},
	{
		Name:     "Utazi Doe",
		Email:    "utazi@gmail.com",
		Password: "aka3",
	},
}

func setUp() {
	ThreadDeleteAll()
	SessionDeleteAll()
	UserDeleteAll()
}
