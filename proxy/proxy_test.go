package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	nginxServer := newNginxServer()

	appStatusUrl := "/app/status"
	createuserURL := "/create/user"

	httpCode, body := nginxServer.handleRequest(appStatusUrl, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody:%s\n", appStatusUrl, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusUrl, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody:%s\n", appStatusUrl, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusUrl, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody:%s\n", appStatusUrl, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody:%s\n", appStatusUrl, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody:%s\n", appStatusUrl, httpCode, body)
}
