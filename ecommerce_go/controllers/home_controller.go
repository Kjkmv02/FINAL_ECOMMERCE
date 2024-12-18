package controllers

import (
	"ecommerce_go/utils"
	"net/http"
)

// Página principal
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "home.html", nil)
}
