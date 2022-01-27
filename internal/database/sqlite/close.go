package sqlite

//Close - to close db in main.go with defer
func (s *Store) Close() {
	s.db.Close()
}
