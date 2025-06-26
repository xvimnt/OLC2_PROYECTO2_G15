.data
temperature_2: .quad 0
.align 2
str0: .asciz "It's hot.\n"
.align 2
str1: .asciz "It's not hot.\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    MOV X9, #25
    // Storing initializer for temperature
    LDR X10, =temperature_2
    STR W9, [X10]
    LDR X9, =temperature_2
    LDRSW X10, [X9]
    MOV X9, #30
    CMP X10, X9
    CSET X10, GT
    // If statement condition check
    CMP W10, #0
    B.EQ .Lelse1
    // 'Then' block
    // --- Start of println call ---
    LDR X0, =str0
    BL printf
    // --- End of println call ---
    B .Lendif2
.Lelse1:
    // 'Else' block
    // --- Start of println call ---
    LDR X0, =str1
    BL printf
    // --- End of println call ---
.Lendif2:
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

