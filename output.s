.data
str0: .asciz "Called A_Function (uppercase)\n"
str1: .asciz "Called a_function (lowercase)\n"

.text
.global A_Function
A_Function:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // Syscall: write(fd=1, buf, count)
    MOV X8, #64
    MOV X0, #1
    LDR X1, =str0
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
    LDR X1, =str1
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
    BL A_Function
    BL a_function
    // Syscall: exit(status=0)
    MOV X8, #93     // exit syscall number
    MOV X0, #0      // exit status code
    SVC #0          // trigger syscall

