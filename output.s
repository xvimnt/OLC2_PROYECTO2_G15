.data
global_integer: .word 100
int_fmt: .asciz "%d\n"

.extern printf
.text
.global _start
_start:
    STP X29, X30, [SP, #-16]
    MOV X29, SP
    // Print integer variable 'global_integer' using printf
    LDR X0, =int_fmt
    LDR X1, =global_integer
    LDR W1, [X1]
    BL printf
    // Syscall: exit(code=0)
    MOV X8, #93
    MOV X0, #0
    SVC #0

