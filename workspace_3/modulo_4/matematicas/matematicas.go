package matematicas

import (
	"math"
)

func AreaCirculo(radio float32) float32 {
	return (math.Pi) * Cuadrado(radio)
}

func Cuadrado(radio float32) float32 {
	return radio * radio
}
