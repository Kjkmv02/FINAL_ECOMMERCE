package controllers

import (
	"ecommerce_go/config"
	"ecommerce_go/utils"
	"encoding/json"
	"net/http"
)

// Estructura del carrito
type CartItem struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
}

// Obtener carrito de un usuario
func GetCartHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id") // Obtener el user_id desde la URL
	rows, err := config.DB.Query("SELECT id, user_id, product_id, quantity, created_at FROM cart WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, "Error al obtener el carrito", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cartItems []CartItem
	for rows.Next() {
		var item CartItem
		err := rows.Scan(&item.ID, &item.UserID, &item.ProductID, &item.Quantity, &item.CreatedAt)
		if err != nil {
			http.Error(w, "Error al leer los items del carrito", http.StatusInternalServerError)
			return
		}
		cartItems = append(cartItems, item)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error al obtener los resultados", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusOK, cartItems)
}

// Agregar al carrito
// Agregar al carrito
func AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	var cartItem struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}

	err := json.NewDecoder(r.Body).Decode(&cartItem)
	if err != nil {
		http.Error(w, "Error al decodificar el carrito", http.StatusBadRequest)
		return
	}

	// Aquí deberías agregar la lógica para insertar el producto en la base de datos de carrito
	_, err = config.DB.Exec("INSERT INTO cart (user_id, product_id, quantity) VALUES (?, ?, ?)", 1, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		http.Error(w, "Error al añadir al carrito", http.StatusInternalServerError)
		return
	}

	// Obtener el carrito actualizado
	rows, err := config.DB.Query("SELECT p.name, c.quantity FROM cart c JOIN products p ON c.product_id = p.id WHERE c.user_id = ?", 1)
	if err != nil {
		http.Error(w, "Error al obtener el carrito", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cart []struct {
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
	}

	for rows.Next() {
		var item struct {
			Name     string `json:"name"`
			Quantity int    `json:"quantity"`
		}
		err := rows.Scan(&item.Name, &item.Quantity)
		if err != nil {
			http.Error(w, "Error al leer el carrito", http.StatusInternalServerError)
			return
		}
		cart = append(cart, item)
	}

	utils.RespondJSON(w, http.StatusOK, cart)
}
