.data
i_2: .quad 0
.align 2
str0: .asciz "%d\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    MOV X9, #0
    // Storing initializer for i
    LDR X10, =i_2
    STR W9, [X10]
loop_start1:
    LDR X9, =i_2
    LDRSW X10, [X9]
    MOV X9, #5
    CMP W10, W9
    CSET W10, LT
    // For loop condition check
    CMP W10, #0
    B.EQ loop_end4
    B loop_body2
loop_post3:
    B loop_start1
loop_body2:
    LDR X9, =i_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to i
    LDR X9, =i_2
    STR W10, [X9]
    LDR X9, =i_2
    LDRSW X10, [X9]
    MOV X9, #3
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif6
    // 'Then' block
    B loop_post3
.Lendif6:
    // --- Start of print call ---
    LDR X9, =i_2
    LDRSW X10, [X9]
    LDR X0, =str0
    MOV X1, X10
    BL printf
    // --- End of print call ---
    B loop_post3
loop_end4:
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

