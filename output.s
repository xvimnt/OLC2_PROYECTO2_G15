.data
.align 2
str0: .asciz "¡Hola, mundo!\n"
a_5: .quad 0
b_5: .quad 0
nombre_7: .quad 0
.align 2
str1: .asciz "¡Hola, %s!\n"
s_9: .quad 0
s_11: .quad 0
.align 2
str2: .asciz "123.45"
.align 3
F3: .double 123.45
.align 2
str4: .asciz "123"
.align 3
F5: .double 123.0
.align 3
F6: .double 0.0

.extern strcmp
.extern printf
.text
.global saludar
saludar:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of println call ---
    LDR X0, =str0
    BL printf
    // --- End of println call ---
.Lsaludar_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

.global obtener_numero
obtener_numero:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    MOV X9, #42
    MOV X0, X9
    B .Lobtener_numero_epilogue
.Lobtener_numero_epilogue:
    LDP X29, X30, [SP], #16
    RET

.global sumar
sumar:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // Store param 'a' from register X0 to memory
    LDR X9, =a_5
    STR X0, [X9]
    // Store param 'b' from register X1 to memory
    LDR X9, =b_5
    STR X1, [X9]
    LDR X9, =a_5
    LDRSW X10, [X9]
    LDR X9, =b_5
    LDRSW X11, [X9]
    ADD X10, X10, X11
    MOV X0, X10
    B .Lsumar_epilogue
.Lsumar_epilogue:
    LDP X29, X30, [SP], #16
    RET

.global saludar_persona
saludar_persona:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // Store param 'nombre' from register X0 to memory
    LDR X9, =nombre_7
    STR X0, [X9]
    // --- Start of println call ---
    LDR X9, =nombre_7
    LDR X10, [X9]
    LDR X0, =str1
    MOV X1, X10
    BL printf
    // --- End of println call ---
.Lsaludar_persona_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

.global atoi
atoi:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // Store param 's' from register X0 to memory
    LDR X9, =s_9
    STR X0, [X9]
    MOV X9, #123
    MOV X0, X9
    B .Latoi_epilogue
.Latoi_epilogue:
    LDP X29, X30, [SP], #16
    RET

.global parse_float
parse_float:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // Store param 's' from register X0 to memory
    LDR X9, =s_11
    STR X0, [X9]
    LDR X9, =s_11
    LDR X10, [X9]
    LDR X9, =str2
    MOV X0, X10
    MOV X1, X9
    BL strcmp
    CMP W0, #0
    CSET W11, EQ
    // If statement condition check
    CMP W11, #0
    B.EQ .Lendif2
    // 'Then' block
    LDR D8, F3
    FMOV D0, D8
    B .Lparse_float_epilogue
.Lendif2:
    LDR X9, =s_11
    LDR X10, [X9]
    LDR X9, =str4
    MOV X0, X10
    MOV X1, X9
    BL strcmp
    CMP W0, #0
    CSET W11, EQ
    // If statement condition check
    CMP W11, #0
    B.EQ .Lendif4
    // 'Then' block
    LDR D8, F5
    FMOV D0, D8
    B .Lparse_float_epilogue
.Lendif4:
    LDR D8, F6
    FMOV D0, D8
    B .Lparse_float_epilogue
.Lparse_float_epilogue:
    LDP X29, X30, [SP], #16
    RET

.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

