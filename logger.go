package snfiber

type logger interface {
	Error(message any, keyvals ...any)
}
