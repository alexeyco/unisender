package api

type Logger interface {
	Println(message string, params map[string]interface{})
}
