package testdb

import "forum/utils/database"

var _ database.Repository = (*TestDB)(nil)
