package controllers

import (
	"ecommerce_go/config"
	"ecommerce_go/utils"
	"net/http"
)

func AdminProductPageHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener todos los productos de la base de datos
	rows, err := config.DB.Query("SELECT id, name, description, price, stock FROM products")
	if err != nil {
		http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []map[string]interface{}
	for rows.Next() {
		var id, stock int
		var name, description string
		var price float64
		err := rows.Scan(&id, &name, &description, &price, &stock)
		if err != nil {
			http.Error(w, "Error al escanear productos", http.StatusInternalServerError)
			return
		}
		products = append(products, map[string]interface{}{
			"id":          id,
			"name":        name,
			"description": description,
			"price":       price,
			"stock":       stock,
		})
	}

	// Obtener todos los usuarios
	rowsUsers, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, "Error al obtener usuarios", http.StatusInternalServerError)
		return
	}
	defer rowsUsers.Close()

	var users []map[string]interface{}
	for rowsUsers.Next() {
		var id int
		var name, email string
		err := rowsUsers.Scan(&id, &name, &email)
		if err != nil {
			http.Error(w, "Error al escanear usuarios", http.StatusInternalServerError)
			return
		}
		users = append(users, map[string]interface{}{
			"id":    id,
			"name":  name,
			"email": email,
		})
	}

	data := map[string]interface{}{
		"products": products,
		"users":    users,
	}

	utils.RenderTemplate(w, "admin.html", data)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("id")
	_, err := config.DB.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		http.Error(w, "Error al eliminar usuario", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "Usuario eliminado exitosamente"})
}
func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		utils.RenderTemplate(w, "admin_login.html", nil)
	} else if r.Method == "POST" {
		code := r.FormValue("code")
		if code == "000" {
			http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		} else {
			http.Error(w, "Código incorrecto", http.StatusForbidden)
		}
	}
}

// Página del panel de ADMIN
func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Recupera todos los productos de la base de datos
	db := config.GetDB() // Asegúrate de tener la función GetDB() configurada en config/db.go
	rows, err := db.Query("SELECT id, name, description, price, stock FROM products")
	if err != nil {
		http.Error(w, "Error al recuperar los productos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product // Asegúrate de tener una estructura Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
			http.Error(w, "Error al leer los productos", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	// Renderiza la página de admin (admin_dashboard.html) y pasa los productos a la plantilla
	utils.RenderTemplate(w, "admin_dashboard.html", products)
}
func AdminUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener todos los usuarios de la base de datos
	db := config.GetDB()
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, "Error al recuperar los usuarios", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			http.Error(w, "Error al leer los datos del usuario", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Renderizar la plantilla con los datos de los usuarios
	utils.RenderTemplate(w, "admin_users.html", users)
}
