package inventaire

import "fmt"

// Product représente un produit dans l'inventaire
type Product struct {
	Name     string
	Price    float64
	Quantity int
}

// Display affiche les informations d’un produit
func (p Product) Display() {
	fmt.Printf("Produit : %s\nPrix : %.2f €\nStock : %d unités\n\n", p.Name, p.Price, p.Quantity)
}
