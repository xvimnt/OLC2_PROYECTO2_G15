.data
str0: .asciz "--- Test 2: Comprehensive Operator Tests ---\n"
str1: .asciz "\n--- A. Arithmetic Operators ---\n"
i1_2: .word 10
i2_2: .word 3
f1_2: .double 12.5
f2_2: .double 2.5
str2: .asciz "int + int:      10 + 3 = %d\n"
str3: .asciz "int + float64:  10 + 2.5 = %f\n"
str4: .asciz "float64 + int:  12.5 + 3 = %f\n"
str5: .asciz "float64 + float64: 12.5 + 2.5 = %f\n"
str6: .asciz "VLang"
str7: .asciz "Cherry"
str8: .asciz "string + string: 'VLang' + 'Cherry' = %s\n"
str9: .asciz "\nint - int:      10 - 3 = %d\n"
str10: .asciz "float64 - int:  12.5 - 10 = %f\n"

.extern malloc
.extern strlen
.extern strcpy
.extern strcat
.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of println call ---
    LDR X0, =str0
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X0, =str1
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =i1_2
    LDRSW X10, [X9]
    LDR X9, =i2_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    LDR X0, =str2
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =i1_2
    LDRSW X10, [X9]
    LDR X9, =f2_2
    LDR D8, [X9]
    // Promoting left operand from INT to FLOAT
    SCVTF D9, X10
    FADD D9, D9, D8
    LDR X0, =str3
    FMOV D0, D9
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =f1_2
    LDR D8, [X9]
    LDR X9, =i2_2
    LDRSW X10, [X9]
    // Promoting right operand from INT to FLOAT
    SCVTF D9, X10
    FADD D8, D8, D9
    LDR X0, =str4
    FMOV D0, D8
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =f1_2
    LDR D8, [X9]
    LDR X9, =f2_2
    LDR D9, [X9]
    FADD D8, D8, D9
    LDR X0, =str5
    FMOV D0, D8
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =str6
    LDR X10, =str7
    // --- Start of string concatenation ---
    STP X19, X20, [SP, #-16]!
    STP X21, X30, [SP, #-16]!
    MOV X19, X9  // Pointer to left string
    MOV X20, X10  // Pointer to right string
    MOV X0, X19
    BL strlen
    MOV X21, X0  // Store length of left string
    MOV X0, X20
    BL strlen
    ADD X0, X0, X21  // Total length
    ADD X0, X0, #1     // Add 1 for null terminator
    BL malloc
    MOV X21, X0      // X21 now holds the new string pointer
    MOV X0, X21      // 1st arg for strcpy: destination
    MOV X1, X19      // 2nd arg for strcpy: source (left string)
    BL strcpy
    MOV X0, X21      // 1st arg for strcat: destination
    MOV X1, X20      // 2nd arg for strcat: source (right string)
    BL strcat
    MOV X11, X21
    LDP X21, X30, [SP], #16
    LDP X19, X20, [SP], #16
    // --- End of string concatenation ---
    LDR X0, =str8
    MOV X1, X11
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =i1_2
    LDRSW X10, [X9]
    LDR X9, =i2_2
    LDRSW X11, [X9]
    SUB X10, X10, X11
    LDR X0, =str9
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =f1_2
    LDR D8, [X9]
    LDR X9, =i1_2
    LDRSW X10, [X9]
    // Promoting right operand from INT to FLOAT
    SCVTF D9, X10
    FSUB D8, D8, D9
    LDR X0, =str10
    FMOV D0, D8
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

