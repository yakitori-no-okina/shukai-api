package action

type Context interface {
	Param(key string) string
	JSON(code int, obj interface{}) error
	Bind(obj interface{}) error
}
