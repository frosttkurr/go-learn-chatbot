package main

import (
	"learn-chatbot/controllers"
	"learn-chatbot/initializers"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	http.ListenAndServe(":8888", http.HandlerFunc(controllers.Handler))
}
