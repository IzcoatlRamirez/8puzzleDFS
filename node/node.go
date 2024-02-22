package node

import ("fmt"
		"github.com/fatih/color"
		"math/rand")

type State [][]int

type Node struct {
    State  State
    Action string
    Parent *Node
    Depth  int
}
func NewNode(state State, action string, parent *Node, depth int) *Node {
    return &Node{
        State:  state,
        Action: action,
        Parent: parent,
        Depth:  depth,
    }
}

func (n *Node) PrintState() {
	rows := len(n.State)
	cols := len(n.State[0])

	printCell := func(value int) {
		if value == 0 {
			fmt.Print("   ")
		} else {
			color.New(color.FgCyan).Printf("%3d", value)
		}
	}

	fmt.Println("+----+----+----+")

	for i := 0; i < rows; i++ {
		fmt.Print("|")
		for j := 0; j < cols; j++ {
			printCell(n.State[i][j])
			fmt.Print(" |")
		}
		fmt.Println("\n+----+----+----+")
	}

	fmt.Println()
}


func RandomState() State {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	rand.Shuffle(len(values), func(i, j int) { values[i], values[j] = values[j], values[i] })

	state := make(State, 3)
	for i := 0; i < 3; i++ {
		state[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			state[i][j] = values[i*3+j]
		}
	}

	return state
}

func GenerateSuccessors(currentNode *Node) []*Node {
	successors := []*Node{}

	// Encuentra la posición de la celda vacía (0)
	emptyRow, emptyCol := FindEmptyCell(currentNode.State)

	// Definir movimientos posibles: arriba, abajo, izquierda, derecha
	moves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, move := range moves {
		newRow, newCol := emptyRow+move[0], emptyCol+move[1]

		// Verificar si el movimiento está dentro de los límites del tablero
		if newRow >= 0 && newRow < len(currentNode.State) && newCol >= 0 && newCol < len(currentNode.State[0]) {
			// Copiar el estado actual para evitar cambios no deseados
			newState := CopyState(currentNode.State)

			// Realizar el intercambio de la celda vacía con la celda adyacente
			newState[emptyRow][emptyCol], newState[newRow][newCol] = newState[newRow][newCol], newState[emptyRow][emptyCol]

			// Crear un nuevo nodo sucesor y agregarlo a la lista
			successor := NewNode(newState, MoveString(move), currentNode, currentNode.Depth+1)
			successors = append(successors, successor)

			// Imprimir el estado del puzzle después de cada movimiento
			// fmt.Printf("Movimiento: %s\n", MoveString(move))
			successor.PrintState()
		}
	}

	return successors
}

// Función para encontrar la posición de la celda vacía en el estado actual
func FindEmptyCell(state State) (int, int) {
	for i, row := range state {
		for j, value := range row {
			if value == 0 {
				return i, j
			}
		}
	}
	return -1, -1 // No debería llegar aquí si el estado es válido
}

// Función para copiar un estado (matriz bidimensional)
func CopyState(state State) State {
	newState := make(State, len(state))
	for i := range state {
		newState[i] = make([]int, len(state[i]))
		copy(newState[i], state[i])
	}
	return newState
}

// Función para convertir un movimiento a una cadena (por ejemplo, "arriba", "abajo", "izquierda", "derecha")
func MoveString(move []int) string {
	switch {
	case move[0] == -1:
		return "arriba"
	case move[0] == 1:
		return "abajo"
	case move[1] == -1:
		return "izquierda"
	case move[1] == 1:
		return "derecha"
	default:
		return "movimiento no válido"
	}
}

func IsGoalState(state State) bool {
    
    goalState := State{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 0},
    }
    for i := 0; i < len(state); i++ {
        for j := 0; j < len(state[i]); j++ {
            if state[i][j] != goalState[i][j] {
                return false
            }
        }
    }

    return true
}
