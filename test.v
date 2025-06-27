fn main() {
     print("\n==== Switch/Case ====");
    mut puntosSwitch int = 0;

    print("\n\n###Validacion Manual");
    print("Switch simple");
    mut dia int = 1;

    switch dia {
    case 1:
        print("Lunes");
        puntosSwitch = puntosSwitch + 1;
    case 2:
        print("Martes");
    case 3:
        print("Miércoles");
    case 4:
        print("Jueves");
    case 5:
        print("Viernes");
    case 6:
        print("Sábado");
    case 7:
        print("Domingo");
    default:
        print("Día inválido");
    }

    print("\nSwitch con default");
    mut numero int = 100;

    switch numero {
    case 1:
        print("No se debería imprimir");
    case 2:
        print("No se debería imprimir");
    default:
        print("Número no reconocido, se ejecuta default");
        puntosSwitch = puntosSwitch + 1;
    }

    print("\nSwitch con break explícito");

    mut numeroBreak int = 2;

    switch numeroBreak {
    case 1:
        print("No se debería imprimir");
    case 2:
        print("Caso 2 - Se ejecuta este y debe detenerse");
        puntosSwitch = puntosSwitch + 1;
        break;
        print("No debería ejecutarse si el break funciona");
        puntosSwitch = puntosSwitch - 1;
    case 3:
        print("No se debería imprimir");
    }
    print("");
}