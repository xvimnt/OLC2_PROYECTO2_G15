.data
str0: .asciz "Hello from the translator!, this should be translated to ARM assembly\n"

.text
.global _start
_start:
    STP X29, X30, [SP, #-16]
    MOV X29, SP
    // Syscall: write(fd=1, buf, count)
    MOV X8, #64
    MOV X0, #1
    LDR X1, =str0
    MOV X2, #70
    SVC #0
    // Syscall: exit(code=0)
    MOV X8, #93
    MOV X0, #0
    SVC #0

