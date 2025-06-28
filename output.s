.data
.align 2
str0: .asciz "123.45"
numeroDecimal1_2: .double 0.0
.align 2
str1: .asciz "123.45\" convertido a float64: %f\n"

.extern atof
.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    LDR X9, =str0
    MOV X0, X9
    BL atof
    FMOV D8, D0
    // Storing initializer for numeroDecimal1
    LDR X9, =numeroDecimal1_2
    STR D8, [X9]
    // --- Start of print call ---
    LDR X9, =numeroDecimal1_2
    LDR D8, [X9]
    LDR X0, =str1
    FMOV D0, D8
    BL printf
    // --- End of print call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

