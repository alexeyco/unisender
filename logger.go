package unisender

type Logger interface {
	Println(message string, params map[string]interface{})
}
