package controllers

import (
	"ecommerce_go/config"
	"ecommerce_go/utils"
	"net/http"
	"strconv"
)

// Estructura del producto
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CreatedAt   string  `json:"created_at"`
}

// Obtener productos
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, description, price, stock, created_at FROM products")
	if err != nil {
		http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt)
		if err != nil {
			http.Error(w, "Error al leer los productos", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error al obtener los resultados", http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, http.StatusOK, products)
}

// Añadir nuevo producto
func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		stock := r.FormValue("stock")

		// Insertar nuevo producto en la base de datos
		db := config.GetDB()
		_, err := db.Exec("INSERT INTO products (name, description, price, stock) VALUES (?, ?, ?, ?)", name, description, price, stock)
		if err != nil {
			http.Error(w, "Error al añadir el producto", http.StatusInternalServerError)
			return
		}

		// Redirigir a la página de admin
		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
	}
}

// Actualizar producto
func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Obtener los datos del formulario
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		stock := r.FormValue("stock")

		// Convertir el ID del producto a entero
		productID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "ID de producto inválido", http.StatusBadRequest)
			return
		}

		// Actualizar el producto en la base de datos
		db := config.GetDB()
		_, err = db.Exec("UPDATE products SET name=?, description=?, price=?, stock=? WHERE id=?", name, description, price, stock, productID)
		if err != nil {
			http.Error(w, "Error al actualizar el producto", http.StatusInternalServerError)
			return
		}

		// Redirigir al dashboard de productos
		http.Redirect(w, r, "/admin/products", http.StatusSeeOther)
		return
	}
}

// Función para eliminar un producto
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	db := config.GetDB()
	_, err := db.Exec("DELETE FROM products WHERE id = ?", productID)
	if err != nil {
		http.Error(w, "Error al eliminar el producto", http.StatusInternalServerError)
		return
	}

	// Redirigir al dashboard de productos
	http.Redirect(w, r, "/admin/products", http.StatusSeeOther)
}
