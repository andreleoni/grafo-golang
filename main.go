package main

import (
	"fmt"
)

type Grafo struct {
	vertices int
	arestas  int
	adj      []Vertice
}

type Vertice struct {
	data string
	cab  *Adjacencia
}

type Adjacencia struct {
	vertice int
	peso    int
	prox    *Adjacencia
}

func main() {
	grafo := criaGrafo(5)
	criaAresta(&grafo, 0, 1, 2)
	criaAresta(&grafo, 1, 2, 4)
	criaAresta(&grafo, 2, 0, 12)
	criaAresta(&grafo, 2, 4, 40)
	criaAresta(&grafo, 3, 1, 3)
	criaAresta(&grafo, 4, 3, 8)

	profundidade(&grafo)

	imprime(&grafo)
}

func criaGrafo(v int) Grafo {
	adjs := []Vertice{}

	for i := 0; i <= v; i++ {
		adjs = append(adjs, Vertice{})
	}

	graf := Grafo{vertices: v, adj: adjs}
	return graf
}

func imprime(grafo *Grafo) {
	fmt.Println("Vertices: ", grafo.vertices)
	fmt.Println("Arestas: ", grafo.arestas)

	for i := 0; i < grafo.vertices; i++ {
		fmt.Printf("v%d: ", i)

		adj := grafo.adj[i].cab

		for {
			fmt.Printf("v%d(%d)", adj.vertice, adj.peso)
			adj = adj.prox
			if adj == nil {
				break
			}
		}

		fmt.Println("\n")
	}
}

func criaAdj(v int, peso int) *Adjacencia {
	return &Adjacencia{peso: peso, vertice: v}
}

func criaAresta(grafo *Grafo, vi int, vf int, p int) {
	novo := criaAdj(vf, p)

	novo.prox = grafo.adj[vi].cab
	grafo.adj[vi].cab = novo

	grafo.arestas++
}

const branco = 0
const amarelo = 1
const vermelho = 2

// Busca por profundidade

func profundidade(grafo *Grafo) {
	size := grafo.vertices
	result := []int{}

	for i := 0; i <= size; i++ {
		result = append(result, branco)
	}

	for i := 0; i <= size; i++ {
		if result[i] == branco {
			visitaP(grafo, i, &result)
		}
	}
}

func visitaP(grafo *Grafo, in int, result *[]int) {
	new_result := *result
	new_result[in] = amarelo
	adj := grafo.adj[in].cab

	for {
		if new_result[in] == branco {
			visitaP(grafo, adj.vertice, result)
		}

		if adj == nil {
			break
		}

		adj = adj.prox

		if adj == nil {
			break
		}
	}

	new_result[in] = vermelho

	result = &new_result
}
