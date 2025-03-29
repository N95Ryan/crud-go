package inventaire

import "fmt"

// Inventory stocke les produits
type Inventory struct {
	Products []Product
}

// AddProduct ajoute un produit à l’inventaire
func (inv *Inventory) AddProduct(p Product) {
	inv.Products = append(inv.Products, p)
	fmt.Println("Produit ajouté :", p.Name)
}

// UpdateStock met à jour la quantité d’un produit existant
func (inv *Inventory) UpdateStock(name string, quantity int) {
	for i := range inv.Products {
		if inv.Products[i].Name == name {
			inv.Products[i].Quantity += quantity
			fmt.Printf("Stock mis à jour : %s -> %d unités\n", name, inv.Products[i].Quantity)
			return
		}
	}
	fmt.Println("Produit non trouvé :", name)
}

// DisplayInventory affiche tous les produits
func (inv Inventory) DisplayInventory() {
	fmt.Println("\n📦 Inventaire :")
	for _, p := range inv.Products {
		p.Display()
	}
}
