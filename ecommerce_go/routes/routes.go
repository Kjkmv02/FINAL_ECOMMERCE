package routes

import (
	"ecommerce_go/controllers"
	"net/http"
)

func SetupRoutes() {
	// Rutas de Administrador
	http.HandleFunc("/admin/login", controllers.AdminLoginHandler)              // Página de login del admin
	http.HandleFunc("/admin/dashboard", controllers.AdminDashboardHandler)      // Página principal del admin
	http.HandleFunc("/admin/products", controllers.AdminProductPageHandler)     // Gestión de productos
	http.HandleFunc("/admin/products/add", controllers.AddProductHandler)       // Añadir un producto
	http.HandleFunc("/admin/products/update", controllers.UpdateProductHandler) // Modificar un producto
	http.HandleFunc("/admin/products/delete", controllers.DeleteProductHandler) // Eliminar un producto
	http.HandleFunc("/admin/users/delete", controllers.DeleteUserHandler)       // Eliminar un usuario
	http.HandleFunc("/admin/products/all", controllers.AdminProductPageHandler) // Página de login del admin

	// Rutas de Usuarios
	http.HandleFunc("/users/register", controllers.UserRegisterHandler)      // Página de registro de usuario
	http.HandleFunc("/users/login", controllers.LoginPageHandler)            // Página de login de usuario
	http.HandleFunc("/users/dashboard", controllers.UserDashboardHandler)    // Página principal del usuario
	http.HandleFunc("/users/cart/add", controllers.AddToCartHandler)         // Añadir al carrito
	http.HandleFunc("/users/cart/checkout", controllers.CheckoutCartHandler) // Realizar compra

	// Página principal de la tienda
	http.HandleFunc("/", controllers.HomePageHandler)              // Página de inicio
	http.HandleFunc("/admin/users", controllers.AdminUsersHandler) // Página de ver usuarios

	// Archivos estáticos
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css")))) // CSS
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))    // JS

}
