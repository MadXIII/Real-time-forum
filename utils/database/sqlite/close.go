package sqlite

//Close ...
func (s *Store) Close() {
	s.db.Close()
}
