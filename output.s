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
str11: .asciz "\nint * int:      10 * 3 = %d\n"
str12: .asciz "float64 * int:  12.5 * 10 = %f\n"
str13: .asciz "\nint / int (truncation): 10 / 3 = %d\n"
str14: .asciz "float64 / float64:      12.5 / 2.5 = %f\n"
str15: .asciz "int / float64:          10 / 2.5 = %f\n"
str16: .asciz "\nint %% int:      10 %% 3 = %d\n"
str17: .asciz "\nUnary negation:   -10 = %d\n"
str18: .asciz "Unary negation: -12.5 = %f\n"
str19: .asciz "\n--- B. Compound Assignment Operators ---\n"
x_i_2: .word 20
str20: .asciz "int += int: (20 += 5) =%d\n"
str21: .asciz "int ++: (25++) =%d\n"
str22: .asciz "int -= int: (25 -= 10) =%d\n"
x_f_2: .double 10.0
str23: .asciz "float64 += int: (10.0 += 5) =%f\n"
str24: .asciz "test"
x_s_2: .quad str24
str25: .asciz "_file"
str26: .asciz "string += string: ('test' += '_file') =%s\n"

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
    // --- Start of println call ---
    LDR X9, =i1_2
    LDRSW X10, [X9]
    LDR X9, =i2_2
    LDRSW X11, [X9]
    MUL X10, X10, X11
    LDR X0, =str11
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
    FMUL D8, D8, D9
    LDR X0, =str12
    FMOV D0, D8
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =i1_2
    LDRSW X10, [X9]
    LDR X9, =i2_2
    LDRSW X11, [X9]
    SDIV X10, X10, X11
    LDR X0, =str13
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =f1_2
    LDR D8, [X9]
    LDR X9, =f2_2
    LDR D9, [X9]
    FDIV D8, D8, D9
    LDR X0, =str14
    FMOV D0, D8
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =i1_2
    LDRSW X10, [X9]
    LDR X9, =f2_2
    LDR D8, [X9]
    // Promoting left operand from INT to FLOAT
    SCVTF D9, X10
    FDIV D9, D9, D8
    LDR X0, =str15
    FMOV D0, D9
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =i1_2
    LDRSW X10, [X9]
    LDR X9, =i2_2
    LDRSW X11, [X9]
    SDIV X9, X10, X11
    MUL X9, X9, X11
    SUB X10, X10, X9
    LDR X0, =str16
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =i1_2
    LDRSW X10, [X9]
    NEG X10, X10
    LDR X0, =str17
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =f1_2
    LDR D8, [X9]
    FNEG D8, D8
    LDR X0, =str18
    FMOV D0, D8
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X0, =str19
    BL printf
    // --- End of println call ---
    // --- Start of compound assignment (+=) to x_i ---
    LDR X10, =x_i_2
    LDR W11, [X10]
    MOV W12, #5
    ADD W11, W11, W12
    STR W11, [X10]
    // --- End of compound assignment (+=) to x_i ---

    // --- Start of println call ---
    LDR X9, =x_i_2
    LDRSW X10, [X9]
    LDR X0, =str20
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of integer inc/dec on x_i ---
    LDR X10, =x_i_2
    LDR W11, [X10]
    ADD W11, W11, #1
    STR W11, [X10]
    // --- End of integer inc/dec on x_i ---

    // --- Start of println call ---
    LDR X9, =x_i_2
    LDRSW X10, [X9]
    LDR X0, =str21
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of compound assignment (-=) to x_i ---
    LDR X10, =x_i_2
    LDR W11, [X10]
    MOV W12, #10
    SUB W11, W11, W12
    STR W11, [X10]
    // --- End of compound assignment (-=) to x_i ---

    // --- Start of println call ---
    LDR X9, =x_i_2
    LDRSW X10, [X9]
    LDR X0, =str22
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of compound assignment (+=) to x_f ---
    LDR X10, =x_f_2
    LDR D8, [X10]
    MOV W11, #5
    SCVTF D9, W11
    FADD D8, D8, D9
    STR D8, [X10]
    // --- End of compound assignment (+=) to x_f ---

    // --- Start of println call ---
    LDR X9, =x_f_2
    LDR D8, [X9]
    LDR X0, =str23
    FMOV D0, D8
    BL printf
    // --- End of println call ---
    LDR X9, =x_s_2
    LDR X10, [X9]
    LDR X11, =str25
    // --- Start of string concatenation ---
    STP X19, X20, [SP, #-16]!
    STP X21, X30, [SP, #-16]!
    MOV X19, X10  // Pointer to left string
    MOV X20, X11  // Pointer to right string
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
    MOV X12, X21
    LDP X21, X30, [SP], #16
    LDP X19, X20, [SP], #16
    // --- End of string concatenation ---
    STR X12, [X9]
    // --- Start of println call ---
    LDR X9, =x_s_2
    LDR X10, [X9]
    LDR X0, =str26
    MOV X1, X10
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

