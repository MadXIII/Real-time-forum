package testsession

import session "forum/internal/sessions"

var _ session.Repository = (*TestSession)(nil)
