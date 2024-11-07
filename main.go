package main

import (
	"fmt"
	"github/poo/models"
)

func main() {

	// Crear un nuevo producto
	nuevoProducto := models.Product{
		ID:          4,
		Name:        "Nombre del Producto",
		Price:       100.0,
		Description: "Descripción del Producto",
		Category:    "C",
	}

	var p models.Product

	// Guardar el nuevo producto
	p.Save(nuevoProducto)

	// Mostrar todos los productos
	fmt.Println("Productos: ", p.GetAll())

	// Mostrar un producto por ID
	fmt.Println("Producto con ID 4: ", p.GetByID(4))

}
