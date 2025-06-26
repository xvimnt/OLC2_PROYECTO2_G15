.data
puntos_2: .quad 0
.align 2
str0: .asciz "=== Prueba Básica Simplificada ===\n"
entero_2: .quad 0
.align 3
F1: .double 3.14
decimal_2: .double 0.0
.align 2
str2: .asciz "Hola, mundo!"
texto_2: .quad 0
booleano_2: .quad 0
.align 2
str3: .asciz "OK entero\n"
.align 3
F4: .double 3.0
.align 2
str5: .asciz "OK decimal\n"
.align 2
str6: .asciz "Hola, mundo!"
.align 2
str7: .asciz "OK texto\n"
.align 2
str8: .asciz "OK booleano\n"
.align 2
str9: .asciz "OK asignación entero\n"
suma_2: .quad 0
.align 2
str10: .asciz "OK suma\n"
resta_2: .quad 0
.align 2
str11: .asciz "OK resta\n"
.align 2
str12: .asciz "OK igualdad\n"
.align 2
str13: .asciz "OK lógica AND\n"
.align 2
str14: .asciz "OK lógica OR negada\n"
.align 2
str15: .asciz "%d\n"
.align 2
str16: .asciz "Texto de prueba\n"
.align 2
str17: .asciz "%s\n"
.align 2
str18: .asciz "true"
.align 2
str19: .asciz "false"
.align 2
str20: .asciz "Puntos obtenidos: %d / 11\n"

.extern strcmp
.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    MOV X9, #0
    // Storing initializer for puntos
    LDR X10, =puntos_2
    STR W9, [X10]
    // --- Start of println call ---
    LDR X0, =str0
    BL printf
    // --- End of println call ---
    MOV X9, #42
    // Storing initializer for entero
    LDR X10, =entero_2
    STR W9, [X10]
    LDR D8, F1
    // Storing initializer for decimal
    LDR X9, =decimal_2
    STR D8, [X9]
    LDR X9, =str2
    // Storing initializer for texto
    LDR X10, =texto_2
    STR X9, [X10]
    MOV X9, #1
    // Storing initializer for booleano
    LDR X10, =booleano_2
    STR W9, [X10]
    LDR X9, =entero_2
    LDRSW X10, [X9]
    MOV X9, #42
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif2
    // 'Then' block
    // --- Start of println call ---
    LDR X0, =str3
    BL printf
    // --- End of println call ---
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
.Lendif2:
    LDR X9, =decimal_2
    LDR D8, [X9]
    LDR D9, F4
    FCMP D8, D9
    CSET W9, GT
    // If statement condition check
    CMP W9, #0
    B.EQ .Lendif4
    // 'Then' block
    // --- Start of println call ---
    LDR X0, =str5
    BL printf
    // --- End of println call ---
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
.Lendif4:
    LDR X9, =texto_2
    LDR X10, [X9]
    LDR X9, =str6
    MOV X0, X10
    MOV X1, X9
    BL strcmp
    CMP W0, #0
    CSET W11, EQ
    // If statement condition check
    CMP W11, #0
    B.EQ .Lendif6
    // 'Then' block
    // --- Start of println call ---
    LDR X0, =str7
    BL printf
    // --- End of println call ---
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
.Lendif6:
    LDR X9, =booleano_2
    LDRB W10, [X9]
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif8
    // 'Then' block
    // --- Start of println call ---
    LDR X0, =str8
    BL printf
    // --- End of println call ---
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
.Lendif8:
    MOV X9, #99
    // Storing value for assignment to entero
    LDR X9, =entero_2
    STR W9, [X9]
    LDR X9, =entero_2
    LDRSW X10, [X9]
    MOV X9, #99
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif10
    // 'Then' block
    // --- Start of println call ---
    LDR X0, =str9
    BL printf
    // --- End of println call ---
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
.Lendif10:
    MOV X9, #10
    MOV X10, #5
    ADD X9, X9, X10
    // Storing initializer for suma
    LDR X10, =suma_2
    STR W9, [X10]
    LDR X9, =suma_2
    LDRSW X10, [X9]
    MOV X9, #15
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif12
    // 'Then' block
    // --- Start of println call ---
    LDR X0, =str10
    BL printf
    // --- End of println call ---
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
.Lendif12:
    MOV X9, #10
    MOV X10, #5
    SUB X9, X9, X10
    // Storing initializer for resta
    LDR X10, =resta_2
    STR W9, [X10]
    LDR X9, =resta_2
    LDRSW X10, [X9]
    MOV X9, #5
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif14
    // 'Then' block
    // --- Start of println call ---
    LDR X0, =str11
    BL printf
    // --- End of println call ---
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
.Lendif14:
    MOV X9, #10
    MOV X10, #10
    CMP W9, W10
    CSET W9, EQ
    // If statement condition check
    CMP W9, #0
    B.EQ .Lendif16
    // 'Then' block
    // --- Start of println call ---
    LDR X0, =str12
    BL printf
    // --- End of println call ---
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
.Lendif16:
    MOV X10, #1
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false20
    MOV X10, #1
    // Left was true, result is right operand
    MOV W9, W10
    B .L_logic_end19
.L_logic_false20:
    MOV W9, #0
.L_logic_end19:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lendif18
    // 'Then' block
    // --- Start of println call ---
    LDR X0, =str13
    BL printf
    // --- End of println call ---
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
.Lendif18:
    MOV X10, #0
    // Short-circuit OR: check left operand
    CMP W10, #0
    B.NE .L_logic_true24
    MOV X10, #0
    // Left was false, result is right operand
    MOV W9, W10
    B .L_logic_end23
.L_logic_true24:
    MOV W9, #1
.L_logic_end23:
    EOR W9, W9, #1
    // If statement condition check
    CMP W9, #0
    B.EQ .Lendif22
    // 'Then' block
    // --- Start of println call ---
    LDR X0, =str14
    BL printf
    // --- End of println call ---
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
.Lendif22:
    // --- Start of println call ---
    MOV X9, #42
    LDR X0, =str15
    MOV X1, X9
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X0, =str16
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    MOV X9, #1
    LDR X0, =str17
    LDR X10, =str18
    LDR X11, =str19
    CMP X9, #0
    CSEL X1, X10, X11, NE
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =puntos_2
    LDRSW X10, [X9]
    LDR X0, =str20
    MOV X1, X10
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

