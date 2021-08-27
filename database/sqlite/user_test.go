package sqlite

import (
	"forum/database"
)

var _ database.UserRepository = (*User)(nil)
var _ database.Repository = (*Store)(nil)
