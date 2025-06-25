.data
uninitialized_int: .word 0
uninitialized_float: .word 0
uninitialized_string: .word 0
uninitialized_bool: .word 0
str0: .asciz "Default int value:%d\n"
str1: .asciz "Default float64 value:%d\n"
str2: .asciz "Default string value: '%d'\n"
str3: .asciz "Default bool value:%d\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // Println call
    LDR X0, =str0
    LDR X9, =uninitialized_int
    LDR W1, [X9]
    BL printf
    // Println call
    LDR X0, =str1
    LDR X9, =uninitialized_float
    LDR W1, [X9]
    BL printf
    // Println call
    LDR X0, =str2
    LDR X9, =uninitialized_string
    LDR W1, [X9]
    BL printf
    // Println call
    LDR X0, =str3
    LDR X9, =uninitialized_bool
    LDR W1, [X9]
    BL printf
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

