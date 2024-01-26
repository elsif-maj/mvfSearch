package handler

import (
	"net/http"

	"github.com/elsif-maj/umbraSearch/app"
)

func SetupRoutes(server *app.Server) {
	http.HandleFunc("/api/snippets", app.MakeAPIFunc(server.HandleSnippets))
	http.HandleFunc("/api/snippets/new", app.MakeAPIFunc(server.HandleNewSnippet))
	http.HandleFunc("/api/snippets/indexAll", app.MakeAPIFunc(server.HandleIndexAllSnippets))
	http.HandleFunc("/api/snippets/search", app.MakeAPIFunc(server.HandleSearchString))
}
