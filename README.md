Este proyecto comienza con una implementación del **algoritmo de búsqueda en profundidad** (DFS, por sus siglas en inglés) para resolver el rompecabezas del 8-puzzle. El 8-puzzle es un juego que consiste en un tablero de 3x3 con ocho fichas numeradas y un espacio vacío. El objetivo es mover las fichas alrededor del tablero para llegar a un estado objetivo.

El proyecto está organizado en tres paquetes principales:

main: Este es el punto de entrada del programa. Aquí se genera el estado inicial del rompecabezas, se verifica si es solucionable y luego se aplica la búsqueda en profundidad para encontrar una solución. Si se encuentra una solución, se imprime en la consola.

node: Este paquete define la estructura de un nodo en el árbol de búsqueda. Cada nodo contiene el estado actual del rompecabezas, la acción que llevó a ese estado, una referencia al nodo padre y la profundidad del nodo en el árbol de búsqueda. Este paquete también proporciona funciones para generar sucesores de un nodo (es decir, los estados que se pueden alcanzar a partir del estado actual con un solo movimiento) y para verificar si un estado es el estado objetivo.

stack: Este paquete implementa una pila, que es la estructura de datos utilizada por el algoritmo de búsqueda en profundidad para mantener un registro de los nodos que aún no se han explorado. La pila se implementa utilizando una lista doblemente enlazada de la biblioteca estándar de Go.

El algoritmo de búsqueda en profundidad funciona de la siguiente manera:

Empieza con el nodo inicial (el estado inicial del rompecabezas) y lo agrega a la pila.
Mientras la pila no esté vacía, saca el nodo superior de la pila. Si este nodo representa el estado objetivo, entonces ha encontrado una solución y el algoritmo se detiene.
Si el nodo no es el estado objetivo, genera todos los sucesores del nodo (los estados que se pueden alcanzar con un solo movimiento) y los agrega a la pila.
Repite los pasos 2 y 3 hasta que se encuentre una solución o se hayan explorado todos los nodos.
Este enfoque es efectivo para resolver el 8-puzzle, pero puede ser ineficiente si el estado objetivo está a una gran profundidad en el árbol de búsqueda, ya que el algoritmo de búsqueda en profundidad puede explorar muchos nodos antes de encontrar la solución
Puedes ver el video de demostración [aquí](https://drive.google.com/file/d/1-QztEqwCj1JzsB3i8UnELwXUIy44w3kP/view?usp=sharing).