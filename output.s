.data
i1_2: .word 10
i2_2: .word 3
f1_2: .double 12.5
f2_2: .double 2.5
str0: .asciz "\nUnary negation:   -10 = %d\n"
str1: .asciz "Unary negation: -12.5 = %f\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of println call ---
    LDR X9, =i1_2
    LDRSW X10, [X9]
    NEG X10, X10
    LDR X0, =str0
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =f1_2
    LDR D8, [X9]
    FNEG D8, D8
    LDR X0, =str1
    FMOV D0, D8
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

