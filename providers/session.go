package providers

import "github.com/gofiber/session/v2"

var sessionP *session.Session

func SessionProvider() *session.Session {
	return sessionP
}

func SetUpSessionProvider(session *session.Session) {
	sessionP = session
}
