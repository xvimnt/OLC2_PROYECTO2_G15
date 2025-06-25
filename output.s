.data
explicit_int: .word -50
str0: .asciz "Explicitly initialized int:%d\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // Println call
    LDR X0, =str0
    LDR X9, =explicit_int
    LDR W1, [X9]
    BL printf
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

