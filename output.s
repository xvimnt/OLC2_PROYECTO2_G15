.data
x_i_2: .word 20
str0: .asciz "int ++: (20++) =%d\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of ++ operation on x_i ---
    LDR X10, =x_i_2
    LDR W11, [X10]
    ADD W11, W11, #1
    STR W11, [X10]
    // --- End of ++ operation on x_i ---

    // --- Start of println call ---
    LDR X9, =x_i_2
    LDRSW X10, [X9]
    LDR X0, =str0
    MOV X1, X10
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

