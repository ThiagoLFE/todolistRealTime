package web

import (
	"net/http"
	"realtime-todolist/internal/components"
)

func basePage(w http.ResponseWriter, req *http.Request) {
	components.BasePage().Render(req.Context(), w)
}
