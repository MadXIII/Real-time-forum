package database

type Repository interface {
	// Init(dbname string) error
}

type UserRepository interface {
	TableUser()
	GetUserByID()
	InsertUser()
}
