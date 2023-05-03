package appProxy

import "main/internal/proxy"

func Run() {
	proxy.RunProxy("localhost:9000", []string{"localhost:8080", "localhost:8082"})
}
