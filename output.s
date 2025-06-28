.data
puntos_2: .quad 0
puntos_entornos_2: .quad 0
a_2: .quad 0
.align 2
str0: .asciz "a = %d\n"
.align 2
str1: .asciz "OK a = 10\n"
b_2: .quad 0
.align 2
str2: .asciz "b = %d\n"
.align 2
str3: .asciz "OK b = 20\n"
c_2: .quad 0
d_2: .quad 0
.align 2
str4: .asciz "c = %d\n"
.align 2
str5: .asciz "d = %d\n"
.align 2
str6: .asciz "OK c = 30\n"
puntos_if_2: .quad 0
.align 2
str7: .asciz "OK true\n"
.align 2
str8: .asciz "OK 1 == 1\n"
.align 2
str9: .asciz "OK 2 > 1\n"
puntos_while_2: .quad 0
i_2: .quad 0
suma1_2: .quad 0
.align 2
str10: .asciz "%d\n"
.align 2
str11: .asciz "OK suma1 == 10\n"
.align 2
str12: .asciz "OK i == 5\n"
j_2: .quad 0
.align 2
str13: .asciz "%d\n"
k_2: .quad 0
.align 2
str14: .asciz "%d\n"
puntos_for_2: .quad 0
suma2_2: .quad 0
x_17: .quad 0
.align 2
str15: .asciz "%d\n"
.align 2
str16: .asciz "OK suma2 == 10\n"
y_20: .quad 0
.align 2
str17: .asciz "%d\n"
z_22: .quad 0
.align 2
str18: .asciz "%d\n"
puntos_case_2: .quad 0
dia_2: .quad 0
.align 2
str19: .asciz "Lunes\n"
.align 2
str20: .asciz "Martes\n"
.align 2
str21: .asciz "Miércoles\n"
.align 2
str22: .asciz "Jueves\n"
.align 2
str23: .asciz "Viernes\n"
.align 2
str24: .asciz "Sábado\n"
.align 2
str25: .asciz "Domingo\n"
.align 2
str26: .asciz "Día inválido\n"
puntos_break_2: .quad 0
suma3_2: .quad 0
n_24: .quad 0
.align 2
str27: .asciz "%d\n"
.align 2
str28: .asciz "OK suma3 == 10\n"
puntos_continue_2: .quad 0
suma_pares_2: .quad 0
m_28: .quad 0
.align 2
str29: .asciz "%d\n"
.align 2
str30: .asciz "OK suma_pares == 20\n"
.align 2
str31: .asciz "Puntos totales: %d / 26\n"

.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    MOV X9, #0
    // Storing initializer for puntos
    LDR X10, =puntos_2
    STR W9, [X10]
    MOV X9, #0
    // Storing initializer for puntos_entornos
    LDR X10, =puntos_entornos_2
    STR W9, [X10]
    MOV X9, #10
    // Storing initializer for a
    LDR X10, =a_2
    STR W9, [X10]
    // --- Start of println call ---
    LDR X9, =a_2
    LDRSW X10, [X9]
    LDR X0, =str0
    MOV X1, X10
    BL printf
    // --- End of println call ---
    LDR X9, =a_2
    LDRSW X10, [X9]
    MOV X9, #10
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif2
    // 'Then' block
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_entornos_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    // --- Start of println call ---
    LDR X0, =str1
    BL printf
    // --- End of println call ---
.Lendif2:
    MOV X9, #10
    // Storing initializer for b
    LDR X10, =b_2
    STR W9, [X10]
    MOV X9, #20
    // Storing value for assignment to b
    LDR X10, =b_2
    STR W9, [X10]
    // --- Start of println call ---
    LDR X9, =b_2
    LDRSW X10, [X9]
    LDR X0, =str2
    MOV X1, X10
    BL printf
    // --- End of println call ---
    LDR X9, =b_2
    LDRSW X10, [X9]
    MOV X9, #20
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif4
    // 'Then' block
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_entornos_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    // --- Start of println call ---
    LDR X0, =str3
    BL printf
    // --- End of println call ---
.Lendif4:
    MOV X9, #10
    // Storing initializer for c
    LDR X10, =c_2
    STR W9, [X10]
    MOV X9, #10
    // Storing initializer for d
    LDR X10, =d_2
    STR W9, [X10]
    MOV X9, #30
    // Storing value for assignment to c
    LDR X10, =c_2
    STR W9, [X10]
    // --- Start of println call ---
    LDR X9, =c_2
    LDRSW X10, [X9]
    LDR X0, =str4
    MOV X1, X10
    BL printf
    // --- End of println call ---
    // --- Start of println call ---
    LDR X9, =d_2
    LDRSW X10, [X9]
    LDR X0, =str5
    MOV X1, X10
    BL printf
    // --- End of println call ---
    LDR X9, =c_2
    LDRSW X10, [X9]
    MOV X9, #30
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif6
    // 'Then' block
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_entornos_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    // --- Start of println call ---
    LDR X0, =str6
    BL printf
    // --- End of println call ---
.Lendif6:
    MOV X9, #0
    // Storing initializer for puntos_if
    LDR X10, =puntos_if_2
    STR W9, [X10]
    MOV X9, #1
    // If statement condition check
    CMP W9, #0
    B.EQ .Lendif8
    // 'Then' block
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_if_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    // --- Start of println call ---
    LDR X0, =str7
    BL printf
    // --- End of println call ---
.Lendif8:
    MOV X9, #1
    MOV X10, #1
    CMP W9, W10
    CSET W9, EQ
    // If statement condition check
    CMP W9, #0
    B.EQ .Lendif10
    // 'Then' block
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_if_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    // --- Start of println call ---
    LDR X0, =str8
    BL printf
    // --- End of println call ---
.Lendif10:
    MOV X9, #2
    MOV X10, #1
    CMP W9, W10
    CSET W9, GT
    // If statement condition check
    CMP W9, #0
    B.EQ .Lendif12
    // 'Then' block
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_if_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    // --- Start of println call ---
    LDR X0, =str9
    BL printf
    // --- End of println call ---
.Lendif12:
    MOV X9, #0
    // Storing initializer for puntos_while
    LDR X10, =puntos_while_2
    STR W9, [X10]
    MOV X9, #0
    // Storing initializer for i
    LDR X10, =i_2
    STR W9, [X10]
    MOV X9, #0
    // Storing initializer for suma1
    LDR X10, =suma1_2
    STR W9, [X10]
loop_start13:
    LDR X9, =i_2
    LDRSW X10, [X9]
    MOV X9, #5
    CMP W10, W9
    CSET W10, LT
    // For loop condition check
    CMP W10, #0
    B.EQ loop_end16
    B loop_body14
loop_post15:
    B loop_start13
loop_body14:
    // --- Start of println call ---
    LDR X9, =i_2
    LDRSW X10, [X9]
    LDR X0, =str10
    MOV X1, X10
    BL printf
    // --- End of println call ---
    LDR X9, =i_2
    LDRSW X10, [X9]
    // --- Start Compound Assignment: += ---
    LDR X9, =suma1_2
    LDRSW X11, [X9]
    ADD X11, X11, X10
    STR W11, [X9]
    // --- End Compound Assignment ---
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =i_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    B loop_post15
loop_end16:
    LDR X9, =suma1_2
    LDRSW X10, [X9]
    MOV X9, #10
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif18
    // 'Then' block
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_while_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    // --- Start of println call ---
    LDR X0, =str11
    BL printf
    // --- End of println call ---
.Lendif18:
    LDR X9, =i_2
    LDRSW X10, [X9]
    MOV X9, #5
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif20
    // 'Then' block
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_while_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    // --- Start of println call ---
    LDR X0, =str12
    BL printf
    // --- End of println call ---
.Lendif20:
    MOV X9, #3
    // Storing initializer for j
    LDR X10, =j_2
    STR W9, [X10]
loop_start21:
    LDR X9, =j_2
    LDRSW X10, [X9]
    MOV X9, #0
    CMP W10, W9
    CSET W10, GT
    // For loop condition check
    CMP W10, #0
    B.EQ loop_end24
    B loop_body22
loop_post23:
    B loop_start21
loop_body22:
    // --- Start of println call ---
    LDR X9, =j_2
    LDRSW X10, [X9]
    LDR X0, =str13
    MOV X1, X10
    BL printf
    // --- End of println call ---
    MOV X9, #1
    // --- Start Compound Assignment: -= ---
    LDR X10, =j_2
    LDRSW X11, [X10]
    SUB X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    B loop_post23
loop_end24:
    MOV X9, #0
    // Storing initializer for k
    LDR X10, =k_2
    STR W9, [X10]
loop_start25:
    LDR X9, =k_2
    LDRSW X10, [X9]
    MOV X9, #10
    CMP W10, W9
    CSET W10, LE
    // For loop condition check
    CMP W10, #0
    B.EQ loop_end28
    B loop_body26
loop_post27:
    B loop_start25
loop_body26:
    // --- Start of println call ---
    LDR X9, =k_2
    LDRSW X10, [X9]
    LDR X0, =str14
    MOV X1, X10
    BL printf
    // --- End of println call ---
    MOV X9, #2
    // --- Start Compound Assignment: += ---
    LDR X10, =k_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    B loop_post27
loop_end28:
    MOV X9, #0
    // Storing initializer for puntos_for
    LDR X10, =puntos_for_2
    STR W9, [X10]
    MOV X9, #0
    // Storing initializer for suma2
    LDR X10, =suma2_2
    STR W9, [X10]
    MOV X9, #0
    // Storing initializer for x
    LDR X10, =x_17
    STR W9, [X10]
loop_start29:
    LDR X9, =x_17
    LDRSW X10, [X9]
    MOV X9, #5
    CMP W10, W9
    CSET W10, LT
    // For loop condition check
    CMP W10, #0
    B.EQ loop_end32
    B loop_body30
loop_post31:
    // --- Start of integer inc/dec on x ---
    LDR X10, =x_17
    LDR W11, [X10]
    ADD W11, W11, #1
    STR W11, [X10]
    // --- End of integer inc/dec on x ---

    B loop_start29
loop_body30:
    // --- Start of println call ---
    LDR X9, =x_17
    LDRSW X10, [X9]
    LDR X0, =str15
    MOV X1, X10
    BL printf
    // --- End of println call ---
    LDR X9, =x_17
    LDRSW X10, [X9]
    // --- Start Compound Assignment: += ---
    LDR X9, =suma2_2
    LDRSW X11, [X9]
    ADD X11, X11, X10
    STR W11, [X9]
    // --- End Compound Assignment ---
    B loop_post31
loop_end32:
    LDR X9, =suma2_2
    LDRSW X10, [X9]
    MOV X9, #10
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif34
    // 'Then' block
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_for_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    // --- Start of println call ---
    LDR X0, =str16
    BL printf
    // --- End of println call ---
.Lendif34:
    MOV X9, #0
    // Storing initializer for y
    LDR X10, =y_20
    STR W9, [X10]
loop_start35:
    LDR X9, =y_20
    LDRSW X10, [X9]
    MOV X9, #3
    CMP W10, W9
    CSET W10, LT
    // For loop condition check
    CMP W10, #0
    B.EQ loop_end38
    B loop_body36
loop_post37:
    // --- Start of integer inc/dec on y ---
    LDR X10, =y_20
    LDR W11, [X10]
    ADD W11, W11, #1
    STR W11, [X10]
    // --- End of integer inc/dec on y ---

    B loop_start35
loop_body36:
    // --- Start of println call ---
    LDR X9, =y_20
    LDRSW X10, [X9]
    LDR X0, =str17
    MOV X1, X10
    BL printf
    // --- End of println call ---
    B loop_post37
loop_end38:
    MOV X9, #0
    // Storing initializer for z
    LDR X10, =z_22
    STR W9, [X10]
loop_start39:
    LDR X9, =z_22
    LDRSW X10, [X9]
    MOV X9, #2
    CMP W10, W9
    CSET W10, LT
    // For loop condition check
    CMP W10, #0
    B.EQ loop_end42
    B loop_body40
loop_post41:
    // --- Start of integer inc/dec on z ---
    LDR X10, =z_22
    LDR W11, [X10]
    ADD W11, W11, #1
    STR W11, [X10]
    // --- End of integer inc/dec on z ---

    B loop_start39
loop_body40:
    // --- Start of println call ---
    LDR X9, =z_22
    LDRSW X10, [X9]
    LDR X0, =str18
    MOV X1, X10
    BL printf
    // --- End of println call ---
    B loop_post41
loop_end42:
    MOV X9, #2
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_for_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    MOV X9, #0
    // Storing initializer for puntos_case
    LDR X10, =puntos_case_2
    STR W9, [X10]
    MOV X9, #3
    // Storing initializer for dia
    LDR X10, =dia_2
    STR W9, [X10]
    LDR X9, =dia_2
    LDRSW X10, [X9]
    // --- Switch Statement ---
    MOV X9, #1
    // Comparing with case: IntegerLiteral: 1
    CMP W10, W9
    BEQ switch_case_045
    MOV X9, #2
    // Comparing with case: IntegerLiteral: 2
    CMP W10, W9
    BEQ switch_case_146
    MOV X9, #3
    // Comparing with case: IntegerLiteral: 3
    CMP W10, W9
    BEQ switch_case_247
    MOV X9, #4
    // Comparing with case: IntegerLiteral: 4
    CMP W10, W9
    BEQ switch_case_348
    MOV X9, #5
    // Comparing with case: IntegerLiteral: 5
    CMP W10, W9
    BEQ switch_case_449
    MOV X9, #6
    // Comparing with case: IntegerLiteral: 6
    CMP W10, W9
    BEQ switch_case_550
    MOV X9, #7
    // Comparing with case: IntegerLiteral: 7
    CMP W10, W9
    BEQ switch_case_651
    B switch_default44
    // --- Switch Case Bodies ---
switch_case_045:
    // --- Start of println call ---
    LDR X0, =str19
    BL printf
    // --- End of println call ---
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X11, =puntos_case_2
    LDRSW X12, [X11]
    ADD X12, X12, X9
    STR W12, [X11]
    // --- End Compound Assignment ---
    B switch_end43
switch_case_146:
    // --- Start of println call ---
    LDR X0, =str20
    BL printf
    // --- End of println call ---
    B switch_end43
switch_case_247:
    // --- Start of println call ---
    LDR X0, =str21
    BL printf
    // --- End of println call ---
    MOV X9, #1
    // --- Start Compound Assignment: += ---
    LDR X11, =puntos_case_2
    LDRSW X12, [X11]
    ADD X12, X12, X9
    STR W12, [X11]
    // --- End Compound Assignment ---
    B switch_end43
switch_case_348:
    // --- Start of println call ---
    LDR X0, =str22
    BL printf
    // --- End of println call ---
    B switch_end43
switch_case_449:
    // --- Start of println call ---
    LDR X0, =str23
    BL printf
    // --- End of println call ---
    B switch_end43
switch_case_550:
    // --- Start of println call ---
    LDR X0, =str24
    BL printf
    // --- End of println call ---
    B switch_end43
switch_case_651:
    // --- Start of println call ---
    LDR X0, =str25
    BL printf
    // --- End of println call ---
    B switch_end43
switch_default44:
    // --- Start of println call ---
    LDR X0, =str26
    BL printf
    // --- End of println call ---
switch_end43:
    MOV X9, #0
    // Storing initializer for puntos_break
    LDR X10, =puntos_break_2
    STR W9, [X10]
    MOV X9, #0
    // Storing initializer for suma3
    LDR X10, =suma3_2
    STR W9, [X10]
    MOV X9, #0
    // Storing initializer for n
    LDR X10, =n_24
    STR W9, [X10]
loop_start52:
    LDR X9, =n_24
    LDRSW X10, [X9]
    MOV X9, #10
    CMP W10, W9
    CSET W10, LT
    // For loop condition check
    CMP W10, #0
    B.EQ loop_end55
    B loop_body53
loop_post54:
    // --- Start of integer inc/dec on n ---
    LDR X10, =n_24
    LDR W11, [X10]
    ADD W11, W11, #1
    STR W11, [X10]
    // --- End of integer inc/dec on n ---

    B loop_start52
loop_body53:
    LDR X9, =n_24
    LDRSW X10, [X9]
    MOV X9, #5
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif57
    // 'Then' block
    B loop_end55
.Lendif57:
    // --- Start of println call ---
    LDR X9, =n_24
    LDRSW X10, [X9]
    LDR X0, =str27
    MOV X1, X10
    BL printf
    // --- End of println call ---
    LDR X9, =n_24
    LDRSW X10, [X9]
    // --- Start Compound Assignment: += ---
    LDR X9, =suma3_2
    LDRSW X11, [X9]
    ADD X11, X11, X10
    STR W11, [X9]
    // --- End Compound Assignment ---
    B loop_post54
loop_end55:
    LDR X9, =suma3_2
    LDRSW X10, [X9]
    MOV X9, #10
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif59
    // 'Then' block
    MOV X9, #3
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_break_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    // --- Start of println call ---
    LDR X0, =str28
    BL printf
    // --- End of println call ---
.Lendif59:
    MOV X9, #0
    // Storing initializer for puntos_continue
    LDR X10, =puntos_continue_2
    STR W9, [X10]
    MOV X9, #0
    // Storing initializer for suma_pares
    LDR X10, =suma_pares_2
    STR W9, [X10]
    MOV X9, #0
    // Storing initializer for m
    LDR X10, =m_28
    STR W9, [X10]
loop_start60:
    LDR X9, =m_28
    LDRSW X10, [X9]
    MOV X9, #10
    CMP W10, W9
    CSET W10, LT
    // For loop condition check
    CMP W10, #0
    B.EQ loop_end63
    B loop_body61
loop_post62:
    // --- Start of integer inc/dec on m ---
    LDR X10, =m_28
    LDR W11, [X10]
    ADD W11, W11, #1
    STR W11, [X10]
    // --- End of integer inc/dec on m ---

    B loop_start60
loop_body61:
    LDR X9, =m_28
    LDRSW X10, [X9]
    MOV X9, #2
    SDIV X11, X10, X9
    MUL X11, X11, X9
    SUB X10, X10, X11
    MOV X9, #0
    CMP W10, W9
    CSET W10, NE
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif65
    // 'Then' block
    B loop_post62
.Lendif65:
    // --- Start of println call ---
    LDR X9, =m_28
    LDRSW X10, [X9]
    LDR X0, =str29
    MOV X1, X10
    BL printf
    // --- End of println call ---
    LDR X9, =m_28
    LDRSW X10, [X9]
    // --- Start Compound Assignment: += ---
    LDR X9, =suma_pares_2
    LDRSW X11, [X9]
    ADD X11, X11, X10
    STR W11, [X9]
    // --- End Compound Assignment ---
    B loop_post62
loop_end63:
    LDR X9, =suma_pares_2
    LDRSW X10, [X9]
    MOV X9, #20
    CMP W10, W9
    CSET W10, EQ
    // If statement condition check
    CMP W10, #0
    B.EQ .Lendif67
    // 'Then' block
    MOV X9, #3
    // --- Start Compound Assignment: += ---
    LDR X10, =puntos_continue_2
    LDRSW X11, [X10]
    ADD X11, X11, X9
    STR W11, [X10]
    // --- End Compound Assignment ---
    // --- Start of println call ---
    LDR X0, =str30
    BL printf
    // --- End of println call ---
.Lendif67:
    LDR X9, =puntos_entornos_2
    LDRSW X10, [X9]
    LDR X9, =puntos_if_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    LDR X9, =puntos_while_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    LDR X9, =puntos_for_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    LDR X9, =puntos_case_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    LDR X9, =puntos_break_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    LDR X9, =puntos_continue_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    // Storing value for assignment to puntos
    LDR X9, =puntos_2
    STR W10, [X9]
    // --- Start of println call ---
    LDR X9, =puntos_2
    LDRSW X10, [X9]
    LDR X0, =str31
    MOV X1, X10
    BL printf
    // --- End of println call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

