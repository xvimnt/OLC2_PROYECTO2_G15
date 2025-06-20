
fn main() {
	mut puntos int = 0

	println("=== Archivo de prueba de funcionalidades intermedias ===")

	// 1. Manejo de entornos (3 puntos)
	println("==== Manejo de entornos ====")
	mut puntosEntornos int = 0

	println("Variable redeclarada en el mismo entorno")
	mut a int = 10
	// a = "" // ! ERROR: Comentar esta línea para que el programa compile
	println("a =", a)

	if a == 10 {
		puntosEntornos = puntosEntornos + 1
		  println("OK Detección de redeclaración en mismo entorno: correcto")
	} else {
		  println("X Detección de redeclaración en mismo entorno: incorrecto")
	}

	println("\nVariable redeclarada en un entorno diferente")
	mut b int = 10
	if true {
        // verificamos con entornos si la variable es manejada
		b = 20
		  println("b dentro del if =", b)

		if b == 20 {
			puntosEntornos = puntosEntornos + 1
			  println("OK Redeclaración en entorno diferente: correcto")
		} else {
			  println("X Redeclaración en entorno diferente: incorrecto")
		}
	}
	  println("b fuera del if =", b)

	  println("\nUso de variable en un entorno superior")
	mut c int = 10 
	mut d int = 10
	if true {
		// d se redefine en este ámbito
		d = 20
		// modificamos c 
        c = 50 
		
		// Modificamos d del ámbito local
		d = 30
	}
	  println("c fuera del if =", c)
	  println("d fuera del if =", d)

	if c == 30 && d == 10 {
		puntosEntornos = puntosEntornos + 1
		  println("OK Uso de variable en entorno superior: correcto")
	} else {
		  println("X Uso de variable en entorno superior: incorrecto")
	}

	// 2. If / Else (3 puntos)
	println("\n==== If / Else ====")
	mut puntosIfElse int = 0

	println("If simple")
	if true {
		println("Condición verdadera")
		puntosIfElse = puntosIfElse + 1
	}

	println("\nIf-Else")
	if true {
		  println("Condición verdadera en if-else")
	} else {
		  println("Condición falsa en if-else")
	}

	if false {
		  println("Esto no debería imprimirse")
	} else {
		  println("Condición falsa, ejecutando else")
		puntosIfElse = puntosIfElse + 1
	}

	println("\nIf-ElseIf-Else")
	if true {
		  println("Primera condición verdadera")
	} else if true {
		  println("Segunda condición verdadera, pero no se ejecuta")
	} else {
		  println("Ninguna condición verdadera")
	}

	if false {
		  println("Primera condición falsa")
	} else if true {
		  println("Segunda condición verdadera")
		puntosIfElse = puntosIfElse + 1
	} else {
		  println("Ninguna condición verdadera")
	}

	// 3. For Tipo While (2 puntos)
	println("\n==== For Tipo While ====")
	mut puntosForWhile int = 0

	println("For como while simple")
	i = 0
	suma = 0
	for i < 5 {
		  println("i =", i)
		suma = suma + i
		i = i + 1
	}

	if suma == 10 {
		puntosForWhile = puntosForWhile + 1
		  println("OK For como while simple: correcto")
	} else {
		  println("X For como while simple: incorrecto")
	}

	println("\nFor como while anidado (patrón X)")
	println("###Validacion Manual")

	mut n int  = 5
	mut x int = 0
	for x < n {
		j = 0
		fila = ""

		for j < n {
			if x == j || x+j == n-1 {
				fila = fila + "*"
			} else {
				fila = fila + " "
			}
			j = j + 1
		}

		  println(fila)
		x = x + 1
	}

	if x == 5 {
		puntosForWhile = puntosForWhile + 1
		  println("OK For como while anidado: correcto")
	} else {
		  println("X For como while anidado: incorrecto")
	}

	// 4. For Clásico (3 puntos)
	  println("\n==== For Clásico ====")
	puntosForClasico = 0

	println("For clásico simple")
	suma = 0
	for i = 0; i < 5; i++ {
		  println("i =", i)
		suma = suma + i
	}

	if suma == 10 {
		puntosForClasico = puntosForClasico + 1
		  println("OK For clásico simple: correcto")
	} else {
		  println("X For clásico simple: incorrecto")
	}

	  println("\nFor clásico anidado (tabla de multiplicar)")
	  println("###Validacion Manual")
	for i = 1; i <= 3; i++ {
		for j = 1; j <= 3; j++ {
			  println(i, "x", j, "=", i*j)
		}
		  println()
	}
	puntosForClasico = puntosForClasico + 2

	// 5. For Range (3 puntos)
	println("\n==== For Range ====")
	mut puntosForRange int = 0

	println("For range con slice")
	numeros = []int{10, 20, 30, 40, 50}
	suma = 0
	sumaIndices = 0

}