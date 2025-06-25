.data
global_integer: .word 100

.text
.global _start
_start:
    STP X29, X30, [SP, #-16]
    MOV X29, SP
    MOV X8, #93
    MOV X0, #0
    SVC #0

