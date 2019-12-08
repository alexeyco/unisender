package unisender

import (
	"fmt"
	"log"
	"strings"
)

type Logger interface {
	LogRequest(method, url string, data map[string]string)
	LogResponse(method, url string, statusCode int, json []byte)
}

type DefaultLogger struct {
}

func (l *DefaultLogger) LogRequest(method, url string, data map[string]string) {
	c := len(data)

	if c == 0 {
		log.Printf("%s %s", method, url)
		return
	}

	rows := make([]string, c)

	i := 0
	for k, v := range data {
		rows[i] = fmt.Sprintf("%s: %s", k, v)
		i++
	}

	log.Printf("%s %s\n%s", method, url, strings.Join(rows, "\n"))
}

func (l *DefaultLogger) LogResponse(method, url string, statusCode int, json []byte) {
	log.Printf("%s %s %d\n%s", method, url, statusCode, string(json))
}
