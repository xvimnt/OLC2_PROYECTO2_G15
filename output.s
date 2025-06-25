.data
str0: .asciz "I am a global variable."
_global_string: .quad str0
str_fmt: .asciz "%s\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]
    MOV X29, SP
    LDR X0, =str_fmt
    LDR X1, =_global_string
    LDR X1, [X1]
    BL printf
    // Return from main, letting C runtime handle exit
    MOV W0, #0      // Return 0 from main
    RET

