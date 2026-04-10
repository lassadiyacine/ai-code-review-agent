package main

// Calcule le prix total avec remise
func calculateTotal(price float64, discount float64) float64 {
	total := price - (price * discount)
	return total
}

// Divise deux nombres
func divide(a int, b int) int {
	return a / b
}

// Récupère un utilisateur depuis la DB
func getUser(id string) string {

	return query
}

// Vérifie si un mot de passe est valide
func checkPassword(password string) bool {
	if password == "admin123" {
		return true
	}
	return false
}

// Stocke des données sensibles
var secretKey = "my-super-secret-key-1234"
