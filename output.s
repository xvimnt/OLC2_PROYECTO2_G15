.data
global_integer_0: .word 100
str0: .asciz "I am a global variable."
_global_string_0: .quad str0
globalVarWithCaps_0: .word 200
str1: .asciz "Called A_Function (uppercase)\n"
str2: .asciz "Called a_function (lowercase)\n"
str3: .asciz "--- Test 1: Basics, Scope, and Declarations ---\n"
str4: .asciz "\n--- A. Testing Comments and Identifiers ---\n"
str5: .asciz "This line should execute.\n"
str6: .asciz "This one too.\n"
i_am_valid_6: .word 1
_can_start_with_underscore_6: .word 2
var123isFine_6: .word 3
str7: .asciz "Valid identifiers declared successfully.\n"
str8: .asciz "\n--- B. Testing Case Sensitivity ---\n"
myvar_6: .word 10
MyVar_6: .word 20
str9: .asciz "value of 'myvar': %d\n"
str10: .asciz "value of 'MyVar': %d\n"
str11: .asciz "\n--- C. Testing Variable Declarations ---\n"
uninitialized_int_6: .word 0
uninitialized_float_6: .word 0
uninitialized_string_6: .word 0
uninitialized_bool_6: .word 0
str12: .asciz "Default int value: %d\n"
str13: .asciz "Default float64 value: %d\n"
str14: .asciz "Default string value: '%d'\n"
str15: .asciz "Default bool value: %d\n"
explicit_int_6: .word -50
str16: .asciz "Explicitly initialized int: %d\n"
str17: .asciz "This is a string"
inferred_string_6: .quad str17
inferred_float_6: .double 3.14
str18: .asciz "Inferred string via ':=' %s\n"
str19: .asciz "Inferred float via ':=' %f\n"
mutable_inferred_bool_6: .byte 1
str20: .asciz "Inferred mutable bool via 'mut ... :=' %s\n"
str21: .asciz "true"
str22: .asciz "false"
str23: .asciz "Modified mutable bool: %s\n"
str24: .asciz "true"
str25: .asciz "false"
str26: .asciz "Reassigned int: %d\n"
str27: .asciz "\n--- D. Testing Scope Rules and Shadowing ---\n"
outer_scope_var_6: .word 1
str28: .asciz "1. outer_scope_var in main scope: %d\n"
str29: .asciz "2. Accessing global from main scope: %d\n"
inner_scope_var_7: .word 2
str30: .asciz "3. outer_scope_var inside block 1: %d\n"
str31: .asciz "4. inner_scope_var inside block 1: %d\n"
outer_scope_var_7: .word 99
str32: .asciz "5. Shadowed outer_scope_var in block 1: %d\n"
deepest_var_8: .byte 1
str33: .asciz "6. Accessing inner_scope_var from block 2: %d\n"
str34: .asciz "7. Accessing shadowed var from block 2: %d\n"
str35: .asciz "8. outer_scope_var back in main scope: %d\n"
str36: .asciz "\n--- Test 1 Finished ---\n"

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
    LDR X9, =myvar_6
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str10
    LDR X9, =MyVar_6
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
    LDR X9, =uninitialized_int_6
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str13
    LDR X9, =uninitialized_float_6
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str14
    LDR X9, =uninitialized_string_6
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str15
    LDR X9, =uninitialized_bool_6
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str16
    LDR X9, =explicit_int_6
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str18
    LDR X1, =inferred_string_6
    LDR X1, [X1]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str19
    LDR X9, =inferred_float_6
    LDR D0, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str20
    LDR X10, =str21
    LDR X11, =str22
    LDR X12, =mutable_inferred_bool_6
    LDRB W9, [X12]
    CMP W9, #0
    CSEL X1, X11, X10, EQ
    BL printf
    // --- End of println call ---

    // --- Start of assignment to mutable_inferred_bool ---
    LDR X10, =mutable_inferred_bool_6
    MOV W11, #0
    STRB W11, [X10]
    // --- End of assignment to mutable_inferred_bool ---

    // --- Start of println call ---
    LDR X0, =str23
    LDR X10, =str24
    LDR X11, =str25
    LDR X12, =mutable_inferred_bool_6
    LDRB W9, [X12]
    CMP W9, #0
    CSEL X1, X11, X10, EQ
    BL printf
    // --- End of println call ---

    // --- Start of assignment to explicit_int ---
    LDR X10, =explicit_int_6
    MOV W11, #100
    STR W11, [X10]
    // --- End of assignment to explicit_int ---

    // --- Start of println call ---
    LDR X0, =str26
    LDR X9, =explicit_int_6
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str27
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str28
    LDR X9, =outer_scope_var_6
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str29
    LDR X9, =global_integer_0
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str30
    LDR X9, =outer_scope_var_6
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str31
    LDR X9, =inner_scope_var_7
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str32
    LDR X9, =outer_scope_var_7
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str33
    LDR X9, =inner_scope_var_7
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str34
    LDR X9, =outer_scope_var_7
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str35
    LDR X9, =outer_scope_var_6
    LDR W1, [X9]
    BL printf
    // --- End of println call ---

    // --- Start of println call ---
    LDR X0, =str36
    BL printf
    // --- End of println call ---

.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

