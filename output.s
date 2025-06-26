.data
.align 2
str0: .asciz "jurassic"
fruit_2: .quad 0
.align 2
str1: .asciz "banana"
.align 2
str2: .asciz "apple"
.align 2
str3: .asciz "orange"
.align 2
str4: .asciz "It's yellow.\n"
.align 2
str5: .asciz "It's red or green.\n"
.align 2
str6: .asciz "It's orange.\n"
.align 2
str7: .asciz "It's some other fruit.\n"

.extern strcmp
.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    LDR X9, =str0
    // Storing initializer for fruit
    LDR X10, =fruit_2
    STR X9, [X10]
    LDR X9, =fruit_2
    LDR X10, [X9]
    // --- Switch Statement ---
    LDR X9, =str1
    // Comparing with case: StringLiteral: "banana"
    SUB SP, SP, #16
    STR X10, [SP]
    MOV X0, X10
    MOV X1, X9
    BL strcmp
    LDR X10, [SP]
    ADD SP, SP, #16
    CMP W0, #0
    BEQ switch_case_03
    LDR X9, =str2
    // Comparing with case: StringLiteral: "apple"
    SUB SP, SP, #16
    STR X10, [SP]
    MOV X0, X10
    MOV X1, X9
    BL strcmp
    LDR X10, [SP]
    ADD SP, SP, #16
    CMP W0, #0
    BEQ switch_case_14
    LDR X9, =str3
    // Comparing with case: StringLiteral: "orange"
    SUB SP, SP, #16
    STR X10, [SP]
    MOV X0, X10
    MOV X1, X9
    BL strcmp
    LDR X10, [SP]
    ADD SP, SP, #16
    CMP W0, #0
    BEQ switch_case_25
    B switch_default2
    // --- Switch Case Bodies ---
switch_case_03:
    // --- Start of println call ---
    LDR X0, =str4
    BL printf
    // --- End of println call ---
    B switch_end1
switch_case_14:
    // --- Start of println call ---
    LDR X0, =str5
    BL printf
    // --- End of println call ---
    B switch_end1
switch_case_25:
    // --- Start of println call ---
    LDR X0, =str6
    BL printf
    // --- End of println call ---
    B switch_end1
switch_default2:
    // --- Start of println call ---
    LDR X0, =str7
    BL printf
    // --- End of println call ---
switch_end1:
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

