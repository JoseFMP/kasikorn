package config

// Config defines configuration for Kasikorn client
// Remember to not to use clear-text passwords in your code ;)
type Config struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
