package main

import (
	"crud-go/inventaire"
	"fmt"
)

func main() {
	fmt.Println("💼 Gestion d'Inventaire en Go")

	// Initialisation de l'inventaire
	inv := inventaire.Inventory{}

	// Ajout de produits
	p1 := inventaire.Product{Name: "Ordinateur", Price: 999.99, Quantity: 5}
	p2 := inventaire.Product{Name: "Clavier", Price: 49.99, Quantity: 10}
	p3 := inventaire.Product{Name: "Souris", Price: 19.99, Quantity: 15}

	inv.AddProduct(p1)
	inv.AddProduct(p2)
	inv.AddProduct(p3)

	// Affichage de l'inventaire
	inv.DisplayInventory()

	// Mise à jour du stock
	inv.UpdateStock("Clavier", 5)
	inv.UpdateStock("Souris", -3)

	// Affichage de l'inventaire après mise à jour
	inv.DisplayInventory()
}
