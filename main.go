package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/IzcoatlRam/puzzle8dfs/node"
	"github.com/IzcoatlRam/puzzle8dfs/stack"
	"github.com/fatih/color"
)

func main() {
	s := node.State{
		{2, 4, 3},
		{1, 5, 6},
		{7, 0, 8},
	}
	withDepth(s)
	fmt.Print("Presiona Enter para continuar con la busqueda por profundidad iterativa...")
	fmt.Scanln()
	withIterativeDepth(s)
}

func DepthFirstSearch(initialState node.State, maxDepth int) *node.Node {
	initialNode := node.NewNode(initialState, "", nil, 0)
	stack := stack.NewStack()
	stack.Push(initialNode)

	for stack.Len() > 0 {
		currentNode := stack.Pop()

		// Detener la búsqueda si se alcanza la profundidad máxima
		if currentNode.Depth > maxDepth {
			continue
		}

		// Verificar si el estado actual es la solución
		if node.IsGoalState(currentNode.State) {
			return currentNode // Has encontrado la solución
		}

		// Generar sucesores y encolarlos
		successors := node.GenerateSuccessors(currentNode)
		for _, successor := range successors {
			stack.Push(successor)
		}
	}
	return nil // No se encontró solución
}

func printSolution(n *node.Node) {
	if n == nil {
		return
	}
	printSolution(n.Parent)
	ClearScreen()
	n.PrintState()
	fmt.Println()
	time.Sleep(1000 * time.Millisecond) // Espera más tiempo
}

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func isSolvable(puzzle []int) bool {
	inversions := 0
	for i := 0; i < len(puzzle); i++ {
		for j := i + 1; j < len(puzzle); j++ {
			if puzzle[i] != 0 && puzzle[j] != 0 && puzzle[i] > puzzle[j] {
				inversions++
			}
		}
	}
	return inversions%2 == 0
}

func withDepth(s node.State) {
	fmt.Println("Estado inicial:")
	puzzle(s)
	//todos los nodos tienen una profundidad de 10
	initialNode := node.NewNode(s, "", nil, 10)

	if solution(initialNode, s, initialNode.Depth) {
		fmt.Println("busqueda por profundidad concluida exitosamente")
	}
}

func solution(initialNode *node.Node, s node.State, Depth int) bool {
	initialNode.PrintState()
	time.Sleep(500 * time.Millisecond) // Espera más tiempo
	ClearScreen()
	solutionNode := DepthFirstSearch(s, Depth)
	if solutionNode != nil {
		fmt.Println("Solución encontrada:")
		printSolution(solutionNode)
		return true
	} else {
		fmt.Println("No se encontró solución.")
		return false
	}
}

func puzzle(s node.State) {
	// Convertir la matriz 2D en una lista de números
	puzzle := make([]int, 0, len(s)*len(s[0]))
	for _, row := range s {
		puzzle = append(puzzle, row...)
	}

	// Verificar si el puzzle tiene solución
	if !isSolvable(puzzle) {
		fmt.Println("Este puzzle no tiene solución.")

	}
}
func withIterativeDepth(s node.State) {
	fmt.Println("Estado inicial:")
	puzzle(s)
	//todos los nodos tienen una profundidad de 10
	initialNode := node.NewNode(s, "", nil, 10)
	for i := 0; i < 10; i++ {
		if solution(initialNode, s, i) {
			fmt.Println("busqueda por profundidad iterativa concluida exitosamente")
			print("Nivel optimo de busqueda:")
			color.New(color.FgGreen).Printf("%3d", i+1)
			break
		}
	}
}
