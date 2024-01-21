package handler

import (
	"net/http"

	"github.com/elsif-maj/umbraSearch/app"
)

func SetupRoutes(server *app.Server) {
	http.HandleFunc("/api/snippets", app.MakeAPIFunc(server.HandleSnippets))       // Placeholder/Test route
	http.HandleFunc("/api/snippets/new", app.MakeAPIFunc(server.HandleNewSnippet)) // Placeholder/Test route
}
