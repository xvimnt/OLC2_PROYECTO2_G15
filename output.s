.data
.align 2
str0: .asciz "\n==== Switch/Case ====\n"
puntosSwitch_2: .quad 0
.align 2
str1: .asciz "\n\n###Validacion Manual\n"
.align 2
str2: .asciz "Switch simple\n"
dia_2: .quad 0
.align 2
str3: .asciz "Lunes\n"
.align 2
str4: .asciz "Martes\n"
.align 2
str5: .asciz "Miércoles\n"
.align 2
str6: .asciz "Jueves\n"
.align 2
str7: .asciz "Viernes\n"
.align 2
str8: .asciz "Sábado\n"
.align 2
str9: .asciz "Domingo\n"
.align 2
str10: .asciz "Día inválido\n"
.align 2
str11: .asciz "\nSwitch con default\n"
numero_2: .quad 0
.align 2
str12: .asciz "No se debería imprimir\n"
.align 2
str13: .asciz "No se debería imprimir\n"
.align 2
str14: .asciz "Número no reconocido, se ejecuta default\n"
.align 2
str15: .asciz "\nSwitch con break explícito\n"
numeroBreak_2: .quad 0
.align 2
str16: .asciz "No se debería imprimir\n"
.align 2
str17: .asciz "Caso 2 - Se ejecuta este y debe detenerse\n"
.align 2
str18: .asciz "No debería ejecutarse si el break funciona\n"
.align 2
str19: .asciz "No se debería imprimir\n"
.align 2
str20: .asciz "\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of print call ---
    LDR X0, =str0
    BL printf
    // --- End of print call ---
    MOV X9, #0
    // Storing initializer for puntosSwitch
    LDR X10, =puntosSwitch_2
    STR W9, [X10]
    // --- Start of print call ---
    LDR X0, =str1
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X0, =str2
    BL printf
    // --- End of print call ---
    MOV X9, #1
    // Storing initializer for dia
    LDR X10, =dia_2
    STR W9, [X10]
    LDR X9, =dia_2
    LDRSW X10, [X9]
    // --- Switch Statement ---
    MOV X9, #1
    // Comparing with case: IntegerLiteral: 1
    CMP W10, W9
    BEQ switch_case_03
    MOV X9, #2
    // Comparing with case: IntegerLiteral: 2
    CMP W10, W9
    BEQ switch_case_14
    MOV X9, #3
    // Comparing with case: IntegerLiteral: 3
    CMP W10, W9
    BEQ switch_case_25
    MOV X9, #4
    // Comparing with case: IntegerLiteral: 4
    CMP W10, W9
    BEQ switch_case_36
    MOV X9, #5
    // Comparing with case: IntegerLiteral: 5
    CMP W10, W9
    BEQ switch_case_47
    MOV X9, #6
    // Comparing with case: IntegerLiteral: 6
    CMP W10, W9
    BEQ switch_case_58
    MOV X9, #7
    // Comparing with case: IntegerLiteral: 7
    CMP W10, W9
    BEQ switch_case_69
    B switch_default2
    // --- Switch Case Bodies ---
switch_case_03:
    // --- Start of print call ---
    LDR X0, =str3
    BL printf
    // --- End of print call ---
    LDR X9, =puntosSwitch_2
    LDRSW X11, [X9]
    MOV X9, #1
    ADD X11, X11, X9
    // Storing value for assignment to puntosSwitch
    LDR X9, =puntosSwitch_2
    STR W11, [X9]
    B switch_end1
switch_case_14:
    // --- Start of print call ---
    LDR X0, =str4
    BL printf
    // --- End of print call ---
    B switch_end1
switch_case_25:
    // --- Start of print call ---
    LDR X0, =str5
    BL printf
    // --- End of print call ---
    B switch_end1
switch_case_36:
    // --- Start of print call ---
    LDR X0, =str6
    BL printf
    // --- End of print call ---
    B switch_end1
switch_case_47:
    // --- Start of print call ---
    LDR X0, =str7
    BL printf
    // --- End of print call ---
    B switch_end1
switch_case_58:
    // --- Start of print call ---
    LDR X0, =str8
    BL printf
    // --- End of print call ---
    B switch_end1
switch_case_69:
    // --- Start of print call ---
    LDR X0, =str9
    BL printf
    // --- End of print call ---
    B switch_end1
switch_default2:
    // --- Start of print call ---
    LDR X0, =str10
    BL printf
    // --- End of print call ---
switch_end1:
    // --- Start of print call ---
    LDR X0, =str11
    BL printf
    // --- End of print call ---
    MOV X9, #100
    // Storing initializer for numero
    LDR X10, =numero_2
    STR W9, [X10]
    LDR X9, =numero_2
    LDRSW X10, [X9]
    // --- Switch Statement ---
    MOV X9, #1
    // Comparing with case: IntegerLiteral: 1
    CMP W10, W9
    BEQ switch_case_012
    MOV X9, #2
    // Comparing with case: IntegerLiteral: 2
    CMP W10, W9
    BEQ switch_case_113
    B switch_default11
    // --- Switch Case Bodies ---
switch_case_012:
    // --- Start of print call ---
    LDR X0, =str12
    BL printf
    // --- End of print call ---
    B switch_end10
switch_case_113:
    // --- Start of print call ---
    LDR X0, =str13
    BL printf
    // --- End of print call ---
    B switch_end10
switch_default11:
    // --- Start of print call ---
    LDR X0, =str14
    BL printf
    // --- End of print call ---
    LDR X9, =puntosSwitch_2
    LDRSW X11, [X9]
    MOV X9, #1
    ADD X11, X11, X9
    // Storing value for assignment to puntosSwitch
    LDR X9, =puntosSwitch_2
    STR W11, [X9]
switch_end10:
    // --- Start of print call ---
    LDR X0, =str15
    BL printf
    // --- End of print call ---
    MOV X9, #2
    // Storing initializer for numeroBreak
    LDR X10, =numeroBreak_2
    STR W9, [X10]
    LDR X9, =numeroBreak_2
    LDRSW X10, [X9]
    // --- Switch Statement ---
    MOV X9, #1
    // Comparing with case: IntegerLiteral: 1
    CMP W10, W9
    BEQ switch_case_015
    MOV X9, #2
    // Comparing with case: IntegerLiteral: 2
    CMP W10, W9
    BEQ switch_case_116
    MOV X9, #3
    // Comparing with case: IntegerLiteral: 3
    CMP W10, W9
    BEQ switch_case_217
    B switch_end14
    // --- Switch Case Bodies ---
switch_case_015:
    // --- Start of print call ---
    LDR X0, =str16
    BL printf
    // --- End of print call ---
    B switch_end14
switch_case_116:
    // --- Start of print call ---
    LDR X0, =str17
    BL printf
    // --- End of print call ---
    LDR X9, =puntosSwitch_2
    LDRSW X11, [X9]
    MOV X9, #1
    ADD X11, X11, X9
    // Storing value for assignment to puntosSwitch
    LDR X9, =puntosSwitch_2
    STR W11, [X9]
    B switch_end14
    // --- Start of print call ---
    LDR X0, =str18
    BL printf
    // --- End of print call ---
    LDR X9, =puntosSwitch_2
    LDRSW X11, [X9]
    MOV X9, #1
    SUB X11, X11, X9
    // Storing value for assignment to puntosSwitch
    LDR X9, =puntosSwitch_2
    STR W11, [X9]
    B switch_end14
switch_case_217:
    // --- Start of print call ---
    LDR X0, =str19
    BL printf
    // --- End of print call ---
    B switch_end14
switch_end14:
switch_end14:
    // --- Start of print call ---
    LDR X0, =str20
    BL printf
    // --- End of print call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

