package controllers

import (
	"ecommerce_go/config"
	"ecommerce_go/utils"
	"net/http"
)

// Estructura del usuario (con campo Password añadido)
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"` // Añadido campo Password
	IsAdmin   bool   `json:"is_admin"`
	CreatedAt string `json:"created_at"`
}

// Obtener usuarios
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, email, is_admin, created_at FROM users")
	if err != nil {
		http.Error(w, "Error al obtener usuarios", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.IsAdmin, &user.CreatedAt)
		if err != nil {
			http.Error(w, "Error al leer los usuarios", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error al obtener los resultados", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusOK, users)
}

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Inserta el usuario en la base de datos (no verificamos duplicados aquí)
		db := config.GetDB()
		_, err := db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, password)
		if err != nil {
			http.Error(w, "Error al registrar el usuario", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/users/dashboard", http.StatusSeeOther)
	}
}

func UserDashboardHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "user_dashboard.html", nil)
}

func CheckoutCartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	productID := r.URL.Query().Get("product_id")
	quantity := r.URL.Query().Get("quantity")

	// Actualizar stock
	_, err := config.DB.Exec("UPDATE products SET stock = stock - ? WHERE id = ? AND stock >= ?", quantity, productID, quantity)
	if err != nil {
		http.Error(w, "Error al actualizar stock", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Compra realizada con éxito"})
}

// Página de inicio de sesión
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		utils.RenderTemplate(w, "login.html", nil)
	}
}

// Página de registro
func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		utils.RenderTemplate(w, "register.html", nil)
	}
}
