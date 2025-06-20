package main

import "fmt"

type Forma interface {
	Area() float64
	Perimetro() float64
}

type Quadrado struct {
	Lado float64
}

type Circulo struct {
	Raio float64
}

func (q Quadrado) Area() float64 {
	return q.Lado * q.Lado
}

func (c Circulo) Area() float64 {
	return 3.14 * c.Raio * c.Raio
}

func (q Quadrado) Perimetro() float64 {
	return 4 * q.Lado
}

func (c Circulo) Perimetro() float64 {
	return 2 * 3.14 * c.Raio
}

func calcularAreaEPerimetroSomados(f Forma) float64 {
	return f.Area() + f.Perimetro()
}


func main() {
	// Exemplo de uso
	quadrado := Quadrado{Lado: 5}
	circulo := Circulo{Raio: 3}

	fmt.Printf("quadrado %.2f\n", calcularAreaEPerimetroSomados(quadrado))
	fmt.Printf("circulo %.2f\n", calcularAreaEPerimetroSomados(circulo))
}
