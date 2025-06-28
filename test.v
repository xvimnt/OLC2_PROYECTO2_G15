fn main() {
    mut puntosTypeOf int = 0;
    // Tipo int
    mut entero int = 42;
    mut tipoEntero string = typeof(entero);
    print("Tipo de 42:", tipoEntero);

    // Tipo float64
    mut decimal float64 = 3.14159;
    mut tipoDecimal string = typeof(decimal);
    print("Tipo de 3.14159:", tipoDecimal);

    // Tipo string
    mut texto string = "Hola, mundo!";
    mut tipoTexto string = typeof(texto);
    print("Tipo de \"Hola, mundo!\":", tipoTexto);

    // Tipo bool
    mut booleano bool = true;
    mut tipoBooleano string = typeof(booleano);
    print("Tipo de true:", tipoBooleano);

    // Tipo slice
    mut slice []int = []int{1, 2, 3};
    mut tipoSlice string = typeof(slice);
    print("Tipo de []int{1, 2, 3}:", tipoSlice);

    if (tipoEntero == "int" && tipoDecimal == "float64" &&
        tipoTexto == "string" && tipoBooleano == "bool" &&
        tipoSlice == "[]int") {
        puntosTypeOf = puntosTypeOf + 1;
        print("OK typeof: correcto");
    } else {
        print("X typeof: incorrecto");
    }
}