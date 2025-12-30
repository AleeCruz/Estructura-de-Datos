package areas

import (
	"math" // 1. "math" en minúscula
)

func AreaCirculo(radio float32) float32 {
	// 2. Todo debe ser float64 para poder multiplicarse
	return math.Pi * cuadrado(radio)
}

// Si quieres mantener tu función cuadrado:
func cuadrado(n float32) float32 {
	return n * n
}
