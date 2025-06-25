.data
i1_2: .word 10
i2_2: .word 3
f1_2: .double 12.5
f2_2: .double 2.5
str0: .asciz "\nint / int (truncation): 10 / 3 = %d\n"
str1: .asciz "float64 / float64:      12.5 / 2.5 = %f\n"
str2: .asciz "int / float64:          10 / 2.5 = %f\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of println call ---
    LDR X9, =i1_2
    LDRSW X10, [X9]
    LDR X9, =i2_2
    LDRSW X11, [X9]
    SDIV X10, X10, X11
    LDR X0, =str0
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =f1_2
    LDR D8, [X9]
    LDR X9, =f2_2
    LDR D9, [X9]
    FDIV D8, D8, D9
    LDR X0, =str1
    FMOV D0, D8
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =i1_2
    LDRSW X10, [X9]
    LDR X9, =f2_2
    LDR D8, [X9]
    // Promoting left operand from INT to FLOAT
    SCVTF D9, X10
    FDIV D9, D9, D8
    LDR X0, =str2
    FMOV D0, D9
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

