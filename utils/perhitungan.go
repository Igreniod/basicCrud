package utils

func CalculateDiscountedPrice(originalPrice float64, discount int) float64 {
	// Logika perhitungan diskon
	discountPercentage := float64(discount) / 100
	discountAmount := originalPrice * discountPercentage
	discountedPrice := originalPrice - discountAmount

	return discountedPrice
}
