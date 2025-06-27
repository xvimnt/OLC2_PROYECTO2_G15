.data
enteroNulo_2: .quad 0
.align 3
decimalNulo_2: .double 0.0
.align 2
str0: .asciz ""
textoNulo_2: .quad str0
booleanoNulo_2: .quad 0
.align 3
F1: .double 0.0
.align 2
str2: .asciz ""
.align 2
str3: .asciz "OK Valores por defecto: correcto\n"
.align 2
str4: .asciz "X Valores por defecto: incorrecto\n"

.extern strcmp
.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // Declaring enteroNulo without initializer
    // Declaring decimalNulo without initializer
    // Declaring textoNulo without initializer
    // Declaring booleanoNulo without initializer
    LDR X12, =enteroNulo_2
    LDRSW X13, [X12]
    MOV X12, #0
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false8
    LDR X12, =decimalNulo_2
    LDR D8, [X12]
    LDR D9, F1
    FCMP D8, D9
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W11, W12
    B .L_logic_end7
.L_logic_false8:
    MOV W11, #0
.L_logic_end7:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false6
    LDR X11, =textoNulo_2
    LDR X12, [X11]
    LDR X11, =str2
    MOV X0, X12
    MOV X1, X11
    BL strcmp
    CMP W0, #0
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W10, W13
    B .L_logic_end5
.L_logic_false6:
    MOV W10, #0
.L_logic_end5:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false4
    LDR X10, =booleanoNulo_2
    LDRB W11, [X10]
    MOV X10, #0
    CMP W11, W10
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W9, W11
    B .L_logic_end3
.L_logic_false4:
    MOV W9, #0
.L_logic_end3:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse1
    // 'Then' block
    // --- Start of print call ---
    LDR X0, =str3
    BL printf
    // --- End of print call ---
    B .Lendif2
.Lelse1:
    // 'Else' block
    // --- Start of print call ---
    LDR X0, =str4
    BL printf
    // --- End of print call ---
.Lendif2:
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

