package testsession

import session "forum/sessions"

var _ session.Repository = (*TestSession)(nil)
