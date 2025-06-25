.data
prec_result_2: .byte 0
.align 2
str0: .asciz "5 * 2 + 3 > 12 && !false is %s\n"
.align 2
str1: .asciz "true"
.align 2
str2: .asciz "false"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    MOV X9, #5
    MOV X10, #2
    MUL X9, X9, X10
    MOV X10, #3
    ADD X9, X9, X10
    MOV X10, #12
    CMP X9, X10
    CSET X9, GT
    MOV X10, #0
    EOR W10, W10, #1
    AND W9, W9, W10
    LDR X10, =prec_result_2
    STRB W9, [X10]
    // --- Start of println call ---
    LDR X9, =prec_result_2
    LDRB W10, [X9]
    LDR X0, =str0
    LDR X9, =str1
    LDR X11, =str2
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

