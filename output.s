.data
global_integer: .word 100
str0: .asciz "I am a global variable."
_global_string: .quad str0
globalVarWithCaps: .word 200
str1: .asciz "Called A_Function (uppercase)\n"
str2: .asciz "Called a_function (lowercase)\n"
i_am_valid: .word 1
str3: .asciz "value of 'i_am_valid': %d\n"
str4: .asciz "value of 'global_integer': %d\n"
str5: .asciz "value of 'globalVarWithCaps': %d\n"
str6: .asciz "value of '_global_string': %s\n"

.extern printf
.text
.global A_Function
A_Function:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // Syscall: write(fd=1, buf, count)
    MOV X8, #64
    MOV X0, #1
    LDR X1, =str1
    MOV X2, #30
    SVC #0
.LA_Function_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

.global a_function
a_function:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // Syscall: write(fd=1, buf, count)
    MOV X8, #64
    MOV X0, #1
    LDR X1, =str2
    MOV X2, #30
    SVC #0
.La_function_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    LDR X0, =str3
    LDR X1, =i_am_valid
    LDR W1, [X1]
    BL printf
    LDR X0, =str4
    LDR X1, =global_integer
    LDR W1, [X1]
    BL printf
    LDR X0, =str5
    LDR X1, =globalVarWithCaps
    LDR W1, [X1]
    BL printf
    LDR X0, =str6
    LDR X1, =_global_string
    LDR X1, [X1]
    BL printf
    BL A_Function
    BL a_function
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

