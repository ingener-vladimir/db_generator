package repository

const (
	countUsersQuery    = "select count(*) from logins_info"
	insertUserQuery    = "insert into logins_info (email, password) values (?, ?)"
	insertAccountQuery = "insert into users (login_id, name, surname, birthday, sex, hobby, city, avatar) values (?, ?, ?, ?, ?, ?, ?, ?)"
)
