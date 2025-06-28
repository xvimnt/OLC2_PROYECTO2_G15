.data
puntosTypeOf_2: .quad 0
entero_2: .quad 0
.align 2
str0: .asciz "int"
tipoEntero_2: .quad 0
.align 2
str1: .asciz "Tipo de 42: %s\n"
.align 3
F2: .double 3.14159
decimal_2: .double 0.0
.align 2
str3: .asciz "float64"
tipoDecimal_2: .quad 0
.align 2
str4: .asciz "Tipo de 3.14159: %s\n"
.align 2
str5: .asciz "Hola, mundo!"
texto_2: .quad 0
.align 2
str6: .asciz "string"
tipoTexto_2: .quad 0
.align 2
str7: .asciz "Tipo de \"Hola, mundo!\": %s\n"
booleano_2: .quad 0
.align 2
str8: .asciz "bool"
tipoBooleano_2: .quad 0
.align 2
str9: .asciz "Tipo de true: %s\n"
slice_data1: .quad 1, 2, 3
slice_descriptor2:
    .quad slice_data1  // Pointer to data
    .quad 3    // Length
    .quad 3    // Capacity
slice_2: .quad 0
.align 2
str10: .asciz "[]int"
tipoSlice_2: .quad 0
.align 2
str11: .asciz "Tipo de []int{1, 2, 3}: %s\n"
.align 2
str12: .asciz "int"
.align 2
str13: .asciz "float64"
.align 2
str14: .asciz "string"
.align 2
str15: .asciz "bool"
.align 2
str16: .asciz "[]int"
.align 2
str17: .asciz "OK typeof: correcto\n"
.align 2
str18: .asciz "X typeof: incorrecto\n"

.extern strcmp
.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    MOV X9, #0
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =puntosTypeOf_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    MOV X9, #42
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =entero_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    LDR X9, =entero_2
    LDRSW X10, [X9]
    LDR X9, =str0
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =tipoEntero_2
    LDR X10, [SP]
    STR X10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =tipoEntero_2
    LDR X10, [X9]
    LDR X0, =str1
    MOV X1, X10
    BL printf
    // --- End of print call ---
    LDR D8, F2
    SUB SP, SP, #16
    STR D8, [SP]
    LDR X9, =decimal_2
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    LDR X9, =decimal_2
    LDR D8, [X9]
    LDR X9, =str3
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =tipoDecimal_2
    LDR X10, [SP]
    STR X10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =tipoDecimal_2
    LDR X10, [X9]
    LDR X0, =str4
    MOV X1, X10
    BL printf
    // --- End of print call ---
    LDR X9, =str5
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =texto_2
    LDR X10, [SP]
    STR X10, [X9]
    ADD SP, SP, #16
    LDR X9, =texto_2
    LDR X10, [X9]
    LDR X9, =str6
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =tipoTexto_2
    LDR X10, [SP]
    STR X10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =tipoTexto_2
    LDR X10, [X9]
    LDR X0, =str7
    MOV X1, X10
    BL printf
    // --- End of print call ---
    MOV X9, #1
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =booleano_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    LDR X9, =booleano_2
    LDRB W10, [X9]
    LDR X9, =str8
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =tipoBooleano_2
    LDR X10, [SP]
    STR X10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =tipoBooleano_2
    LDR X10, [X9]
    LDR X0, =str9
    MOV X1, X10
    BL printf
    // --- End of print call ---
    LDR X9, =slice_descriptor2
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =slice_2
    LDR X10, [SP]
    STR X10, [X9]
    ADD SP, SP, #16
    LDR X9, =slice_2
    LDR X10, [X9]
    LDR X9, =str10
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =tipoSlice_2
    LDR X10, [SP]
    STR X10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =tipoSlice_2
    LDR X10, [X9]
    LDR X0, =str11
    MOV X1, X10
    BL printf
    // --- End of print call ---
    LDR X13, =tipoEntero_2
    LDR X14, [X13]
    LDR X13, =str12
    MOV X0, X14
    MOV X1, X13
    BL strcmp
    CMP W0, #0
    CSET W15, EQ
    // Short-circuit AND: check left operand
    CMP W15, #0
    B.EQ .L_logic_false12
    LDR X13, =tipoDecimal_2
    LDR X14, [X13]
    LDR X13, =str13
    MOV X0, X14
    MOV X1, X13
    BL strcmp
    CMP W0, #0
    CSET W15, EQ
    // Left was true, result is right operand
    MOV W12, W15
    B .L_logic_end11
.L_logic_false12:
    MOV W12, #0
.L_logic_end11:
    // Short-circuit AND: check left operand
    CMP W12, #0
    B.EQ .L_logic_false10
    LDR X12, =tipoTexto_2
    LDR X13, [X12]
    LDR X12, =str14
    MOV X0, X13
    MOV X1, X12
    BL strcmp
    CMP W0, #0
    CSET W14, EQ
    // Left was true, result is right operand
    MOV W11, W14
    B .L_logic_end9
.L_logic_false10:
    MOV W11, #0
.L_logic_end9:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false8
    LDR X11, =tipoBooleano_2
    LDR X12, [X11]
    LDR X11, =str15
    MOV X0, X12
    MOV X1, X11
    BL strcmp
    CMP W0, #0
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W10, W13
    B .L_logic_end7
.L_logic_false8:
    MOV W10, #0
.L_logic_end7:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false6
    LDR X10, =tipoSlice_2
    LDR X11, [X10]
    LDR X10, =str16
    MOV X0, X11
    MOV X1, X10
    BL strcmp
    CMP W0, #0
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W9, W12
    B .L_logic_end5
.L_logic_false6:
    MOV W9, #0
.L_logic_end5:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse3
    // 'Then' block
    LDR X9, =puntosTypeOf_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosTypeOf
    LDR X9, =puntosTypeOf_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X0, =str17
    BL printf
    // --- End of print call ---
    B .Lendif4
.Lelse3:
    // 'Else' block
    // --- Start of print call ---
    LDR X0, =str18
    BL printf
    // --- End of print call ---
.Lendif4:
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

