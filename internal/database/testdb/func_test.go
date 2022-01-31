package testdb

import "forum/internal/database"

var _ database.Repository = (*TestDB)(nil)
