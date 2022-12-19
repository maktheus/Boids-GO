package main

import (
	"fmt"
	"math"
)

// Point é uma estrutura que representa uma coordenada x, y em um plano 2D.
type Point struct {
	x float64
	y float64
}

// Boid é uma estrutura que representa um boid (pássaro simulado).
type Boid struct {
	position Point  // Posição atual do boid no plano 2D.
	velocity Point  // Velocidade atual do boid no plano 2D.
}

// updatePosition atualiza a posição do boid de acordo com sua velocidade.
func (b *Boid) updatePosition() {
	b.position.x += b.velocity.x
	b.position.y += b.velocity.y
}

// distance retorna a distância euclidiana entre dois boids.
func (b *Boid) distance(other *Boid) float64 {
	dx := b.position.x - other.position.x
	dy := b.position.y - other.position.y
	return math.Sqrt(dx*dx + dy*dy)
}

// align ajusta a velocidade do boid de acordo com a velocidade média dos outros boids próximos.
func (b *Boid) align(neighbors []*Boid) {
	var sum Point
	count := 0

	// Encontra a velocidade média dos vizinhos próximos.
	for _, neighbor := range neighbors {
		if b.distance(neighbor) < 10 {
			sum.x += neighbor.velocity.x
			sum.y += neighbor.velocity.y
			count++
		}
	}

	if count > 0 {
		sum.x /= float64(count)
		sum.y /= float64(count)

		// Ajusta a velocidade do boid para se aproximar da velocidade média dos vizinhos.
		b.velocity.x = (b.velocity.x + sum.x) / 2
		b.velocity.y = (b.velocity.y + sum.y) / 2
	}
}

// separate ajusta a velocidade do boid para se afastar dos outros boids próximos.
func (b *Boid) separate(neighbors []*Boid) {
	var sum Point
	count := 0

	// Encontra a direção média de afastamento dos vizinhos próximos.
	for _, neighbor := range neighbors {
		if b.distance(neighbor) < 50 {
			diff := Point{b.position.x - neighbor.position.x, b.position.y - neighbor.position.y}
			distance := math.Sqrt(diff.x*diff.x + diff.y*diff.y)
			sum.x += diff.x / distance
			sum.y += diff.y / distance
			count++
		}
	}

	if count > 0 {
		sum.x /= float64(count)
		sum.y /= float64(count)

		// Ajusta a velocidade do boid para se afastar da direção média dos vizinhos.
		b.velocity.x = (b.velocity.x + sum.x) / 2
		b.velocity.y = (b.velocity.y + sum.y) / 2
	}
}

// cohesion ajusta a velocidade do boid para se aproximar da posição média dos outros boids próximos.
func (b *Boid) cohesion(neighbors []*Boid) {
	var sum Point
	count := 0

	// Encontra a posição média dos vizinhos próximos.
	for _, neighbor := range neighbors {
		if b.distance(neighbor) < 100 {
			sum.x += neighbor.position.x
			sum.y += neighbor.position.y
			count++
		}
	}

	if count > 0 {
		sum.x /= float64(count)
		sum.y /= float64(count)

		// Calcula a direção para se aproximar da posição média dos vizinhos.
		direction := Point{sum.x - b.position.x, sum.y - b.position.y}

		// Ajusta a velocidade do boid para se aproximar da posição média dos vizinhos.
		b.velocity.x = (b.velocity.x + direction.x) / 2
		b.velocity.y = (b.velocity.y + direction.y) / 2
	}
}

// flock atualiza a velocidade do boid de acordo com o comportamento de enxame.
func (b *Boid) flock(neighbors []*Boid) {
	b.align(neighbors)
	b.separate(neighbors)
	b.cohesion(neighbors)
}

func main() {
	// Cria uma lista de boids.
	boids := []*Boid{
		&Boid{position: Point{x: 100, y: 100}, velocity: Point{x: 1, y: 1}},
		&Boid{position: Point{x: 200, y: 100}, velocity: Point{x: 1, y: 1}},
		&Boid{position: Point{x: 300, y: 100}, velocity: Point{x: 1, y: 1}},
		&Boid{position: Point{x: 400, y: 100}, velocity: Point{x: 1, y: 1}},
	}

	// Executa o enxame de boids por um determinado número de passos.
	for i := 0; i < 100; i++ {
		// Atualiza a velocidade de cada boid de acordo com o comportamento de enxame.
		for _, b := range boids {
			b.flock(boids)
		}

		// Atualiza a posição de cada boid de acordo com sua velocidade.
		for _, b := range boids {
			b.updatePosition()
		}

		// Imprime a posição atual de cada boid.
		for _, b := range boids {
			fmt.Printf("Boid na posição (%f, %f)\n", b.position.x, b.position.y)
		}
	}
}
