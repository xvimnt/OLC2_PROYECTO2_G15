fn main() {
    mut enteroNulo int;
    mut decimalNulo float64;
    mut textoNulo string;
    mut booleanoNulo bool;

    if (enteroNulo == 0 && decimalNulo == 0.0 && textoNulo == "" && booleanoNulo == false) {
        print("OK Valores por defecto: correcto");
    } else {
        print("X Valores por defecto: incorrecto");
    }
}	