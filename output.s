.data
.align 2
str0: .asciz "a"
.align 2
str1: .asciz "a"
.align 2
str2: .asciz "string == string: ('a' == 'a') is %s\n"
.align 2
str3: .asciz "true"
.align 2
str4: .asciz "false"
.align 2
str5: .asciz "a"
.align 2
str6: .asciz "A"
.align 2
str7: .asciz "string != string: ('a' != 'A') is %s\n"
.align 2
str8: .asciz "bool == bool: (true == true) is %s\n"
.align 2
str9: .asciz "rune == rune: ('z' == 'z') is %s\n"
.align 3
F10: .double 10.0
.align 2
str11: .asciz "float == int: (10.0 == 10) is %s\n"
.align 3
F12: .double 10.1
.align 2
str13: .asciz "int != float: (10 != 10.1) is %s\n"

.extern strcmp
.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of println call ---
    LDR X9, =str0
    LDR X10, =str1
    MOV X0, X9
    MOV X1, X10
    BL strcmp
    CMP W0, #0
    CSET X11, EQ
    LDR X0, =str2
    LDR X9, =str3
    LDR X10, =str4
    CMP X11, #0
    CSEL X1, X9, X10, NE
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =str5
    LDR X10, =str6
    MOV X0, X9
    MOV X1, X10
    BL strcmp
    CMP W0, #0
    CSET X11, NE
    LDR X0, =str7
    LDR X9, =str3
    LDR X10, =str4
    CMP X11, #0
    CSEL X1, X9, X10, NE
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    MOV X9, #1
    MOV X10, #1
    CMP X9, X10
    CSET X9, EQ
    LDR X0, =str8
    LDR X10, =str3
    LDR X11, =str4
    CMP X9, #0
    CSEL X1, X10, X11, NE
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    MOV X9, #122
    MOV X10, #122
    CMP X9, X10
    CSET X9, EQ
    LDR X0, =str9
    LDR X10, =str3
    LDR X11, =str4
    CMP X9, #0
    CSEL X1, X10, X11, NE
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR D8, F10
    MOV X9, #10
    // Promoting right operand from INT to FLOAT
    SCVTF D9, X9
    FCMP D8, D9
    CSET X9, EQ
    LDR X0, =str11
    LDR X10, =str3
    LDR X11, =str4
    CMP X9, #0
    CSEL X1, X10, X11, NE
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    MOV X9, #10
    LDR D8, F12
    // Promoting left operand from INT to FLOAT
    SCVTF D9, X9
    FCMP D9, D8
    CSET X9, NE
    LDR X0, =str13
    LDR X10, =str3
    LDR X11, =str4
    CMP X9, #0
    CSEL X1, X10, X11, NE
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

