fn main() {
    mut puntos int = 0

    println("=== Archivo de prueba de slices ===")

    // 1. Creación de slices (3 puntos)
    println("==== Creación de slices ====")
    mut puntosCreacion int = 0

    println("Creación con literales:")
    numeros := []int {1, 2, 3, 4, 5}
    println("###Validacion Manual")
    println("numeros:", numeros)
    println("OK Creación con literales: correcto")
    puntosCreacion = puntosCreacion + 3

    // 2. Acceso de Elementos (3 puntos)
    println("\n==== Acceso de Elementos ====")
    mut puntosAcceso int = 0

    println("Acceso por índice:")
    mut primerElemento int = numeros[0]
    println("Primer elemento:", primerElemento)

    if primerElemento == 1 {
        puntosAcceso = puntosAcceso + 1
        println("OK Acceso por índice: correcto")
    } else {
        println("X Acceso por índice: incorrecto")
    }

    println("Acceso fuera de rango:")
    // println(numeros[5]) // ! ERROR: Comentar esta línea para que el programa compile
    println("OK Acceso fuera de rango: correcto")

    println("\nModificación de elementos:")
    numeros[0] = 10
    println("numeros después de modificar:", numeros)

    if numeros[0] == 10 {
        puntosAcceso = puntosAcceso + 2
        println("OK Modificación de elementos: correcto")
    } else {
        println("X Modificación de elementos: incorrecto")
    }

    // 3. Array Multidimensional (4 puntos)
    println("\n==== Array Multidimensional ====")
    mut puntosMultidimensional int = 0

    println("Creación de matriz 2D con literales:")
    matriz := [][]int {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    println("matriz:")
    println("###Validacion Manual")
    println(matriz)
    puntosMultidimensional = puntosMultidimensional + 2

    println("\nMatrices irregulares:")
    matrizIrregular := [][]int {
        {1, 2, 3, 4},
        {5, 6},
        {7, 8, 9},
    }
    println("matrizIrregular:")
    println("###Validacion Manual")
    println(matrizIrregular)
    puntosMultidimensional = puntosMultidimensional + 2

    // 4. Acceso Array Multidimensional (4 puntos)
    println("\n==== Acceso Array Multidimensional ====")
    mut puntosAccesoMultidimensional int = 0

    println("Acceso a elementos de matriz 2D:")
    mut elemento11 int = matriz[0][0]
    mut elemento23 int = matriz[1][2]
    mut elemento33 int = matriz[2][2]
    println("Elemento [0][0]:", elemento11)
    println("Elemento [1][2]:", elemento23)
    println("Elemento [2][2]:", elemento33)

    if elemento11 == 1 && elemento23 == 6 && elemento33 == 9 {
        puntosAccesoMultidimensional = puntosAccesoMultidimensional + 1
        println("OK Acceso a elementos de matriz 2D: correcto")
    } else {
        println("X Acceso a elementos de matriz 2D: incorrecto")
    }

    println("\nModificación de elementos en matriz 2D:")
    matriz[0][0] = 100
    matriz[1][1] = 500
    matriz[2][2] = 900

    if matriz[0][0] == 100 && matriz[1][1] == 500 && matriz[2][2] == 900 {
        puntosAccesoMultidimensional = puntosAccesoMultidimensional + 1
        println("OK Modificación de elementos en matriz 2D: correcto")
    } else {
        println("X Modificación de elementos en matriz 2D: incorrecto")
    }

    println("\nAcceso a filas completas:")
    mut primeraFila []int = matriz[0]
    println("Primera fila de matriz:", primeraFila)

    if primeraFila[0] == 100 && primeraFila[1] == 2 && primeraFila[2] == 3 {
        puntosAccesoMultidimensional = puntosAccesoMultidimensional + 2
        println("OK Acceso a filas completas: correcto")
    } else {
        println("X Acceso a filas completas: incorrecto")
    }

    // 5. Función indexOf (1 punto)
    println("\n==== Función indexOf ====")
    mut puntosIndex int = 0

    println("Búsqueda de elementos con indexOf:")
    numeros = []int {10, 20, 30, 40, 50}
    mut indice1 int = indexOf(numeros, 30)
    mut indice2 int = indexOf(numeros, 60) // No existe, debería retornar -1
    println("Índice de 30:", indice1)
    println("Índice de 60:", indice2)

    if indice1 == 2 && indice2 == -1 {
        puntosIndex = puntosIndex + 1
        println("OK indexOf: correcto")
    } else {
        println("X indexOf: incorrecto")
    }

    // 6. Función join (1 punto)
    println("\n==== Función join ====")
    mut puntosJoin int = 0

    println("Unión de strings con join:")
    palabras := []string {"Hola", "mundo", "desde", "Go"}
    mut frase string = join(palabras, " ")
    mut fraseConComas string = join(palabras, ", ")
    println("Frase con espacios:", frase)
    println("Frase con comas:", fraseConComas)

    if frase == "Hola mundo desde Go" && fraseConComas == "Hola, mundo, desde, Go" {
        puntosJoin = puntosJoin + 1
        println("OK join: correcto")
    } else {
        println("X join: incorrecto")
    }

    // 7. Función len (1 punto)
    println("\n==== Función len ====")
    mut puntosLen int = 0

    println("Longitud de slices con len:")
    mut longitud1 int = len(numeros)
    mut longitud2 int = len(matrizIrregular)
    mut longitud3 int = len(matrizIrregular[1])

    if longitud1 == 5 && longitud2 == 3 && longitud3 == 2 {
        puntosLen = puntosLen + 1
        println("OK len: correcto")
    } else {
        println("X len: incorrecto")
    }

    // 8. Función append (3 puntos)
    println("\n==== Función append ====")
    mut puntosAppend int = 0

    println("Agregar elementos con append:")
    numeros = []int {1, 2, 3}
    numeros = append(numeros, 4)
    println("numeros después de append(numeros, 4):", numeros)

    if numeros[3] == 4 {
        puntosAppend = puntosAppend + 1
        println("OK Agregar un elemento: correcto")
    } else {
        println("X Agregar un elemento: incorrecto")
    }

    println("\nAgregar un slice a otro con append:")
    mtx1 := [][]int {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    mtx1 = append(mtx1, numeros)
    println("mtx1 después de append(mtx1, numeros):")
    println("###Validacion Manual")
    println(mtx1)
    puntosAppend = puntosAppend + 2

    // Resumen de puntos
    puntos = puntosCreacion + puntosAcceso + puntosMultidimensional +
        puntosAccesoMultidimensional + puntosIndex + puntosJoin +
        puntosLen + puntosAppend

    println("\n=== Errores ===")
    println("###Validacion Manual")
    println("Errores esperados ?/1")

    println("\n=== Tabla de Resultados ===")
    println("+----------------------------------+--------+-------+")
    println("| Característica                   | Puntos | Total |")
    println("+----------------------------------+--------+-------+")
    println("| Creación de slices               |", puntosCreacion, "   | 3     |")
    println("| Acceso de Elementos              |", puntosAcceso, "   | 3     |")
    println("| Array Multidimensional           |", puntosMultidimensional, "   | 4     |")
    println("| Acceso Array Multidimensional    |", puntosAccesoMultidimensional, "   | 4     |")
    println("| Función indexOf                  |", puntosIndex, "   | 1     |")
    println("| Función join                     |", puntosJoin, "   | 1     |")
    println("| Función len                      |", puntosLen, "   | 1     |")
    println("| Función append                   |", puntosAppend, "   | 3     |")
    println("+----------------------------------+--------+-------+")
    println("| TOTAL                            |", puntos, "  | 20    |")
    println("+----------------------------------+--------+-------+")
}