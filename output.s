.data
logged_in_2: .quad 0
.align 2
str0: .asciz "admin"
user_role_2: .quad 0
.align 2
str1: .asciz "admin"
.align 2
str2: .asciz "moderator"
.align 2
str3: .asciz "Access level: Privileged\n"

.extern strcmp
.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    MOV X9, #1
    // Storing initializer for logged_in
    LDR X10, =logged_in_2
    STR W9, [X10]
    LDR X9, =str0
    // Storing initializer for user_role
    LDR X10, =user_role_2
    STR X9, [X10]
    LDR X10, =logged_in_2
    LDRB W11, [X10]
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false4
    LDR X11, =user_role_2
    LDR X12, [X11]
    LDR X11, =str1
    MOV X0, X12
    MOV X1, X11
    BL strcmp
    CMP W0, #0
    CSET W13, EQ
    // Short-circuit OR: check left operand
    CMP W13, #0
    B.NE .L_logic_true6
    LDR X11, =user_role_2
    LDR X12, [X11]
    LDR X11, =str2
    MOV X0, X12
    MOV X1, X11
    BL strcmp
    CMP W0, #0
    CSET W13, EQ
    // Left was false, result is right operand
    MOV W10, W13
    B .L_logic_end5
.L_logic_true6:
    MOV W10, #1
.L_logic_end5:
    // Left was true, result is right operand
    MOV W9, W10
    B .L_logic_end3
.L_logic_false4:
    MOV W9, #0
.L_logic_end3:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lendif2
    // 'Then' block
    // --- Start of println call ---
    LDR X0, =str3
    BL printf
    // --- End of println call ---
.Lendif2:
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

