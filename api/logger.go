package api

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

// Logger logger interface.
type Logger interface {
	LogRequest(method, url string, values url.Values)
	LogResponse(method, url string, statusCode int, json []byte)
}

// DefaultLogger default logger.
type DefaultLogger struct {
}

// LogRequest logs request.
func (l *DefaultLogger) LogRequest(method, url string, values url.Values) {
	c := len(values)

	if c == 0 {
		log.Printf("%s %s", method, url)
		return
	}

	rows := make([]string, c)

	i := 0
	for k, v := range values {
		// TODO: check v[0] is correct
		rows[i] = fmt.Sprintf("%s: %s", k, v[0])
		i++
	}

	log.Printf("%s %s\n%s", method, url, strings.Join(rows, "\n"))
}

// LogResponse logs response.
func (l *DefaultLogger) LogResponse(method, url string, statusCode int, json []byte) {
	log.Printf("%s %s %d\n%s", method, url, statusCode, string(json))
}
