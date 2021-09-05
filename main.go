package main

import "fmt"

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
