package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"ejemplo.com/api/api/db"
	"ejemplo.com/api/api/models"
	"github.com/gorilla/mux"
)

// Obtener todos los productos
func GetProductos(w http.ResponseWriter, r *http.Request) {
	var productos []models.Producto
	result, err := db.DB.Query("SELECT * FROM Producto")
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close()

	for result.Next() {
		var producto models.Producto
		err := result.Scan(&producto.Id, &producto.Nombre, &producto.Precio, &producto.Stock)
		if err != nil {
			log.Fatal(err)
		}
		productos = append(productos, producto)
	}

	json.NewEncoder(w).Encode(productos)

}

// Obtener un producto por su ID
func GetProducto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var producto models.Producto
	result := db.DB.QueryRow("SELECT * FROM Producto WHERE id = ?", params["id"])
	err := result.Scan(&producto.Id, &producto.Nombre, &producto.Precio, &producto.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(producto)
}

// Crear nuevo producto
func CreateProductos(w http.ResponseWriter, r *http.Request) {
	var producto models.Producto
	json.NewDecoder(r.Body).Decode(&producto)
	result, err := db.DB.Exec("INSERT INTO Producto (nombre, precio, stock) VALUES (?, ?, ?)",
		producto.Nombre, producto.Precio, producto.Stock)
	if err != nil {
		log.Fatal(err)
	}

	producto.Id, _ = result.LastInsertId()
	json.NewEncoder(w).Encode(producto)
}

// Actualizar un producto existente
func UpdateProducto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var producto models.Producto
	json.NewDecoder(r.Body).Decode(&producto)
	_, err := db.DB.Exec("UPDATE Producto SET nombre = ?, precio = ?, stock = ?, WHERE id = ?",
		producto.Nombre, producto.Precio, producto.Stock, params["id"])
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(producto)
}

// Eliminar un producto existente
func DeleteProducto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_, err := db.DB.Exec("DELETE FROM Producto WHERE id = ?", params["id"])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusNoContent)
}
