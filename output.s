.data
slice_data1: .quad 10, 20, 30
slice_descriptor2:
    .quad slice_data1  // Pointer to data
    .quad 3    // Length
    .quad 3    // Capacity
numbers_2: .quad 0

.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    LDR X9, =slice_descriptor2
    // Storing initializer for numbers
    LDR X10, =numbers_2
    STR X9, [X10]
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

