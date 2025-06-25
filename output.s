.data
str0: .asciz "VLang"
str1: .asciz "Cherry"
str2: .asciz "string + string: 'VLang' + 'Cherry' = %s\n"

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
    LDR X9, =str0
    LDR X10, =str1
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
    LDR X0, =str2
    MOV X1, X11
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

