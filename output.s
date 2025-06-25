.data
global_integer: .word 100
str0: .asciz "I am a global variable."
_global_string: .quad str0
globalVarWithCaps: .word 200
str1: .asciz "Called A_Function (uppercase)\n"
str2: .asciz "Called a_function (lowercase)\n"
str3: .asciz "--- Test 1: Basics, Scope, and Declarations ---\n"
str4: .asciz "\n--- A. Testing Comments and Identifiers ---\n"
str5: .asciz "This line should execute.\n"
str6: .asciz "This one too.\n"
i_am_valid: .word 1
_can_start_with_underscore: .word 2
var123isFine: .word 3
str7: .asciz "Valid identifiers declared successfully.\n"
str8: .asciz "\n--- B. Testing Case Sensitivity ---\n"
myvar: .word 10
MyVar: .word 20
str9: .asciz "value of 'myvar': %d\n"
str10: .asciz "value of 'MyVar': %d\n"
str11: .asciz "\n--- C. Testing Variable Declarations ---\n"
uninitialized_int: .word 0
uninitialized_float: .word 0
uninitialized_string: .word 0
uninitialized_bool: .word 0
str12: .asciz "Default int value: %d\n"
str13: .asciz "Default float64 value: %d\n"
str14: .asciz "Default string value: '%d'\n"
str15: .asciz "Default bool value: %d\n"
explicit_int: .word -50
str16: .asciz "Explicitly initialized int: %d\n"
str17: .asciz "This is a string"
inferred_string: .quad str17
inferred_float: .double 3.14
str18: .asciz "Inferred string via ':=' %s\n"
str19: .asciz "Inferred float via ':=' %f\n"
mutable_inferred_bool: .byte 1
str20: .asciz "Inferred mutable bool via 'mut ... :=' %s\n"
str21: .asciz "true"
str22: .asciz "false"
str23: .asciz "Modified mutable bool: %s\n"
str24: .asciz "true"
str25: .asciz "false"

.extern printf
.text
.global A_Function
A_Function:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of println call ---
    LDR X0, =str1
    BL printf
    // --- End of println call ---

.LA_Function_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

.global a_function
a_function:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of println call ---
    LDR X0, =str2
    BL printf
    // --- End of println call ---

.La_function_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    // --- Start of println call ---
    LDR X0, =str3
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str4
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str5
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str6
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str7
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str8
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str9
    LDR X9, =myvar
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str10
    LDR X9, =MyVar
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    BL A_Function
    BL a_function
    // --- Start of println call ---
    LDR X0, =str11
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str12
    LDR X9, =uninitialized_int
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str13
    LDR X9, =uninitialized_float
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str14
    LDR X9, =uninitialized_string
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str15
    LDR X9, =uninitialized_bool
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str16
    LDR X9, =explicit_int
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str18
    LDR X1, =inferred_string
    LDR X1, [X1]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str19
    LDR X9, =inferred_float
    LDR D0, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str20
    LDR X10, =str21
    LDR X11, =str22
    LDR X12, =mutable_inferred_bool
    LDRB W9, [X12]
    CMP W9, #0
    CSEL X1, X11, X10, EQ
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str23
    LDR X10, =str24
    LDR X11, =str25
    LDR X12, =mutable_inferred_bool
    LDRB W9, [X12]
    CMP W9, #0
    CSEL X1, X11, X10, EQ
    BL printf
    // --- End of println call ---

.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

