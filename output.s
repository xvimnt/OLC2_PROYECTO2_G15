.data
x_i_2: .word 20
str0: .asciz "int += int: (20 += 5) =%d\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of compound assignment (+=) to x_i ---
    LDR X10, =x_i_2
    LDR W11, [X10]
    MOV W12, #5
    ADD W11, W11, W12
    STR W11, [X10]
    // --- End of compound assignment (+=) to x_i ---

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

