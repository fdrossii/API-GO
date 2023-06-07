package main

import (
	"log"
	"net/http"

	"ejemplo.com/api/api/db"
	"ejemplo.com/api/api/services"
	"github.com/gorilla/mux"
)

//Estructura para el modelo de producto

type Producto struct {
	Id     int64   `json:"id"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
	Stock  int     `json:"stock"`
}

func main() {

	db.Connection()

	//Inicialización del enrutador
	router := mux.NewRouter()

	// Rutas de la API
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server ON"))
	})
	router.HandleFunc("/productos", services.GetProductos).Methods("GET")
	router.HandleFunc("/productos/{id}", services.GetProducto).Methods("GET")
	router.HandleFunc("/productos", services.CreateProductos).Methods("POST")
	router.HandleFunc("/productos/{id}", services.UpdateProducto).Methods("PUT")
	router.HandleFunc("/productos/{id}", services.DeleteProducto).Methods("DELETE")

	//Iniciar el servidor en el puerto 8080
	log.Fatal(http.ListenAndServe(":8080", router))

	//Cerrar conexión con la BD
	defer db.DB.Close()
}
