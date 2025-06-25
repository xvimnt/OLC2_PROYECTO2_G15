.data
x_f_2: .double 10.0
float_one1: .double 1.0
str0: .asciz "float64 --: (10.0 --) =%f\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of float inc/dec on x_f ---
    LDR X10, =x_f_2
    LDR D8, [X10]
    LDR X11, =float_one1
    LDR D9, [X11]
    FSUB D8, D8, D9
    STR D8, [X10]
    // --- End of float inc/dec on x_f ---

    // --- Start of println call ---
    LDR X9, =x_f_2
    LDR D8, [X9]
    LDR X0, =str0
    FMOV D0, D8
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

