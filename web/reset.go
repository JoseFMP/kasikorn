package web

// GetNew provides a new instance of the web portal session.
// It reuses the same configuration. The goal of this function is to avoid to need
// to pass the configuration of the session again if we just want to have a clea session
// with the same configuration
func (session *Session) GetNew() *Session {
	newSession := InitSession(*session.initialConfig)
	return newSession
}
