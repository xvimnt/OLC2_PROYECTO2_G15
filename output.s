.data
str0: .asciz "\n--- A. Arithmetic Operators ---\n"
i1_2: .word 10
i2_2: .word 3
f1_2: .double 12.5
f2_2: .double 2.5
str1: .asciz "int + int:      10 + 3 = %d\n"
str2: .asciz "int + float64:  10 + 2.5 = %f\n"
str3: .asciz "float64 + int:  12.5 + 3 = %f\n"
str4: .asciz "float64 + float64: 12.5 + 2.5 = %f\n"

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
    // --- Start of println call ---
    LDR X9, =i1_2
    LDRSW X10, [X9]
    LDR X9, =i2_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    LDR X0, =str1
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =i1_2
    LDRSW X10, [X9]
    LDR X9, =f2_2
    LDR D8, [X9]
    // Promoting left operand from INT to FLOAT
    SCVTF D9, X10
    FADD D9, D9, D8
    LDR X0, =str2
    FMOV D0, D9
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =f1_2
    LDR D8, [X9]
    LDR X9, =i2_2
    LDRSW X10, [X9]
    // Promoting right operand from INT to FLOAT
    SCVTF D9, X10
    FADD D8, D8, D9
    LDR X0, =str3
    FMOV D0, D8
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =f1_2
    LDR D8, [X9]
    LDR X9, =f2_2
    LDR D9, [X9]
    FADD D8, D8, D9
    LDR X0, =str4
    FMOV D0, D8
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

