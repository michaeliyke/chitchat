package main

var _ = users

func users() (users []User, err error) {
	rows, err := Db.Query("SELECT id, name FROM users")
	verify(err)
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Name)
		if !verify(err) {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// var _ = nextUserId

func nextUserId() (id int) {
	// id = 0
	err := Db.QueryRow("SELECT nextval(id) FROM users").Scan(&id)
	ensure(err)
	return
}

var _ = user

func user(id int) (user User, err error) {
	row := Db.QueryRow("SELECT id,name FROM users WHERE id = $1", id)
	err = row.Scan(&user.Id, &user.Name)
	return
}
