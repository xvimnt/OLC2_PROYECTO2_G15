.data
score_1: .quad 0
.align 2
str0: .asciz "Invalid score"
.align 2
str1: .asciz "A"
.align 2
str2: .asciz "B"
.align 2
str3: .asciz "C or below"
.align 2
str4: .asciz "Grade for score 95: %s\n"

.extern printf
.text
.global get_grade
get_grade:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // Store param 'score' from register X0 to memory
    LDR X9, =score_1
    STR X0, [X9]
    LDR X9, =score_1
    LDRSW X10, [X9]
    MOV X9, #0
    CMP X10, X9
    CSET X10, LT
    LDR X9, =score_1
    LDRSW X11, [X9]
    MOV X9, #100
    CMP X11, X9
    CSET X11, GT
    ORR W10, W10, W11
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif2
    // 'Then' block
    LDR X9, =str0
    MOV X0, X9
    B .Lget_grade_epilogue
.Lendif2:
    LDR X9, =score_1
    LDRSW X10, [X9]
    MOV X9, #90
    CMP X10, X9
    CSET X10, GE
    // If statement condition check
    CMP W10, #0
    B.EQ .Lelse3
    // 'Then' block
    LDR X9, =str1
    MOV X0, X9
    B .Lget_grade_epilogue
    B .Lendif4
.Lelse3:
    // 'Else' block
    LDR X9, =score_1
    LDRSW X10, [X9]
    MOV X9, #80
    CMP X10, X9
    CSET X10, GE
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif6
    // 'Then' block
    LDR X9, =str2
    MOV X0, X9
    B .Lget_grade_epilogue
.Lendif6:
.Lendif4:
    LDR X9, =str3
    MOV X0, X9
    B .Lget_grade_epilogue
.Lget_grade_epilogue:
    LDP X29, X30, [SP], #16
    RET

.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of println call ---
    MOV X9, #95
    MOV X0, X9
    BL get_grade
    MOV X9, X0
    LDR X0, =str4
    MOV X1, X9
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

