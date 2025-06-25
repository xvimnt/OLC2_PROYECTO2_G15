.data
p_2: .byte 1
q_2: .byte 0
.align 2
str0: .asciz "true && false: %s\n"
.align 2
str1: .asciz "true"
.align 2
str2: .asciz "false"
.align 2
str3: .asciz "true || false: %s\n"
.align 2
str4: .asciz "!true: %s\n"
.align 2
str5: .asciz "!false: %s\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of println call ---
    LDR X9, =p_2
    LDRB W10, [X9]
    LDR X9, =q_2
    LDRB W11, [X9]
    AND W10, W10, W11
    LDR X0, =str0
    LDR X9, =str1
    LDR X11, =str2
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =p_2
    LDRB W10, [X9]
    LDR X9, =q_2
    LDRB W11, [X9]
    ORR W10, W10, W11
    LDR X0, =str3
    LDR X9, =str1
    LDR X11, =str2
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =p_2
    LDRB W10, [X9]
    EOR W10, W10, #1
    LDR X0, =str4
    LDR X9, =str1
    LDR X11, =str2
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =q_2
    LDRB W10, [X9]
    EOR W10, W10, #1
    LDR X0, =str5
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

