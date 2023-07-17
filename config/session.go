package config

import "github.com/gorilla/sessions"

const SESSION_ID = "session"

var Store = sessions.NewCookieStore([]byte("123asdasdb12983y13nasd"))
