package testsession

import session "forum/utils/sessions"

var _ session.Repository = (*TestSession)(nil)
