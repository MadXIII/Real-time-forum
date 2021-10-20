package testdb

import "forum/database"

var _ database.Repository = (*TestDB)(nil)
