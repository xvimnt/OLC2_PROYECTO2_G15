.data
str0: .asciz "\n--- B. Compound Assignment Operators ---\n"
x_i_2: .word 20
str1: .asciz "int += int: (20 += 5) =%d\n"
str2: .asciz "int ++: (25++) =%d\n"
str3: .asciz "int -= int: (25 -= 10) =%d\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of println call ---
    LDR X0, =str0
    BL printf
    // --- End of println call ---
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
    LDR X0, =str1
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of ++ operation on x_i ---
    LDR X10, =x_i_2
    LDR W11, [X10]
    ADD W11, W11, #1
    STR W11, [X10]
    // --- End of ++ operation on x_i ---

    // --- Start of println call ---
    LDR X9, =x_i_2
    LDRSW X10, [X9]
    LDR X0, =str2
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of compound assignment (-=) to x_i ---
    LDR X10, =x_i_2
    LDR W11, [X10]
    MOV W12, #10
    SUB W11, W11, W12
    STR W11, [X10]
    // --- End of compound assignment (-=) to x_i ---

    // --- Start of println call ---
    LDR X9, =x_i_2
    LDRSW X10, [X9]
    LDR X0, =str3
    MOV X1, X10
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

