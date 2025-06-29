.data
puntos_2: .quad 0
.align 2
str0: .asciz "=== Archivo de prueba básico ==="
.align 2
str1: .asciz "%s"
.align 2
str2: .asciz "\n"
.align 2
str3: .asciz "Validaciones manuales esperadas: 3"
.align 2
str4: .asciz "%s"
.align 2
str5: .asciz "\n"
.align 2
str6: .asciz "==== Declaración de variables ===="
.align 2
str7: .asciz "%s"
.align 2
str8: .asciz "\n"
puntosDeclaracion_2: .quad 0
.align 2
str9: .asciz "Declaración explícita con tipo y valor"
.align 2
str10: .asciz "%s"
.align 2
str11: .asciz "\n"
entero_2: .quad 0
.align 3
F12: .double 3.14
decimal_2: .double 0.0
.align 2
str13: .asciz "Hola!"
texto_2: .quad 0
booleano_2: .quad 0
.align 2
str14: .asciz "\n\n###Validacion Manual"
.align 2
str15: .asciz "%s"
.align 2
str16: .asciz "\n"
.align 2
str17: .asciz "entero:"
.align 2
str18: .asciz "%s"
.align 2
str19: .asciz " "
.align 2
str20: .asciz "%d"
.align 2
str21: .asciz "\n"
.align 2
str22: .asciz "decimal:"
.align 2
str23: .asciz "%s"
.align 2
str24: .asciz " "
.align 2
str25: .asciz "%f"
.align 2
str26: .asciz "\n"
.align 2
str27: .asciz "texto:"
.align 2
str28: .asciz "%s"
.align 2
str29: .asciz " "
.align 2
str30: .asciz "%s"
.align 2
str31: .asciz "\n"
.align 2
str32: .asciz "booleano:"
.align 2
str33: .asciz "%s"
.align 2
str34: .asciz " "
.align 2
str35: .asciz "%s"
.align 2
str36: .asciz "true"
.align 2
str37: .asciz "false"
.align 2
str38: .asciz "\n"
.align 2
str39: .asciz ""
.align 2
str40: .asciz "%s"
.align 2
str41: .asciz "\n"
.align 3
F42: .double 3.0
.align 2
str43: .asciz "Hola!"
.align 2
str44: .asciz "OK Declaración explícita: correcto"
.align 2
str45: .asciz "%s"
.align 2
str46: .asciz "\n"
.align 2
str47: .asciz "X Declaración explícita: incorrecto"
.align 2
str48: .asciz "%s"
.align 2
str49: .asciz "\n"
.align 2
str50: .asciz "Declaración sin valor"
.align 2
str51: .asciz "%s"
.align 2
str52: .asciz "\n"
enteroSinValor_2: .quad 0
.align 3
decimalSinValor_2: .double 0.0
.align 2
str53: .asciz ""
textoSinValor_2: .quad str53
booleanoSinValor_2: .quad 0
.align 2
str54: .asciz "enteroSinValor:"
.align 2
str55: .asciz "%s"
.align 2
str56: .asciz " "
.align 2
str57: .asciz "%d"
.align 2
str58: .asciz "\n"
.align 2
str59: .asciz "decimalSinValor:"
.align 2
str60: .asciz "%s"
.align 2
str61: .asciz " "
.align 2
str62: .asciz "%f"
.align 2
str63: .asciz "\n"
.align 2
str64: .asciz "textoSinValor:"
.align 2
str65: .asciz "%s"
.align 2
str66: .asciz " "
.align 2
str67: .asciz "%s"
.align 2
str68: .asciz "\n"
.align 2
str69: .asciz "booleanoSinValor:"
.align 2
str70: .asciz "%s"
.align 2
str71: .asciz " "
.align 2
str72: .asciz "%s"
.align 2
str73: .asciz "\n"
.align 3
F74: .double 0.0
.align 2
str75: .asciz ""
.align 2
str76: .asciz "OK Declaración sin valor: correcto"
.align 2
str77: .asciz "%s"
.align 2
str78: .asciz "\n"
.align 2
str79: .asciz "X Declaración sin valor: incorrecto"
.align 2
str80: .asciz "%s"
.align 2
str81: .asciz "\n"
.align 2
str82: .asciz "Declaración con inferencia de tipo (convertida a explícita)"
.align 2
str83: .asciz "%s"
.align 2
str84: .asciz "\n"
enteroInferido_2: .quad 0
.align 3
F85: .double 2.71
decimalInferido_2: .double 0.0
.align 2
str86: .asciz "Adios!"
textoInferido_2: .quad 0
booleanoInferido_2: .quad 0
.align 2
str87: .asciz "enteroInferido:"
.align 2
str88: .asciz "%s"
.align 2
str89: .asciz " "
.align 2
str90: .asciz "%d"
.align 2
str91: .asciz "\n"
.align 2
str92: .asciz "decimalInferido:"
.align 2
str93: .asciz "%s"
.align 2
str94: .asciz " "
.align 2
str95: .asciz "%f"
.align 2
str96: .asciz "\n"
.align 2
str97: .asciz "textoInferido:"
.align 2
str98: .asciz "%s"
.align 2
str99: .asciz " "
.align 2
str100: .asciz "%s"
.align 2
str101: .asciz "\n"
.align 2
str102: .asciz "booleanoInferido:"
.align 2
str103: .asciz "%s"
.align 2
str104: .asciz " "
.align 2
str105: .asciz "%s"
.align 2
str106: .asciz "\n"
.align 3
F107: .double 2.7
.align 2
str108: .asciz "Adios!"
.align 2
str109: .asciz "OK Declaración con inferencia: correcto"
.align 2
str110: .asciz "%s"
.align 2
str111: .asciz "\n"
.align 2
str112: .asciz "X Declaración con inferencia: incorrecto"
.align 2
str113: .asciz "%s"
.align 2
str114: .asciz "\n"
.align 2
str115: .asciz "\n==== Asignación de variables ===="
.align 2
str116: .asciz "%s"
.align 2
str117: .asciz "\n"
puntosAsignacion_2: .quad 0
.align 2
str118: .asciz "Asignación con tipo correcto"
.align 2
str119: .asciz "%s"
.align 2
str120: .asciz "\n"
.align 3
F121: .double 9.9
.align 2
str122: .asciz "Nuevo"
.align 2
str123: .asciz "\n\n###Validacion Manual"
.align 2
str124: .asciz "%s"
.align 2
str125: .asciz "\n"
.align 2
str126: .asciz "entero:"
.align 2
str127: .asciz "%s"
.align 2
str128: .asciz " "
.align 2
str129: .asciz "%d"
.align 2
str130: .asciz "\n"
.align 2
str131: .asciz "decimal:"
.align 2
str132: .asciz "%s"
.align 2
str133: .asciz " "
.align 2
str134: .asciz "%f"
.align 2
str135: .asciz "\n"
.align 2
str136: .asciz "texto:"
.align 2
str137: .asciz "%s"
.align 2
str138: .asciz " "
.align 2
str139: .asciz "%s"
.align 2
str140: .asciz "\n"
.align 2
str141: .asciz "booleano:"
.align 2
str142: .asciz "%s"
.align 2
str143: .asciz " "
.align 2
str144: .asciz "%s"
.align 2
str145: .asciz "\n"
.align 2
str146: .asciz ""
.align 2
str147: .asciz "%s"
.align 2
str148: .asciz "\n"
.align 3
F149: .double 9.9
.align 2
str150: .asciz "Nuevo"
.align 2
str151: .asciz "OK Asignación simple: correcto"
.align 2
str152: .asciz "%s"
.align 2
str153: .asciz "\n"
.align 2
str154: .asciz "X Asignación simple: incorrecto"
.align 2
str155: .asciz "%s"
.align 2
str156: .asciz "\n"
.align 2
str157: .asciz "Asignación con expresiones"
.align 2
str158: .asciz "%s"
.align 2
str159: .asciz "\n"
.align 2
str160: .asciz "!"
.align 2
str161: .asciz "entero:"
.align 2
str162: .asciz "%s"
.align 2
str163: .asciz " "
.align 2
str164: .asciz "%d"
.align 2
str165: .asciz "\n"
.align 2
str166: .asciz "decimal:"
.align 2
str167: .asciz "%s"
.align 2
str168: .asciz " "
.align 2
str169: .asciz "%f"
.align 2
str170: .asciz "\n"
.align 2
str171: .asciz "texto:"
.align 2
str172: .asciz "%s"
.align 2
str173: .asciz " "
.align 2
str174: .asciz "%s"
.align 2
str175: .asciz "\n"
.align 2
str176: .asciz "booleano:"
.align 2
str177: .asciz "%s"
.align 2
str178: .asciz " "
.align 2
str179: .asciz "%s"
.align 2
str180: .asciz "\n"
.align 3
F181: .double 19.8
.align 2
str182: .asciz "Nuevo!"
.align 2
str183: .asciz "OK Asignación con expresiones: correcto"
.align 2
str184: .asciz "%s"
.align 2
str185: .asciz "\n"
.align 2
str186: .asciz "X Asignación con expresiones: incorrecto"
.align 2
str187: .asciz "%s"
.align 2
str188: .asciz "\n"
.align 2
str189: .asciz "Asignación con tipo incorrecto"
.align 2
str190: .asciz "%s"
.align 2
str191: .asciz "\n"
.align 2
str192: .asciz "OK Asignación con tipo incorrecto: correcto"
.align 2
str193: .asciz "%s"
.align 2
str194: .asciz "\n"
.align 2
str195: .asciz "\n==== Operaciones Aritméticas ===="
.align 2
str196: .asciz "%s"
.align 2
str197: .asciz "\n"
puntosOperacionesAritmeticas_2: .quad 0
.align 2
str198: .asciz "Suma"
.align 2
str199: .asciz "%s"
.align 2
str200: .asciz "\n"
resultadoSuma1_2: .quad 0
.align 3
F201: .double 10.5
.align 3
F202: .double 5.5
resultadoSuma2_2: .double 0.0
.align 3
F203: .double 5.5
resultadoSuma3_2: .double 0.0
.align 3
F204: .double 10.5
resultadoSuma4_2: .double 0.0
.align 2
str205: .asciz "10 + 5 ="
.align 2
str206: .asciz "%s"
.align 2
str207: .asciz " "
.align 2
str208: .asciz "%d"
.align 2
str209: .asciz "\n"
.align 2
str210: .asciz "10.5 + 5.5 ="
.align 2
str211: .asciz "%s"
.align 2
str212: .asciz " "
.align 2
str213: .asciz "%f"
.align 2
str214: .asciz "\n"
.align 2
str215: .asciz "10 + 5.5 ="
.align 2
str216: .asciz "%s"
.align 2
str217: .asciz " "
.align 2
str218: .asciz "%f"
.align 2
str219: .asciz "\n"
.align 2
str220: .asciz "10.5 + 5 ="
.align 2
str221: .asciz "%s"
.align 2
str222: .asciz " "
.align 2
str223: .asciz "%f"
.align 2
str224: .asciz "\n"
.align 3
F225: .double 16.0
.align 3
F226: .double 15.5
.align 3
F227: .double 15.5
.align 2
str228: .asciz "OK Suma: correcto"
.align 2
str229: .asciz "%s"
.align 2
str230: .asciz "\n"
.align 2
str231: .asciz "X Suma: incorrecto"
.align 2
str232: .asciz "%s"
.align 2
str233: .asciz "\n"
.align 2
str234: .asciz "Multiplicación"
.align 2
str235: .asciz "%s"
.align 2
str236: .asciz "\n"
resultadoMult1_2: .quad 0
.align 3
F237: .double 5.5
.align 3
F238: .double 2.0
resultadoMult2_2: .double 0.0
.align 3
F239: .double 2.5
resultadoMult3_2: .double 0.0
.align 3
F240: .double 5.5
resultadoMult4_2: .double 0.0
.align 2
str241: .asciz "5 * 3 ="
.align 2
str242: .asciz "%s"
.align 2
str243: .asciz " "
.align 2
str244: .asciz "%d"
.align 2
str245: .asciz "\n"
.align 2
str246: .asciz "5.5 * 2.0 ="
.align 2
str247: .asciz "%s"
.align 2
str248: .asciz " "
.align 2
str249: .asciz "%f"
.align 2
str250: .asciz "\n"
.align 2
str251: .asciz "5 * 2.5 ="
.align 2
str252: .asciz "%s"
.align 2
str253: .asciz " "
.align 2
str254: .asciz "%f"
.align 2
str255: .asciz "\n"
.align 2
str256: .asciz "5.5 * 2 ="
.align 2
str257: .asciz "%s"
.align 2
str258: .asciz " "
.align 2
str259: .asciz "%f"
.align 2
str260: .asciz "\n"
.align 3
F261: .double 11.0
.align 3
F262: .double 12.5
.align 3
F263: .double 11.0
.align 2
str264: .asciz "OK Multiplicación: correcto"
.align 2
str265: .asciz "%s"
.align 2
str266: .asciz "\n"
.align 2
str267: .asciz "X Multiplicación: incorrecto"
.align 2
str268: .asciz "%s"
.align 2
str269: .asciz "\n"
.align 2
str270: .asciz "División"
.align 2
str271: .asciz "%s"
.align 2
str272: .asciz "\n"
resultadoDiv1_2: .quad 0
.align 3
F273: .double 10.0
.align 3
F274: .double 4.0
resultadoDiv2_2: .double 0.0
.align 3
F275: .double 4.0
resultadoDiv3_2: .double 0.0
.align 3
F276: .double 10.0
resultadoDiv4_2: .double 0.0
.align 2
str277: .asciz "10 / 2 ="
.align 2
str278: .asciz "%s"
.align 2
str279: .asciz " "
.align 2
str280: .asciz "%d"
.align 2
str281: .asciz "\n"
.align 2
str282: .asciz "10.0 / 4.0 ="
.align 2
str283: .asciz "%s"
.align 2
str284: .asciz " "
.align 2
str285: .asciz "%f"
.align 2
str286: .asciz "\n"
.align 2
str287: .asciz "10 / 4.0 ="
.align 2
str288: .asciz "%s"
.align 2
str289: .asciz " "
.align 2
str290: .asciz "%f"
.align 2
str291: .asciz "\n"
.align 2
str292: .asciz "10.0 / 4 ="
.align 2
str293: .asciz "%s"
.align 2
str294: .asciz " "
.align 2
str295: .asciz "%f"
.align 2
str296: .asciz "\n"
.align 3
F297: .double 2.5
.align 3
F298: .double 2.5
.align 3
F299: .double 2.5
.align 2
str300: .asciz "OK División: correcto"
.align 2
str301: .asciz "%s"
.align 2
str302: .asciz "\n"
.align 2
str303: .asciz "X División: incorrecto"
.align 2
str304: .asciz "%s"
.align 2
str305: .asciz "\n"
.align 2
str306: .asciz "\n==== Operaciones Relacionales ===="
.align 2
str307: .asciz "%s"
.align 2
str308: .asciz "\n"
puntosOperacionesRelacionales_2: .quad 0
.align 2
str309: .asciz "Igualdad"
.align 2
str310: .asciz "%s"
.align 2
str311: .asciz "\n"
resultadoIgualdad1_2: .quad 0
resultadoIgualdad2_2: .quad 0
.align 3
F312: .double 10.5
.align 3
F313: .double 10.5
resultadoIgualdad3_2: .quad 0
.align 3
F314: .double 10.5
.align 3
F315: .double 5.5
resultadoIgualdad4_2: .quad 0
.align 2
str316: .asciz "Hola"
.align 2
str317: .asciz "Hola"
resultadoIgualdad5_2: .quad 0
.align 2
str318: .asciz "Hola"
.align 2
str319: .asciz "Mundo"
resultadoIgualdad6_2: .quad 0
.align 2
str320: .asciz "10 == 10:"
.align 2
str321: .asciz "%s"
.align 2
str322: .asciz " "
.align 2
str323: .asciz "%s"
.align 2
str324: .asciz "\n"
.align 2
str325: .asciz "10 == 5:"
.align 2
str326: .asciz "%s"
.align 2
str327: .asciz " "
.align 2
str328: .asciz "%s"
.align 2
str329: .asciz "\n"
.align 2
str330: .asciz "10.5 == 10.5:"
.align 2
str331: .asciz "%s"
.align 2
str332: .asciz " "
.align 2
str333: .asciz "%s"
.align 2
str334: .asciz "\n"
.align 2
str335: .asciz "10.5 == 5.5:"
.align 2
str336: .asciz "%s"
.align 2
str337: .asciz " "
.align 2
str338: .asciz "%s"
.align 2
str339: .asciz "\n"
.align 2
str340: .asciz "\"Hola\" == \"Hola\":"
.align 2
str341: .asciz "%s"
.align 2
str342: .asciz " "
.align 2
str343: .asciz "%s"
.align 2
str344: .asciz "\n"
.align 2
str345: .asciz "\"Hola\" == \"Mundo\":"
.align 2
str346: .asciz "%s"
.align 2
str347: .asciz " "
.align 2
str348: .asciz "%s"
.align 2
str349: .asciz "\n"
.align 2
str350: .asciz "OK Igualdad: correcto"
.align 2
str351: .asciz "%s"
.align 2
str352: .asciz "\n"
.align 2
str353: .asciz "X Igualdad: incorrecto"
.align 2
str354: .asciz "%s"
.align 2
str355: .asciz "\n"
.align 2
str356: .asciz "Mayor/Menor"
.align 2
str357: .asciz "%s"
.align 2
str358: .asciz "\n"
resultadoComp1_2: .quad 0
resultadoComp2_2: .quad 0
.align 3
F359: .double 10.5
.align 3
F360: .double 5.5
resultadoComp3_2: .quad 0
.align 3
F361: .double 10.5
.align 3
F362: .double 5.5
resultadoComp4_2: .quad 0
.align 2
str363: .asciz "10 > 5:"
.align 2
str364: .asciz "%s"
.align 2
str365: .asciz " "
.align 2
str366: .asciz "%s"
.align 2
str367: .asciz "\n"
.align 2
str368: .asciz "10 < 5:"
.align 2
str369: .asciz "%s"
.align 2
str370: .asciz " "
.align 2
str371: .asciz "%s"
.align 2
str372: .asciz "\n"
.align 2
str373: .asciz "10.5 > 5.5:"
.align 2
str374: .asciz "%s"
.align 2
str375: .asciz " "
.align 2
str376: .asciz "%s"
.align 2
str377: .asciz "\n"
.align 2
str378: .asciz "10.5 < 5.5:"
.align 2
str379: .asciz "%s"
.align 2
str380: .asciz " "
.align 2
str381: .asciz "%s"
.align 2
str382: .asciz "\n"
.align 2
str383: .asciz "OK Mayor/Menor: correcto"
.align 2
str384: .asciz "%s"
.align 2
str385: .asciz "\n"
.align 2
str386: .asciz "X Mayor/Menor: incorrecto"
.align 2
str387: .asciz "%s"
.align 2
str388: .asciz "\n"
.align 2
str389: .asciz "Mayor o igual/Menor o igual"
.align 2
str390: .asciz "%s"
.align 2
str391: .asciz "\n"
resultadoComp5_2: .quad 0
resultadoComp6_2: .quad 0
.align 3
F392: .double 10.5
.align 3
F393: .double 5.5
resultadoComp7_2: .quad 0
.align 3
F394: .double 10.5
.align 3
F395: .double 10.5
resultadoComp8_2: .quad 0
.align 2
str396: .asciz "10 >= 10:"
.align 2
str397: .asciz "%s"
.align 2
str398: .asciz " "
.align 2
str399: .asciz "%s"
.align 2
str400: .asciz "\n"
.align 2
str401: .asciz "10 <= 5:"
.align 2
str402: .asciz "%s"
.align 2
str403: .asciz " "
.align 2
str404: .asciz "%s"
.align 2
str405: .asciz "\n"
.align 2
str406: .asciz "10.5 >= 5.5:"
.align 2
str407: .asciz "%s"
.align 2
str408: .asciz " "
.align 2
str409: .asciz "%s"
.align 2
str410: .asciz "\n"
.align 2
str411: .asciz "10.5 <= 10.5:"
.align 2
str412: .asciz "%s"
.align 2
str413: .asciz " "
.align 2
str414: .asciz "%s"
.align 2
str415: .asciz "\n"
.align 2
str416: .asciz "OK Mayor o igual/Menor o igual: correcto"
.align 2
str417: .asciz "%s"
.align 2
str418: .asciz "\n"
.align 2
str419: .asciz "X Mayor o igual/Menor o igual: incorrecto"
.align 2
str420: .asciz "%s"
.align 2
str421: .asciz "\n"
.align 2
str422: .asciz "\n==== Operaciones Lógicas ===="
.align 2
str423: .asciz "%s"
.align 2
str424: .asciz "\n"
puntosOperacionesLogicas_2: .quad 0
.align 2
str425: .asciz "AND"
.align 2
str426: .asciz "%s"
.align 2
str427: .asciz "\n"
resultadoAnd1_2: .quad 0
resultadoAnd2_2: .quad 0
resultadoAnd3_2: .quad 0
resultadoAnd4_2: .quad 0
.align 2
str428: .asciz "true && true:"
.align 2
str429: .asciz "%s"
.align 2
str430: .asciz " "
.align 2
str431: .asciz "%s"
.align 2
str432: .asciz "\n"
.align 2
str433: .asciz "true && false:"
.align 2
str434: .asciz "%s"
.align 2
str435: .asciz " "
.align 2
str436: .asciz "%s"
.align 2
str437: .asciz "\n"
.align 2
str438: .asciz "(10 == 10) && (5 == 5):"
.align 2
str439: .asciz "%s"
.align 2
str440: .asciz " "
.align 2
str441: .asciz "%s"
.align 2
str442: .asciz "\n"
.align 2
str443: .asciz "(10 == 10) && (5 == 6):"
.align 2
str444: .asciz "%s"
.align 2
str445: .asciz " "
.align 2
str446: .asciz "%s"
.align 2
str447: .asciz "\n"
.align 2
str448: .asciz "OK AND: correcto"
.align 2
str449: .asciz "%s"
.align 2
str450: .asciz "\n"
.align 2
str451: .asciz "X AND: incorrecto"
.align 2
str452: .asciz "%s"
.align 2
str453: .asciz "\n"
.align 2
str454: .asciz "OR"
.align 2
str455: .asciz "%s"
.align 2
str456: .asciz "\n"
resultadoOr1_2: .quad 0
resultadoOr2_2: .quad 0
resultadoOr3_2: .quad 0
resultadoOr4_2: .quad 0
.align 2
str457: .asciz "true || false:"
.align 2
str458: .asciz "%s"
.align 2
str459: .asciz " "
.align 2
str460: .asciz "%s"
.align 2
str461: .asciz "\n"
.align 2
str462: .asciz "false || false:"
.align 2
str463: .asciz "%s"
.align 2
str464: .asciz " "
.align 2
str465: .asciz "%s"
.align 2
str466: .asciz "\n"
.align 2
str467: .asciz "(10 == 10) || (5 == 6):"
.align 2
str468: .asciz "%s"
.align 2
str469: .asciz " "
.align 2
str470: .asciz "%s"
.align 2
str471: .asciz "\n"
.align 2
str472: .asciz "(10 == 11) || (5 == 6):"
.align 2
str473: .asciz "%s"
.align 2
str474: .asciz " "
.align 2
str475: .asciz "%s"
.align 2
str476: .asciz "\n"
.align 2
str477: .asciz "OK OR: correcto"
.align 2
str478: .asciz "%s"
.align 2
str479: .asciz "\n"
.align 2
str480: .asciz "X OR: incorrecto"
.align 2
str481: .asciz "%s"
.align 2
str482: .asciz "\n"
.align 2
str483: .asciz "NOT"
.align 2
str484: .asciz "%s"
.align 2
str485: .asciz "\n"
resultadoNot1_2: .quad 0
resultadoNot2_2: .quad 0
resultadoNot3_2: .quad 0
resultadoNot4_2: .quad 0
.align 2
str486: .asciz "!true:"
.align 2
str487: .asciz "%s"
.align 2
str488: .asciz " "
.align 2
str489: .asciz "%s"
.align 2
str490: .asciz "\n"
.align 2
str491: .asciz "!false:"
.align 2
str492: .asciz "%s"
.align 2
str493: .asciz " "
.align 2
str494: .asciz "%s"
.align 2
str495: .asciz "\n"
.align 2
str496: .asciz "!(10 == 10):"
.align 2
str497: .asciz "%s"
.align 2
str498: .asciz " "
.align 2
str499: .asciz "%s"
.align 2
str500: .asciz "\n"
.align 2
str501: .asciz "!(10 == 11):"
.align 2
str502: .asciz "%s"
.align 2
str503: .asciz " "
.align 2
str504: .asciz "%s"
.align 2
str505: .asciz "\n"
.align 2
str506: .asciz "OK NOT: correcto"
.align 2
str507: .asciz "%s"
.align 2
str508: .asciz "\n"
.align 2
str509: .asciz "X NOT: incorrecto"
.align 2
str510: .asciz "%s"
.align 2
str511: .asciz "\n"
.align 2
str512: .asciz "\n==== print ===="
.align 2
str513: .asciz "%s"
.align 2
str514: .asciz "\n"
puntosPrintln_2: .quad 0
.align 2
str515: .asciz "\n\n###Validacion Manual"
.align 2
str516: .asciz "%s"
.align 2
str517: .asciz "\n"
.align 2
str518: .asciz "Impresión de valores simples"
.align 2
str519: .asciz "%s"
.align 2
str520: .asciz "\n"
.align 2
str521: .asciz "%d"
.align 2
str522: .asciz "\n"
.align 3
F523: .double 3.14
.align 2
str524: .asciz "%f"
.align 2
str525: .asciz "\n"
.align 2
str526: .asciz "Texto"
.align 2
str527: .asciz "%s"
.align 2
str528: .asciz "\n"
.align 2
str529: .asciz "%s"
.align 2
str530: .asciz "\n"
.align 2
str531: .asciz ""
.align 2
str532: .asciz "%s"
.align 2
str533: .asciz "\n"
.align 2
str534: .asciz "OK Impresión de valores simples: correcto"
.align 2
str535: .asciz "%s"
.align 2
str536: .asciz "\n"
.align 2
str537: .asciz "Impresión de múltiples valores"
.align 2
str538: .asciz "%s"
.align 2
str539: .asciz "\n"
.align 2
str540: .asciz "Números:"
.align 2
str541: .asciz "%s"
.align 2
str542: .asciz " "
.align 2
str543: .asciz "%d"
.align 3
F544: .double 3.14
.align 2
str545: .asciz "%f"
.align 2
str546: .asciz "\n"
.align 2
str547: .asciz "Booleano:"
.align 2
str548: .asciz "%s"
.align 2
str549: .asciz " "
.align 2
str550: .asciz "%s"
.align 2
str551: .asciz "Texto:"
.align 2
str552: .asciz "%s"
.align 2
str553: .asciz "Hola"
.align 2
str554: .asciz "%s"
.align 2
str555: .asciz "\n"
.align 2
str556: .asciz "OK Impresión de múltiples valores: correcto"
.align 2
str557: .asciz "%s"
.align 2
str558: .asciz "\n"
.align 2
str559: .asciz "Impresión de expresiones"
.align 2
str560: .asciz "%s"
.align 2
str561: .asciz "\n"
.align 2
str562: .asciz "Suma:"
.align 2
str563: .asciz "%s"
.align 2
str564: .asciz " "
.align 2
str565: .asciz "%d"
.align 2
str566: .asciz "\n"
.align 2
str567: .asciz "Comparación:"
.align 2
str568: .asciz "%s"
.align 2
str569: .asciz " "
.align 2
str570: .asciz "%s"
.align 2
str571: .asciz "\n"
.align 2
str572: .asciz "Lógica:"
.align 2
str573: .asciz "%s"
.align 2
str574: .asciz " "
.align 2
str575: .asciz "%s"
.align 2
str576: .asciz "\n"
.align 2
str577: .asciz "OK Impresión de expresiones: correcto"
.align 2
str578: .asciz "%s"
.align 2
str579: .asciz "\n"
.align 2
str580: .asciz "\n==== Manejo de valor nulo ===="
.align 2
str581: .asciz "%s"
.align 2
str582: .asciz "\n"
puntosValorNulo_2: .quad 0
.align 2
str583: .asciz "Valores por defecto"
.align 2
str584: .asciz "%s"
.align 2
str585: .asciz "\n"
enteroNulo_2: .quad 0
.align 3
decimalNulo_2: .double 0.0
textoNulo_2: .quad str53
booleanoNulo_2: .quad 0
.align 2
str586: .asciz "\n\n###Validacion Manual"
.align 2
str587: .asciz "%s"
.align 2
str588: .asciz "\n"
.align 2
str589: .asciz "enteroNulo:"
.align 2
str590: .asciz "%s"
.align 2
str591: .asciz " "
.align 2
str592: .asciz "%d"
.align 2
str593: .asciz "\n"
.align 2
str594: .asciz "decimalNulo:"
.align 2
str595: .asciz "%s"
.align 2
str596: .asciz " "
.align 2
str597: .asciz "%f"
.align 2
str598: .asciz "\n"
.align 2
str599: .asciz "textoNulo:"
.align 2
str600: .asciz "%s"
.align 2
str601: .asciz " "
.align 2
str602: .asciz "%s"
.align 2
str603: .asciz "\n"
.align 2
str604: .asciz "booleanoNulo:"
.align 2
str605: .asciz "%s"
.align 2
str606: .asciz " "
.align 2
str607: .asciz "%s"
.align 2
str608: .asciz "\n"
.align 2
str609: .asciz ""
.align 2
str610: .asciz "%s"
.align 2
str611: .asciz "\n"
.align 3
F612: .double 0.0
.align 2
str613: .asciz ""
.align 2
str614: .asciz "OK Valores por defecto: correcto"
.align 2
str615: .asciz "%s"
.align 2
str616: .asciz "\n"
.align 2
str617: .asciz "X Valores por defecto: incorrecto"
.align 2
str618: .asciz "%s"
.align 2
str619: .asciz "\n"
.align 2
str620: .asciz "Operaciones con nil"
.align 2
str621: .asciz "%s"
.align 2
str622: .asciz "\n"
.align 2
str623: .asciz "OK Operaciones con nil: correcto"
.align 2
str624: .asciz "%s"
.align 2
str625: .asciz "\n"
.align 2
str626: .asciz "\n=== Tabla de Resultados ==="
.align 2
str627: .asciz "%s"
.align 2
str628: .asciz "\n"
.align 2
str629: .asciz "+--------------------------+--------+-------+"
.align 2
str630: .asciz "%s"
.align 2
str631: .asciz "\n"
.align 2
str632: .asciz "| Característica           | Puntos | Total |"
.align 2
str633: .asciz "%s"
.align 2
str634: .asciz "\n"
.align 2
str635: .asciz "+--------------------------+--------+-------+"
.align 2
str636: .asciz "%s"
.align 2
str637: .asciz "\n"
.align 2
str638: .asciz "| Declaración de variables | "
.align 2
str639: .asciz "%s"
.align 2
str640: .asciz " "
.align 2
str641: .asciz "%d"
.align 2
str642: .asciz "    | 3     |"
.align 2
str643: .asciz "%s"
.align 2
str644: .asciz "\n"
.align 2
str645: .asciz "| Asignación de variables  | "
.align 2
str646: .asciz "%s"
.align 2
str647: .asciz " "
.align 2
str648: .asciz "%d"
.align 2
str649: .asciz "    | 3     |"
.align 2
str650: .asciz "%s"
.align 2
str651: .asciz "\n"
.align 2
str652: .asciz "| Operaciones Aritméticas  | "
.align 2
str653: .asciz "%s"
.align 2
str654: .asciz " "
.align 2
str655: .asciz "%d"
.align 2
str656: .asciz "    | 3     |"
.align 2
str657: .asciz "%s"
.align 2
str658: .asciz "\n"
.align 2
str659: .asciz "| Operaciones Relacionales | "
.align 2
str660: .asciz "%s"
.align 2
str661: .asciz " "
.align 2
str662: .asciz "%d"
.align 2
str663: .asciz "    | 3     |"
.align 2
str664: .asciz "%s"
.align 2
str665: .asciz "\n"
.align 2
str666: .asciz "| Operaciones Lógicas      | "
.align 2
str667: .asciz "%s"
.align 2
str668: .asciz " "
.align 2
str669: .asciz "%d"
.align 2
str670: .asciz "    | 3     |"
.align 2
str671: .asciz "%s"
.align 2
str672: .asciz "\n"
.align 2
str673: .asciz "| print                    | "
.align 2
str674: .asciz "%s"
.align 2
str675: .asciz " "
.align 2
str676: .asciz "%d"
.align 2
str677: .asciz "    | 3     |"
.align 2
str678: .asciz "%s"
.align 2
str679: .asciz "\n"
.align 2
str680: .asciz "| Manejo de valor nulo     | "
.align 2
str681: .asciz "%s"
.align 2
str682: .asciz " "
.align 2
str683: .asciz "%d"
.align 2
str684: .asciz "    | 2     |"
.align 2
str685: .asciz "%s"
.align 2
str686: .asciz "\n"
.align 2
str687: .asciz "+--------------------------+--------+-------+"
.align 2
str688: .asciz "%s"
.align 2
str689: .asciz "\n"
.align 2
str690: .asciz "| TOTAL                    | "
.align 2
str691: .asciz "%s"
.align 2
str692: .asciz " "
.align 2
str693: .asciz "%d"
.align 2
str694: .asciz "   | 20    |"
.align 2
str695: .asciz "%s"
.align 2
str696: .asciz "\n"
.align 2
str697: .asciz "+--------------------------+--------+-------+"
.align 2
str698: .asciz "%s"
.align 2
str699: .asciz "\n"

.extern malloc
.extern strlen
.extern strcpy
.extern strcat
.extern strcmp
.extern printf
.text
.global main
main:
    STP X29, X30, [SP, #-16]!
    MOV X29, SP
    MOV X9, #0
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =puntos_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str0
    LDR X0, =str1
    MOV X1, X9
    BL printf
    LDR X0, =str2
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str3
    LDR X0, =str4
    MOV X1, X9
    BL printf
    LDR X0, =str5
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str6
    LDR X0, =str7
    MOV X1, X9
    BL printf
    LDR X0, =str8
    BL printf
    // --- End of print call ---
    MOV X9, #0
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =puntosDeclaracion_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str9
    LDR X0, =str10
    MOV X1, X9
    BL printf
    LDR X0, =str11
    BL printf
    // --- End of print call ---
    MOV X9, #42
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =entero_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    LDR D8, F12
    SUB SP, SP, #16
    STR D8, [SP]
    LDR X9, =decimal_2
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    LDR X9, =str13
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =texto_2
    LDR X10, [SP]
    STR X10, [X9]
    ADD SP, SP, #16
    MOV X9, #1
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =booleano_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str14
    LDR X0, =str15
    MOV X1, X9
    BL printf
    LDR X0, =str16
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str17
    LDR X0, =str18
    MOV X1, X9
    BL printf
    LDR X0, =str19
    BL printf
    LDR X9, =entero_2
    LDRSW X10, [X9]
    LDR X0, =str20
    MOV X1, X10
    BL printf
    LDR X0, =str21
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str22
    LDR X0, =str23
    MOV X1, X9
    BL printf
    LDR X0, =str24
    BL printf
    LDR X9, =decimal_2
    LDR D8, [X9]
    LDR X0, =str25
    FMOV D0, D8
    BL printf
    LDR X0, =str26
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str27
    LDR X0, =str28
    MOV X1, X9
    BL printf
    LDR X0, =str29
    BL printf
    LDR X9, =texto_2
    LDR X10, [X9]
    LDR X0, =str30
    MOV X1, X10
    BL printf
    LDR X0, =str31
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str32
    LDR X0, =str33
    MOV X1, X9
    BL printf
    LDR X0, =str34
    BL printf
    LDR X9, =booleano_2
    LDRB W10, [X9]
    LDR X0, =str35
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str38
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str39
    LDR X0, =str40
    MOV X1, X9
    BL printf
    LDR X0, =str41
    BL printf
    // --- End of print call ---
    LDR X12, =entero_2
    LDRSW X13, [X12]
    MOV X12, #42
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false8
    LDR X12, =decimal_2
    LDR D8, [X12]
    LDR D9, F42
    FCMP D8, D9
    CSET W12, GT
    // Left was true, result is right operand
    MOV W11, W12
    B .L_logic_end7
.L_logic_false8:
    MOV W11, #0
.L_logic_end7:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false6
    LDR X11, =texto_2
    LDR X12, [X11]
    LDR X11, =str43
    MOV X0, X12
    MOV X1, X11
    BL strcmp
    CMP W0, #0
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W10, W13
    B .L_logic_end5
.L_logic_false6:
    MOV W10, #0
.L_logic_end5:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false4
    LDR X10, =booleano_2
    LDRB W11, [X10]
    MOV X10, #1
    CMP W11, W10
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W9, W11
    B .L_logic_end3
.L_logic_false4:
    MOV W9, #0
.L_logic_end3:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse1
    // 'Then' block
    LDR X9, =puntosDeclaracion_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosDeclaracion
    LDR X9, =puntosDeclaracion_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str44
    LDR X0, =str45
    MOV X1, X9
    BL printf
    LDR X0, =str46
    BL printf
    // --- End of print call ---
    B .Lendif2
.Lelse1:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str47
    LDR X0, =str48
    MOV X1, X9
    BL printf
    LDR X0, =str49
    BL printf
    // --- End of print call ---
.Lendif2:
    // --- Start of print call ---
    LDR X9, =str50
    LDR X0, =str51
    MOV X1, X9
    BL printf
    LDR X0, =str52
    BL printf
    // --- End of print call ---
    // Declaring enteroSinValor without initializer
    // Declaring decimalSinValor without initializer
    // Declaring textoSinValor without initializer
    // Declaring booleanoSinValor without initializer
    // --- Start of print call ---
    LDR X9, =str54
    LDR X0, =str55
    MOV X1, X9
    BL printf
    LDR X0, =str56
    BL printf
    LDR X9, =enteroSinValor_2
    LDRSW X10, [X9]
    LDR X0, =str57
    MOV X1, X10
    BL printf
    LDR X0, =str58
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str59
    LDR X0, =str60
    MOV X1, X9
    BL printf
    LDR X0, =str61
    BL printf
    LDR X9, =decimalSinValor_2
    LDR D8, [X9]
    LDR X0, =str62
    FMOV D0, D8
    BL printf
    LDR X0, =str63
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str64
    LDR X0, =str65
    MOV X1, X9
    BL printf
    LDR X0, =str66
    BL printf
    LDR X9, =textoSinValor_2
    LDR X10, [X9]
    LDR X0, =str67
    MOV X1, X10
    BL printf
    LDR X0, =str68
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str69
    LDR X0, =str70
    MOV X1, X9
    BL printf
    LDR X0, =str71
    BL printf
    LDR X9, =booleanoSinValor_2
    LDRB W10, [X9]
    LDR X0, =str72
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str73
    BL printf
    // --- End of print call ---
    LDR X12, =enteroSinValor_2
    LDRSW X13, [X12]
    MOV X12, #0
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false16
    LDR X12, =decimalSinValor_2
    LDR D8, [X12]
    LDR D9, F74
    FCMP D8, D9
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W11, W12
    B .L_logic_end15
.L_logic_false16:
    MOV W11, #0
.L_logic_end15:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false14
    LDR X11, =textoSinValor_2
    LDR X12, [X11]
    LDR X11, =str75
    MOV X0, X12
    MOV X1, X11
    BL strcmp
    CMP W0, #0
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W10, W13
    B .L_logic_end13
.L_logic_false14:
    MOV W10, #0
.L_logic_end13:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false12
    LDR X10, =booleanoSinValor_2
    LDRB W11, [X10]
    MOV X10, #0
    CMP W11, W10
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W9, W11
    B .L_logic_end11
.L_logic_false12:
    MOV W9, #0
.L_logic_end11:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse9
    // 'Then' block
    LDR X9, =puntosDeclaracion_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosDeclaracion
    LDR X9, =puntosDeclaracion_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str76
    LDR X0, =str77
    MOV X1, X9
    BL printf
    LDR X0, =str78
    BL printf
    // --- End of print call ---
    B .Lendif10
.Lelse9:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str79
    LDR X0, =str80
    MOV X1, X9
    BL printf
    LDR X0, =str81
    BL printf
    // --- End of print call ---
.Lendif10:
    // --- Start of print call ---
    LDR X9, =str82
    LDR X0, =str83
    MOV X1, X9
    BL printf
    LDR X0, =str84
    BL printf
    // --- End of print call ---
    MOV X9, #100
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =enteroInferido_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    LDR D8, F85
    SUB SP, SP, #16
    STR D8, [SP]
    LDR X9, =decimalInferido_2
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    LDR X9, =str86
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =textoInferido_2
    LDR X10, [SP]
    STR X10, [X9]
    ADD SP, SP, #16
    MOV X9, #0
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =booleanoInferido_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str87
    LDR X0, =str88
    MOV X1, X9
    BL printf
    LDR X0, =str89
    BL printf
    LDR X9, =enteroInferido_2
    LDRSW X10, [X9]
    LDR X0, =str90
    MOV X1, X10
    BL printf
    LDR X0, =str91
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str92
    LDR X0, =str93
    MOV X1, X9
    BL printf
    LDR X0, =str94
    BL printf
    LDR X9, =decimalInferido_2
    LDR D8, [X9]
    LDR X0, =str95
    FMOV D0, D8
    BL printf
    LDR X0, =str96
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str97
    LDR X0, =str98
    MOV X1, X9
    BL printf
    LDR X0, =str99
    BL printf
    LDR X9, =textoInferido_2
    LDR X10, [X9]
    LDR X0, =str100
    MOV X1, X10
    BL printf
    LDR X0, =str101
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str102
    LDR X0, =str103
    MOV X1, X9
    BL printf
    LDR X0, =str104
    BL printf
    LDR X9, =booleanoInferido_2
    LDRB W10, [X9]
    LDR X0, =str105
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str106
    BL printf
    // --- End of print call ---
    LDR X12, =enteroInferido_2
    LDRSW X13, [X12]
    MOV X12, #100
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false24
    LDR X12, =decimalInferido_2
    LDR D8, [X12]
    LDR D9, F107
    FCMP D8, D9
    CSET W12, GT
    // Left was true, result is right operand
    MOV W11, W12
    B .L_logic_end23
.L_logic_false24:
    MOV W11, #0
.L_logic_end23:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false22
    LDR X11, =textoInferido_2
    LDR X12, [X11]
    LDR X11, =str108
    MOV X0, X12
    MOV X1, X11
    BL strcmp
    CMP W0, #0
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W10, W13
    B .L_logic_end21
.L_logic_false22:
    MOV W10, #0
.L_logic_end21:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false20
    LDR X10, =booleanoInferido_2
    LDRB W11, [X10]
    MOV X10, #0
    CMP W11, W10
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W9, W11
    B .L_logic_end19
.L_logic_false20:
    MOV W9, #0
.L_logic_end19:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse17
    // 'Then' block
    LDR X9, =puntosDeclaracion_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosDeclaracion
    LDR X9, =puntosDeclaracion_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str109
    LDR X0, =str110
    MOV X1, X9
    BL printf
    LDR X0, =str111
    BL printf
    // --- End of print call ---
    B .Lendif18
.Lelse17:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str112
    LDR X0, =str113
    MOV X1, X9
    BL printf
    LDR X0, =str114
    BL printf
    // --- End of print call ---
.Lendif18:
    // --- Start of print call ---
    LDR X9, =str115
    LDR X0, =str116
    MOV X1, X9
    BL printf
    LDR X0, =str117
    BL printf
    // --- End of print call ---
    MOV X9, #0
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =puntosAsignacion_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str118
    LDR X0, =str119
    MOV X1, X9
    BL printf
    LDR X0, =str120
    BL printf
    // --- End of print call ---
    MOV X9, #99
    // Storing value for assignment to entero
    LDR X10, =entero_2
    SUB SP, SP, #16
    STR X9, [SP]
    LDR W9, [SP]
    STR W9, [X10]
    ADD SP, SP, #16
    LDR D8, F121
    // Storing value for assignment to decimal
    LDR X9, =decimal_2
    SUB SP, SP, #16
    STR D8, [SP]
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    LDR X9, =str122
    // Storing value for assignment to texto
    LDR X10, =texto_2
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, [SP]
    STR X9, [X10]
    ADD SP, SP, #16
    LDR X9, =booleano_2
    LDRB W10, [X9]
    EOR W10, W10, #1
    // Storing value for assignment to booleano
    LDR X9, =booleano_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str123
    LDR X0, =str124
    MOV X1, X9
    BL printf
    LDR X0, =str125
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str126
    LDR X0, =str127
    MOV X1, X9
    BL printf
    LDR X0, =str128
    BL printf
    LDR X9, =entero_2
    LDRSW X10, [X9]
    LDR X0, =str129
    MOV X1, X10
    BL printf
    LDR X0, =str130
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str131
    LDR X0, =str132
    MOV X1, X9
    BL printf
    LDR X0, =str133
    BL printf
    LDR X9, =decimal_2
    LDR D8, [X9]
    LDR X0, =str134
    FMOV D0, D8
    BL printf
    LDR X0, =str135
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str136
    LDR X0, =str137
    MOV X1, X9
    BL printf
    LDR X0, =str138
    BL printf
    LDR X9, =texto_2
    LDR X10, [X9]
    LDR X0, =str139
    MOV X1, X10
    BL printf
    LDR X0, =str140
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str141
    LDR X0, =str142
    MOV X1, X9
    BL printf
    LDR X0, =str143
    BL printf
    LDR X9, =booleano_2
    LDRB W10, [X9]
    LDR X0, =str144
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str145
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str146
    LDR X0, =str147
    MOV X1, X9
    BL printf
    LDR X0, =str148
    BL printf
    // --- End of print call ---
    LDR X12, =entero_2
    LDRSW X13, [X12]
    MOV X12, #99
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false32
    LDR X12, =decimal_2
    LDR D8, [X12]
    LDR D9, F149
    FCMP D8, D9
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W11, W12
    B .L_logic_end31
.L_logic_false32:
    MOV W11, #0
.L_logic_end31:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false30
    LDR X11, =texto_2
    LDR X12, [X11]
    LDR X11, =str150
    MOV X0, X12
    MOV X1, X11
    BL strcmp
    CMP W0, #0
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W10, W13
    B .L_logic_end29
.L_logic_false30:
    MOV W10, #0
.L_logic_end29:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false28
    LDR X10, =booleano_2
    LDRB W11, [X10]
    MOV X10, #0
    CMP W11, W10
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W9, W11
    B .L_logic_end27
.L_logic_false28:
    MOV W9, #0
.L_logic_end27:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse25
    // 'Then' block
    LDR X9, =puntosAsignacion_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosAsignacion
    LDR X9, =puntosAsignacion_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str151
    LDR X0, =str152
    MOV X1, X9
    BL printf
    LDR X0, =str153
    BL printf
    // --- End of print call ---
    B .Lendif26
.Lelse25:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str154
    LDR X0, =str155
    MOV X1, X9
    BL printf
    LDR X0, =str156
    BL printf
    // --- End of print call ---
.Lendif26:
    // --- Start of print call ---
    LDR X9, =str157
    LDR X0, =str158
    MOV X1, X9
    BL printf
    LDR X0, =str159
    BL printf
    // --- End of print call ---
    LDR X9, =entero_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to entero
    LDR X9, =entero_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    LDR X9, =decimal_2
    LDR D8, [X9]
    MOV X9, #2
    // Promoting right operand from INT to FLOAT
    SCVTF D9, W9
    FMUL D8, D8, D9
    // Storing value for assignment to decimal
    LDR X9, =decimal_2
    SUB SP, SP, #16
    STR D8, [SP]
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    LDR X9, =texto_2
    LDR X10, [X9]
    LDR X9, =str160
    // --- String Concatenation ---
    SUB SP, SP, #16
    STP X10, X9, [SP]
    LDR X0, [SP, #0]
    BL strlen
    MOV X9, X0
    LDR X0, [SP, #8]
    BL strlen
    MOV X10, X0
    ADD X0, X9, X10
    ADD X0, X0, #1
    BL malloc
    MOV X9, X0
    LDR X1, [SP, #0]
    MOV X0, X9
    BL strcpy
    LDR X1, [SP, #8]
    MOV X0, X9
    BL strcat
    ADD SP, SP, #16
    // Storing value for assignment to texto
    LDR X10, =texto_2
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, [SP]
    STR X9, [X10]
    ADD SP, SP, #16
    LDR X9, =booleano_2
    LDRB W10, [X9]
    EOR W10, W10, #1
    // Storing value for assignment to booleano
    LDR X9, =booleano_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str161
    LDR X0, =str162
    MOV X1, X9
    BL printf
    LDR X0, =str163
    BL printf
    LDR X9, =entero_2
    LDRSW X10, [X9]
    LDR X0, =str164
    MOV X1, X10
    BL printf
    LDR X0, =str165
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str166
    LDR X0, =str167
    MOV X1, X9
    BL printf
    LDR X0, =str168
    BL printf
    LDR X9, =decimal_2
    LDR D8, [X9]
    LDR X0, =str169
    FMOV D0, D8
    BL printf
    LDR X0, =str170
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str171
    LDR X0, =str172
    MOV X1, X9
    BL printf
    LDR X0, =str173
    BL printf
    LDR X9, =texto_2
    LDR X10, [X9]
    LDR X0, =str174
    MOV X1, X10
    BL printf
    LDR X0, =str175
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str176
    LDR X0, =str177
    MOV X1, X9
    BL printf
    LDR X0, =str178
    BL printf
    LDR X9, =booleano_2
    LDRB W10, [X9]
    LDR X0, =str179
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str180
    BL printf
    // --- End of print call ---
    LDR X12, =entero_2
    LDRSW X13, [X12]
    MOV X12, #100
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false40
    LDR X12, =decimal_2
    LDR D8, [X12]
    LDR D9, F181
    FCMP D8, D9
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W11, W12
    B .L_logic_end39
.L_logic_false40:
    MOV W11, #0
.L_logic_end39:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false38
    LDR X11, =texto_2
    LDR X12, [X11]
    LDR X11, =str182
    MOV X0, X12
    MOV X1, X11
    BL strcmp
    CMP W0, #0
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W10, W13
    B .L_logic_end37
.L_logic_false38:
    MOV W10, #0
.L_logic_end37:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false36
    LDR X10, =booleano_2
    LDRB W11, [X10]
    MOV X10, #1
    CMP W11, W10
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W9, W11
    B .L_logic_end35
.L_logic_false36:
    MOV W9, #0
.L_logic_end35:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse33
    // 'Then' block
    LDR X9, =puntosAsignacion_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosAsignacion
    LDR X9, =puntosAsignacion_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str183
    LDR X0, =str184
    MOV X1, X9
    BL printf
    LDR X0, =str185
    BL printf
    // --- End of print call ---
    B .Lendif34
.Lelse33:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str186
    LDR X0, =str187
    MOV X1, X9
    BL printf
    LDR X0, =str188
    BL printf
    // --- End of print call ---
.Lendif34:
    // --- Start of print call ---
    LDR X9, =str189
    LDR X0, =str190
    MOV X1, X9
    BL printf
    LDR X0, =str191
    BL printf
    // --- End of print call ---
    LDR X9, =puntosAsignacion_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosAsignacion
    LDR X9, =puntosAsignacion_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str192
    LDR X0, =str193
    MOV X1, X9
    BL printf
    LDR X0, =str194
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str195
    LDR X0, =str196
    MOV X1, X9
    BL printf
    LDR X0, =str197
    BL printf
    // --- End of print call ---
    MOV X9, #0
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =puntosOperacionesAritmeticas_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str198
    LDR X0, =str199
    MOV X1, X9
    BL printf
    LDR X0, =str200
    BL printf
    // --- End of print call ---
    MOV X9, #10
    MOV X10, #5
    ADD X9, X9, X10
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoSuma1_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    LDR D8, F201
    LDR D9, F202
    FADD D8, D8, D9
    SUB SP, SP, #16
    STR D8, [SP]
    LDR X9, =resultadoSuma2_2
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    MOV X9, #10
    LDR D8, F203
    // Promoting left operand from INT to FLOAT
    SCVTF D9, W9
    FADD D9, D9, D8
    SUB SP, SP, #16
    STR D9, [SP]
    LDR X9, =resultadoSuma3_2
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    LDR D8, F204
    MOV X9, #5
    // Promoting right operand from INT to FLOAT
    SCVTF D9, W9
    FADD D8, D8, D9
    SUB SP, SP, #16
    STR D8, [SP]
    LDR X9, =resultadoSuma4_2
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str205
    LDR X0, =str206
    MOV X1, X9
    BL printf
    LDR X0, =str207
    BL printf
    LDR X9, =resultadoSuma1_2
    LDRSW X10, [X9]
    LDR X0, =str208
    MOV X1, X10
    BL printf
    LDR X0, =str209
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str210
    LDR X0, =str211
    MOV X1, X9
    BL printf
    LDR X0, =str212
    BL printf
    LDR X9, =resultadoSuma2_2
    LDR D8, [X9]
    LDR X0, =str213
    FMOV D0, D8
    BL printf
    LDR X0, =str214
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str215
    LDR X0, =str216
    MOV X1, X9
    BL printf
    LDR X0, =str217
    BL printf
    LDR X9, =resultadoSuma3_2
    LDR D8, [X9]
    LDR X0, =str218
    FMOV D0, D8
    BL printf
    LDR X0, =str219
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str220
    LDR X0, =str221
    MOV X1, X9
    BL printf
    LDR X0, =str222
    BL printf
    LDR X9, =resultadoSuma4_2
    LDR D8, [X9]
    LDR X0, =str223
    FMOV D0, D8
    BL printf
    LDR X0, =str224
    BL printf
    // --- End of print call ---
    LDR X12, =resultadoSuma1_2
    LDRSW X13, [X12]
    MOV X12, #15
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false48
    LDR X12, =resultadoSuma2_2
    LDR D8, [X12]
    LDR D9, F225
    FCMP D8, D9
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W11, W12
    B .L_logic_end47
.L_logic_false48:
    MOV W11, #0
.L_logic_end47:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false46
    LDR X11, =resultadoSuma3_2
    LDR D8, [X11]
    LDR D9, F226
    FCMP D8, D9
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W10, W11
    B .L_logic_end45
.L_logic_false46:
    MOV W10, #0
.L_logic_end45:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false44
    LDR X10, =resultadoSuma4_2
    LDR D8, [X10]
    LDR D9, F227
    FCMP D8, D9
    CSET W10, EQ
    // Left was true, result is right operand
    MOV W9, W10
    B .L_logic_end43
.L_logic_false44:
    MOV W9, #0
.L_logic_end43:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse41
    // 'Then' block
    LDR X9, =puntosOperacionesAritmeticas_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosOperacionesAritmeticas
    LDR X9, =puntosOperacionesAritmeticas_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str228
    LDR X0, =str229
    MOV X1, X9
    BL printf
    LDR X0, =str230
    BL printf
    // --- End of print call ---
    B .Lendif42
.Lelse41:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str231
    LDR X0, =str232
    MOV X1, X9
    BL printf
    LDR X0, =str233
    BL printf
    // --- End of print call ---
.Lendif42:
    // --- Start of print call ---
    LDR X9, =str234
    LDR X0, =str235
    MOV X1, X9
    BL printf
    LDR X0, =str236
    BL printf
    // --- End of print call ---
    MOV X9, #5
    MOV X10, #3
    MUL X9, X9, X10
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoMult1_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    LDR D8, F237
    LDR D9, F238
    FMUL D8, D8, D9
    SUB SP, SP, #16
    STR D8, [SP]
    LDR X9, =resultadoMult2_2
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    MOV X9, #5
    LDR D8, F239
    // Promoting left operand from INT to FLOAT
    SCVTF D9, W9
    FMUL D9, D9, D8
    SUB SP, SP, #16
    STR D9, [SP]
    LDR X9, =resultadoMult3_2
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    LDR D8, F240
    MOV X9, #2
    // Promoting right operand from INT to FLOAT
    SCVTF D9, W9
    FMUL D8, D8, D9
    SUB SP, SP, #16
    STR D8, [SP]
    LDR X9, =resultadoMult4_2
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str241
    LDR X0, =str242
    MOV X1, X9
    BL printf
    LDR X0, =str243
    BL printf
    LDR X9, =resultadoMult1_2
    LDRSW X10, [X9]
    LDR X0, =str244
    MOV X1, X10
    BL printf
    LDR X0, =str245
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str246
    LDR X0, =str247
    MOV X1, X9
    BL printf
    LDR X0, =str248
    BL printf
    LDR X9, =resultadoMult2_2
    LDR D8, [X9]
    LDR X0, =str249
    FMOV D0, D8
    BL printf
    LDR X0, =str250
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str251
    LDR X0, =str252
    MOV X1, X9
    BL printf
    LDR X0, =str253
    BL printf
    LDR X9, =resultadoMult3_2
    LDR D8, [X9]
    LDR X0, =str254
    FMOV D0, D8
    BL printf
    LDR X0, =str255
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str256
    LDR X0, =str257
    MOV X1, X9
    BL printf
    LDR X0, =str258
    BL printf
    LDR X9, =resultadoMult4_2
    LDR D8, [X9]
    LDR X0, =str259
    FMOV D0, D8
    BL printf
    LDR X0, =str260
    BL printf
    // --- End of print call ---
    LDR X12, =resultadoMult1_2
    LDRSW X13, [X12]
    MOV X12, #15
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false56
    LDR X12, =resultadoMult2_2
    LDR D8, [X12]
    LDR D9, F261
    FCMP D8, D9
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W11, W12
    B .L_logic_end55
.L_logic_false56:
    MOV W11, #0
.L_logic_end55:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false54
    LDR X11, =resultadoMult3_2
    LDR D8, [X11]
    LDR D9, F262
    FCMP D8, D9
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W10, W11
    B .L_logic_end53
.L_logic_false54:
    MOV W10, #0
.L_logic_end53:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false52
    LDR X10, =resultadoMult4_2
    LDR D8, [X10]
    LDR D9, F263
    FCMP D8, D9
    CSET W10, EQ
    // Left was true, result is right operand
    MOV W9, W10
    B .L_logic_end51
.L_logic_false52:
    MOV W9, #0
.L_logic_end51:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse49
    // 'Then' block
    LDR X9, =puntosOperacionesAritmeticas_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosOperacionesAritmeticas
    LDR X9, =puntosOperacionesAritmeticas_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str264
    LDR X0, =str265
    MOV X1, X9
    BL printf
    LDR X0, =str266
    BL printf
    // --- End of print call ---
    B .Lendif50
.Lelse49:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str267
    LDR X0, =str268
    MOV X1, X9
    BL printf
    LDR X0, =str269
    BL printf
    // --- End of print call ---
.Lendif50:
    // --- Start of print call ---
    LDR X9, =str270
    LDR X0, =str271
    MOV X1, X9
    BL printf
    LDR X0, =str272
    BL printf
    // --- End of print call ---
    MOV X9, #10
    MOV X10, #2
    SDIV X9, X9, X10
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoDiv1_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    LDR D8, F273
    LDR D9, F274
    FDIV D8, D8, D9
    SUB SP, SP, #16
    STR D8, [SP]
    LDR X9, =resultadoDiv2_2
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    MOV X9, #10
    LDR D8, F275
    // Promoting left operand from INT to FLOAT
    SCVTF D9, W9
    FDIV D9, D9, D8
    SUB SP, SP, #16
    STR D9, [SP]
    LDR X9, =resultadoDiv3_2
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    LDR D8, F276
    MOV X9, #4
    // Promoting right operand from INT to FLOAT
    SCVTF D9, W9
    FDIV D8, D8, D9
    SUB SP, SP, #16
    STR D8, [SP]
    LDR X9, =resultadoDiv4_2
    LDR D8, [SP]
    STR D8, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str277
    LDR X0, =str278
    MOV X1, X9
    BL printf
    LDR X0, =str279
    BL printf
    LDR X9, =resultadoDiv1_2
    LDRSW X10, [X9]
    LDR X0, =str280
    MOV X1, X10
    BL printf
    LDR X0, =str281
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str282
    LDR X0, =str283
    MOV X1, X9
    BL printf
    LDR X0, =str284
    BL printf
    LDR X9, =resultadoDiv2_2
    LDR D8, [X9]
    LDR X0, =str285
    FMOV D0, D8
    BL printf
    LDR X0, =str286
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str287
    LDR X0, =str288
    MOV X1, X9
    BL printf
    LDR X0, =str289
    BL printf
    LDR X9, =resultadoDiv3_2
    LDR D8, [X9]
    LDR X0, =str290
    FMOV D0, D8
    BL printf
    LDR X0, =str291
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str292
    LDR X0, =str293
    MOV X1, X9
    BL printf
    LDR X0, =str294
    BL printf
    LDR X9, =resultadoDiv4_2
    LDR D8, [X9]
    LDR X0, =str295
    FMOV D0, D8
    BL printf
    LDR X0, =str296
    BL printf
    // --- End of print call ---
    LDR X12, =resultadoDiv1_2
    LDRSW X13, [X12]
    MOV X12, #5
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false64
    LDR X12, =resultadoDiv2_2
    LDR D8, [X12]
    LDR D9, F297
    FCMP D8, D9
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W11, W12
    B .L_logic_end63
.L_logic_false64:
    MOV W11, #0
.L_logic_end63:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false62
    LDR X11, =resultadoDiv3_2
    LDR D8, [X11]
    LDR D9, F298
    FCMP D8, D9
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W10, W11
    B .L_logic_end61
.L_logic_false62:
    MOV W10, #0
.L_logic_end61:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false60
    LDR X10, =resultadoDiv4_2
    LDR D8, [X10]
    LDR D9, F299
    FCMP D8, D9
    CSET W10, EQ
    // Left was true, result is right operand
    MOV W9, W10
    B .L_logic_end59
.L_logic_false60:
    MOV W9, #0
.L_logic_end59:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse57
    // 'Then' block
    LDR X9, =puntosOperacionesAritmeticas_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosOperacionesAritmeticas
    LDR X9, =puntosOperacionesAritmeticas_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str300
    LDR X0, =str301
    MOV X1, X9
    BL printf
    LDR X0, =str302
    BL printf
    // --- End of print call ---
    B .Lendif58
.Lelse57:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str303
    LDR X0, =str304
    MOV X1, X9
    BL printf
    LDR X0, =str305
    BL printf
    // --- End of print call ---
.Lendif58:
    // --- Start of print call ---
    LDR X9, =str306
    LDR X0, =str307
    MOV X1, X9
    BL printf
    LDR X0, =str308
    BL printf
    // --- End of print call ---
    MOV X9, #0
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =puntosOperacionesRelacionales_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str309
    LDR X0, =str310
    MOV X1, X9
    BL printf
    LDR X0, =str311
    BL printf
    // --- End of print call ---
    MOV X9, #10
    MOV X10, #10
    CMP W9, W10
    CSET W9, EQ
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoIgualdad1_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    MOV X9, #10
    MOV X10, #5
    CMP W9, W10
    CSET W9, EQ
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoIgualdad2_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    LDR D8, F312
    LDR D9, F313
    FCMP D8, D9
    CSET W9, EQ
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoIgualdad3_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    LDR D8, F314
    LDR D9, F315
    FCMP D8, D9
    CSET W9, EQ
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoIgualdad4_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    LDR X9, =str316
    LDR X10, =str317
    MOV X0, X9
    MOV X1, X10
    BL strcmp
    CMP W0, #0
    CSET W11, EQ
    SUB SP, SP, #16
    STR X11, [SP]
    LDR X9, =resultadoIgualdad5_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    LDR X9, =str318
    LDR X10, =str319
    MOV X0, X9
    MOV X1, X10
    BL strcmp
    CMP W0, #0
    CSET W11, EQ
    SUB SP, SP, #16
    STR X11, [SP]
    LDR X9, =resultadoIgualdad6_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str320
    LDR X0, =str321
    MOV X1, X9
    BL printf
    LDR X0, =str322
    BL printf
    LDR X9, =resultadoIgualdad1_2
    LDRB W10, [X9]
    LDR X0, =str323
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str324
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str325
    LDR X0, =str326
    MOV X1, X9
    BL printf
    LDR X0, =str327
    BL printf
    LDR X9, =resultadoIgualdad2_2
    LDRB W10, [X9]
    LDR X0, =str328
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str329
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str330
    LDR X0, =str331
    MOV X1, X9
    BL printf
    LDR X0, =str332
    BL printf
    LDR X9, =resultadoIgualdad3_2
    LDRB W10, [X9]
    LDR X0, =str333
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str334
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str335
    LDR X0, =str336
    MOV X1, X9
    BL printf
    LDR X0, =str337
    BL printf
    LDR X9, =resultadoIgualdad4_2
    LDRB W10, [X9]
    LDR X0, =str338
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str339
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str340
    LDR X0, =str341
    MOV X1, X9
    BL printf
    LDR X0, =str342
    BL printf
    LDR X9, =resultadoIgualdad5_2
    LDRB W10, [X9]
    LDR X0, =str343
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str344
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str345
    LDR X0, =str346
    MOV X1, X9
    BL printf
    LDR X0, =str347
    BL printf
    LDR X9, =resultadoIgualdad6_2
    LDRB W10, [X9]
    LDR X0, =str348
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str349
    BL printf
    // --- End of print call ---
    LDR X14, =resultadoIgualdad1_2
    LDRB W15, [X14]
    MOV X14, #1
    CMP W15, W14
    CSET W15, EQ
    // Short-circuit AND: check left operand
    CMP W15, #0
    B.EQ .L_logic_false76
    LDR X14, =resultadoIgualdad2_2
    LDRB W15, [X14]
    MOV X14, #0
    CMP W15, W14
    CSET W15, EQ
    // Left was true, result is right operand
    MOV W13, W15
    B .L_logic_end75
.L_logic_false76:
    MOV W13, #0
.L_logic_end75:
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false74
    LDR X13, =resultadoIgualdad3_2
    LDRB W14, [X13]
    MOV X13, #1
    CMP W14, W13
    CSET W14, EQ
    // Left was true, result is right operand
    MOV W12, W14
    B .L_logic_end73
.L_logic_false74:
    MOV W12, #0
.L_logic_end73:
    // Short-circuit AND: check left operand
    CMP W12, #0
    B.EQ .L_logic_false72
    LDR X12, =resultadoIgualdad4_2
    LDRB W13, [X12]
    MOV X12, #0
    CMP W13, W12
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W11, W13
    B .L_logic_end71
.L_logic_false72:
    MOV W11, #0
.L_logic_end71:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false70
    LDR X11, =resultadoIgualdad5_2
    LDRB W12, [X11]
    MOV X11, #1
    CMP W12, W11
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W10, W12
    B .L_logic_end69
.L_logic_false70:
    MOV W10, #0
.L_logic_end69:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false68
    LDR X10, =resultadoIgualdad6_2
    LDRB W11, [X10]
    MOV X10, #0
    CMP W11, W10
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W9, W11
    B .L_logic_end67
.L_logic_false68:
    MOV W9, #0
.L_logic_end67:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse65
    // 'Then' block
    LDR X9, =puntosOperacionesRelacionales_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosOperacionesRelacionales
    LDR X9, =puntosOperacionesRelacionales_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str350
    LDR X0, =str351
    MOV X1, X9
    BL printf
    LDR X0, =str352
    BL printf
    // --- End of print call ---
    B .Lendif66
.Lelse65:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str353
    LDR X0, =str354
    MOV X1, X9
    BL printf
    LDR X0, =str355
    BL printf
    // --- End of print call ---
.Lendif66:
    // --- Start of print call ---
    LDR X9, =str356
    LDR X0, =str357
    MOV X1, X9
    BL printf
    LDR X0, =str358
    BL printf
    // --- End of print call ---
    MOV X9, #10
    MOV X10, #5
    CMP W9, W10
    CSET W9, GT
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoComp1_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    MOV X9, #10
    MOV X10, #5
    CMP W9, W10
    CSET W9, LT
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoComp2_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    LDR D8, F359
    LDR D9, F360
    FCMP D8, D9
    CSET W9, GT
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoComp3_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    LDR D8, F361
    LDR D9, F362
    FCMP D8, D9
    CSET W9, LT
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoComp4_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str363
    LDR X0, =str364
    MOV X1, X9
    BL printf
    LDR X0, =str365
    BL printf
    LDR X9, =resultadoComp1_2
    LDRB W10, [X9]
    LDR X0, =str366
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str367
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str368
    LDR X0, =str369
    MOV X1, X9
    BL printf
    LDR X0, =str370
    BL printf
    LDR X9, =resultadoComp2_2
    LDRB W10, [X9]
    LDR X0, =str371
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str372
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str373
    LDR X0, =str374
    MOV X1, X9
    BL printf
    LDR X0, =str375
    BL printf
    LDR X9, =resultadoComp3_2
    LDRB W10, [X9]
    LDR X0, =str376
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str377
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str378
    LDR X0, =str379
    MOV X1, X9
    BL printf
    LDR X0, =str380
    BL printf
    LDR X9, =resultadoComp4_2
    LDRB W10, [X9]
    LDR X0, =str381
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str382
    BL printf
    // --- End of print call ---
    LDR X12, =resultadoComp1_2
    LDRB W13, [X12]
    MOV X12, #1
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false84
    LDR X12, =resultadoComp2_2
    LDRB W13, [X12]
    MOV X12, #0
    CMP W13, W12
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W11, W13
    B .L_logic_end83
.L_logic_false84:
    MOV W11, #0
.L_logic_end83:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false82
    LDR X11, =resultadoComp3_2
    LDRB W12, [X11]
    MOV X11, #1
    CMP W12, W11
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W10, W12
    B .L_logic_end81
.L_logic_false82:
    MOV W10, #0
.L_logic_end81:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false80
    LDR X10, =resultadoComp4_2
    LDRB W11, [X10]
    MOV X10, #0
    CMP W11, W10
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W9, W11
    B .L_logic_end79
.L_logic_false80:
    MOV W9, #0
.L_logic_end79:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse77
    // 'Then' block
    LDR X9, =puntosOperacionesRelacionales_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosOperacionesRelacionales
    LDR X9, =puntosOperacionesRelacionales_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str383
    LDR X0, =str384
    MOV X1, X9
    BL printf
    LDR X0, =str385
    BL printf
    // --- End of print call ---
    B .Lendif78
.Lelse77:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str386
    LDR X0, =str387
    MOV X1, X9
    BL printf
    LDR X0, =str388
    BL printf
    // --- End of print call ---
.Lendif78:
    // --- Start of print call ---
    LDR X9, =str389
    LDR X0, =str390
    MOV X1, X9
    BL printf
    LDR X0, =str391
    BL printf
    // --- End of print call ---
    MOV X9, #10
    MOV X10, #10
    CMP W9, W10
    CSET W9, GE
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoComp5_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    MOV X9, #10
    MOV X10, #5
    CMP W9, W10
    CSET W9, LE
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoComp6_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    LDR D8, F392
    LDR D9, F393
    FCMP D8, D9
    CSET W9, GE
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoComp7_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    LDR D8, F394
    LDR D9, F395
    FCMP D8, D9
    CSET W9, LE
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoComp8_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str396
    LDR X0, =str397
    MOV X1, X9
    BL printf
    LDR X0, =str398
    BL printf
    LDR X9, =resultadoComp5_2
    LDRB W10, [X9]
    LDR X0, =str399
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str400
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str401
    LDR X0, =str402
    MOV X1, X9
    BL printf
    LDR X0, =str403
    BL printf
    LDR X9, =resultadoComp6_2
    LDRB W10, [X9]
    LDR X0, =str404
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str405
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str406
    LDR X0, =str407
    MOV X1, X9
    BL printf
    LDR X0, =str408
    BL printf
    LDR X9, =resultadoComp7_2
    LDRB W10, [X9]
    LDR X0, =str409
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str410
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str411
    LDR X0, =str412
    MOV X1, X9
    BL printf
    LDR X0, =str413
    BL printf
    LDR X9, =resultadoComp8_2
    LDRB W10, [X9]
    LDR X0, =str414
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str415
    BL printf
    // --- End of print call ---
    LDR X12, =resultadoComp5_2
    LDRB W13, [X12]
    MOV X12, #1
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false92
    LDR X12, =resultadoComp6_2
    LDRB W13, [X12]
    MOV X12, #0
    CMP W13, W12
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W11, W13
    B .L_logic_end91
.L_logic_false92:
    MOV W11, #0
.L_logic_end91:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false90
    LDR X11, =resultadoComp7_2
    LDRB W12, [X11]
    MOV X11, #1
    CMP W12, W11
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W10, W12
    B .L_logic_end89
.L_logic_false90:
    MOV W10, #0
.L_logic_end89:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false88
    LDR X10, =resultadoComp8_2
    LDRB W11, [X10]
    MOV X10, #1
    CMP W11, W10
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W9, W11
    B .L_logic_end87
.L_logic_false88:
    MOV W9, #0
.L_logic_end87:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse85
    // 'Then' block
    LDR X9, =puntosOperacionesRelacionales_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosOperacionesRelacionales
    LDR X9, =puntosOperacionesRelacionales_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str416
    LDR X0, =str417
    MOV X1, X9
    BL printf
    LDR X0, =str418
    BL printf
    // --- End of print call ---
    B .Lendif86
.Lelse85:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str419
    LDR X0, =str420
    MOV X1, X9
    BL printf
    LDR X0, =str421
    BL printf
    // --- End of print call ---
.Lendif86:
    // --- Start of print call ---
    LDR X9, =str422
    LDR X0, =str423
    MOV X1, X9
    BL printf
    LDR X0, =str424
    BL printf
    // --- End of print call ---
    MOV X9, #0
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =puntosOperacionesLogicas_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str425
    LDR X0, =str426
    MOV X1, X9
    BL printf
    LDR X0, =str427
    BL printf
    // --- End of print call ---
    MOV X10, #1
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false94
    MOV X10, #1
    // Left was true, result is right operand
    MOV W9, W10
    B .L_logic_end93
.L_logic_false94:
    MOV W9, #0
.L_logic_end93:
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoAnd1_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    MOV X10, #1
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false96
    MOV X10, #0
    // Left was true, result is right operand
    MOV W9, W10
    B .L_logic_end95
.L_logic_false96:
    MOV W9, #0
.L_logic_end95:
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoAnd2_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    MOV X10, #10
    MOV X11, #10
    CMP W10, W11
    CSET W10, EQ
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false98
    MOV X10, #5
    MOV X11, #5
    CMP W10, W11
    CSET W10, EQ
    // Left was true, result is right operand
    MOV W9, W10
    B .L_logic_end97
.L_logic_false98:
    MOV W9, #0
.L_logic_end97:
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoAnd3_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    MOV X10, #10
    MOV X11, #10
    CMP W10, W11
    CSET W10, EQ
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false100
    MOV X10, #5
    MOV X11, #6
    CMP W10, W11
    CSET W10, EQ
    // Left was true, result is right operand
    MOV W9, W10
    B .L_logic_end99
.L_logic_false100:
    MOV W9, #0
.L_logic_end99:
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoAnd4_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str428
    LDR X0, =str429
    MOV X1, X9
    BL printf
    LDR X0, =str430
    BL printf
    LDR X9, =resultadoAnd1_2
    LDRB W10, [X9]
    LDR X0, =str431
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str432
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str433
    LDR X0, =str434
    MOV X1, X9
    BL printf
    LDR X0, =str435
    BL printf
    LDR X9, =resultadoAnd2_2
    LDRB W10, [X9]
    LDR X0, =str436
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str437
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str438
    LDR X0, =str439
    MOV X1, X9
    BL printf
    LDR X0, =str440
    BL printf
    LDR X9, =resultadoAnd3_2
    LDRB W10, [X9]
    LDR X0, =str441
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str442
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str443
    LDR X0, =str444
    MOV X1, X9
    BL printf
    LDR X0, =str445
    BL printf
    LDR X9, =resultadoAnd4_2
    LDRB W10, [X9]
    LDR X0, =str446
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str447
    BL printf
    // --- End of print call ---
    LDR X12, =resultadoAnd1_2
    LDRB W13, [X12]
    MOV X12, #1
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false108
    LDR X12, =resultadoAnd2_2
    LDRB W13, [X12]
    MOV X12, #0
    CMP W13, W12
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W11, W13
    B .L_logic_end107
.L_logic_false108:
    MOV W11, #0
.L_logic_end107:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false106
    LDR X11, =resultadoAnd3_2
    LDRB W12, [X11]
    MOV X11, #1
    CMP W12, W11
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W10, W12
    B .L_logic_end105
.L_logic_false106:
    MOV W10, #0
.L_logic_end105:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false104
    LDR X10, =resultadoAnd4_2
    LDRB W11, [X10]
    MOV X10, #0
    CMP W11, W10
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W9, W11
    B .L_logic_end103
.L_logic_false104:
    MOV W9, #0
.L_logic_end103:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse101
    // 'Then' block
    LDR X9, =puntosOperacionesLogicas_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosOperacionesLogicas
    LDR X9, =puntosOperacionesLogicas_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str448
    LDR X0, =str449
    MOV X1, X9
    BL printf
    LDR X0, =str450
    BL printf
    // --- End of print call ---
    B .Lendif102
.Lelse101:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str451
    LDR X0, =str452
    MOV X1, X9
    BL printf
    LDR X0, =str453
    BL printf
    // --- End of print call ---
.Lendif102:
    // --- Start of print call ---
    LDR X9, =str454
    LDR X0, =str455
    MOV X1, X9
    BL printf
    LDR X0, =str456
    BL printf
    // --- End of print call ---
    MOV X10, #1
    // Short-circuit OR: check left operand
    CMP W10, #0
    B.NE .L_logic_true110
    MOV X10, #0
    // Left was false, result is right operand
    MOV W9, W10
    B .L_logic_end109
.L_logic_true110:
    MOV W9, #1
.L_logic_end109:
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoOr1_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    MOV X10, #0
    // Short-circuit OR: check left operand
    CMP W10, #0
    B.NE .L_logic_true112
    MOV X10, #0
    // Left was false, result is right operand
    MOV W9, W10
    B .L_logic_end111
.L_logic_true112:
    MOV W9, #1
.L_logic_end111:
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoOr2_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    MOV X10, #10
    MOV X11, #10
    CMP W10, W11
    CSET W10, EQ
    // Short-circuit OR: check left operand
    CMP W10, #0
    B.NE .L_logic_true114
    MOV X10, #5
    MOV X11, #6
    CMP W10, W11
    CSET W10, EQ
    // Left was false, result is right operand
    MOV W9, W10
    B .L_logic_end113
.L_logic_true114:
    MOV W9, #1
.L_logic_end113:
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoOr3_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    MOV X10, #10
    MOV X11, #11
    CMP W10, W11
    CSET W10, EQ
    // Short-circuit OR: check left operand
    CMP W10, #0
    B.NE .L_logic_true116
    MOV X10, #5
    MOV X11, #6
    CMP W10, W11
    CSET W10, EQ
    // Left was false, result is right operand
    MOV W9, W10
    B .L_logic_end115
.L_logic_true116:
    MOV W9, #1
.L_logic_end115:
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoOr4_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str457
    LDR X0, =str458
    MOV X1, X9
    BL printf
    LDR X0, =str459
    BL printf
    LDR X9, =resultadoOr1_2
    LDRB W10, [X9]
    LDR X0, =str460
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str461
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str462
    LDR X0, =str463
    MOV X1, X9
    BL printf
    LDR X0, =str464
    BL printf
    LDR X9, =resultadoOr2_2
    LDRB W10, [X9]
    LDR X0, =str465
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str466
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str467
    LDR X0, =str468
    MOV X1, X9
    BL printf
    LDR X0, =str469
    BL printf
    LDR X9, =resultadoOr3_2
    LDRB W10, [X9]
    LDR X0, =str470
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str471
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str472
    LDR X0, =str473
    MOV X1, X9
    BL printf
    LDR X0, =str474
    BL printf
    LDR X9, =resultadoOr4_2
    LDRB W10, [X9]
    LDR X0, =str475
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str476
    BL printf
    // --- End of print call ---
    LDR X12, =resultadoOr1_2
    LDRB W13, [X12]
    MOV X12, #1
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false124
    LDR X12, =resultadoOr2_2
    LDRB W13, [X12]
    MOV X12, #0
    CMP W13, W12
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W11, W13
    B .L_logic_end123
.L_logic_false124:
    MOV W11, #0
.L_logic_end123:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false122
    LDR X11, =resultadoOr3_2
    LDRB W12, [X11]
    MOV X11, #1
    CMP W12, W11
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W10, W12
    B .L_logic_end121
.L_logic_false122:
    MOV W10, #0
.L_logic_end121:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false120
    LDR X10, =resultadoOr4_2
    LDRB W11, [X10]
    MOV X10, #0
    CMP W11, W10
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W9, W11
    B .L_logic_end119
.L_logic_false120:
    MOV W9, #0
.L_logic_end119:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse117
    // 'Then' block
    LDR X9, =puntosOperacionesLogicas_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosOperacionesLogicas
    LDR X9, =puntosOperacionesLogicas_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str477
    LDR X0, =str478
    MOV X1, X9
    BL printf
    LDR X0, =str479
    BL printf
    // --- End of print call ---
    B .Lendif118
.Lelse117:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str480
    LDR X0, =str481
    MOV X1, X9
    BL printf
    LDR X0, =str482
    BL printf
    // --- End of print call ---
.Lendif118:
    // --- Start of print call ---
    LDR X9, =str483
    LDR X0, =str484
    MOV X1, X9
    BL printf
    LDR X0, =str485
    BL printf
    // --- End of print call ---
    MOV X9, #1
    EOR W9, W9, #1
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoNot1_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    MOV X9, #0
    EOR W9, W9, #1
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoNot2_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    MOV X9, #10
    MOV X10, #10
    CMP W9, W10
    CSET W9, EQ
    EOR W9, W9, #1
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoNot3_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    MOV X9, #10
    MOV X10, #11
    CMP W9, W10
    CSET W9, EQ
    EOR W9, W9, #1
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =resultadoNot4_2
    LDRB W10, [SP]
    STRB W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str486
    LDR X0, =str487
    MOV X1, X9
    BL printf
    LDR X0, =str488
    BL printf
    LDR X9, =resultadoNot1_2
    LDRB W10, [X9]
    LDR X0, =str489
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str490
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str491
    LDR X0, =str492
    MOV X1, X9
    BL printf
    LDR X0, =str493
    BL printf
    LDR X9, =resultadoNot2_2
    LDRB W10, [X9]
    LDR X0, =str494
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str495
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str496
    LDR X0, =str497
    MOV X1, X9
    BL printf
    LDR X0, =str498
    BL printf
    LDR X9, =resultadoNot3_2
    LDRB W10, [X9]
    LDR X0, =str499
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str500
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str501
    LDR X0, =str502
    MOV X1, X9
    BL printf
    LDR X0, =str503
    BL printf
    LDR X9, =resultadoNot4_2
    LDRB W10, [X9]
    LDR X0, =str504
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str505
    BL printf
    // --- End of print call ---
    LDR X12, =resultadoNot1_2
    LDRB W13, [X12]
    MOV X12, #0
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false132
    LDR X12, =resultadoNot2_2
    LDRB W13, [X12]
    MOV X12, #1
    CMP W13, W12
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W11, W13
    B .L_logic_end131
.L_logic_false132:
    MOV W11, #0
.L_logic_end131:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false130
    LDR X11, =resultadoNot3_2
    LDRB W12, [X11]
    MOV X11, #0
    CMP W12, W11
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W10, W12
    B .L_logic_end129
.L_logic_false130:
    MOV W10, #0
.L_logic_end129:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false128
    LDR X10, =resultadoNot4_2
    LDRB W11, [X10]
    MOV X10, #1
    CMP W11, W10
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W9, W11
    B .L_logic_end127
.L_logic_false128:
    MOV W9, #0
.L_logic_end127:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse125
    // 'Then' block
    LDR X9, =puntosOperacionesLogicas_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosOperacionesLogicas
    LDR X9, =puntosOperacionesLogicas_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str506
    LDR X0, =str507
    MOV X1, X9
    BL printf
    LDR X0, =str508
    BL printf
    // --- End of print call ---
    B .Lendif126
.Lelse125:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str509
    LDR X0, =str510
    MOV X1, X9
    BL printf
    LDR X0, =str511
    BL printf
    // --- End of print call ---
.Lendif126:
    // --- Start of print call ---
    LDR X9, =str512
    LDR X0, =str513
    MOV X1, X9
    BL printf
    LDR X0, =str514
    BL printf
    // --- End of print call ---
    MOV X9, #0
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =puntosPrintln_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str515
    LDR X0, =str516
    MOV X1, X9
    BL printf
    LDR X0, =str517
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str518
    LDR X0, =str519
    MOV X1, X9
    BL printf
    LDR X0, =str520
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    MOV X9, #42
    LDR X0, =str521
    MOV X1, X9
    BL printf
    LDR X0, =str522
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR D8, F523
    LDR X0, =str524
    FMOV D0, D8
    BL printf
    LDR X0, =str525
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str526
    LDR X0, =str527
    MOV X1, X9
    BL printf
    LDR X0, =str528
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    MOV X9, #1
    LDR X0, =str529
    LDR X10, =str36
    LDR X11, =str37
    CMP X9, #0
    CSEL X1, X10, X11, NE
    BL printf
    LDR X0, =str530
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str531
    LDR X0, =str532
    MOV X1, X9
    BL printf
    LDR X0, =str533
    BL printf
    // --- End of print call ---
    LDR X9, =puntosPrintln_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosPrintln
    LDR X9, =puntosPrintln_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str534
    LDR X0, =str535
    MOV X1, X9
    BL printf
    LDR X0, =str536
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str537
    LDR X0, =str538
    MOV X1, X9
    BL printf
    LDR X0, =str539
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str540
    LDR X0, =str541
    MOV X1, X9
    BL printf
    LDR X0, =str542
    BL printf
    MOV X9, #42
    LDR X0, =str543
    MOV X1, X9
    BL printf
    LDR X0, =str542
    BL printf
    LDR D8, F544
    LDR X0, =str545
    FMOV D0, D8
    BL printf
    LDR X0, =str546
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str547
    LDR X0, =str548
    MOV X1, X9
    BL printf
    LDR X0, =str549
    BL printf
    MOV X9, #1
    LDR X0, =str550
    LDR X10, =str36
    LDR X11, =str37
    CMP X9, #0
    CSEL X1, X10, X11, NE
    BL printf
    LDR X0, =str549
    BL printf
    LDR X9, =str551
    LDR X0, =str552
    MOV X1, X9
    BL printf
    LDR X0, =str549
    BL printf
    LDR X9, =str553
    LDR X0, =str554
    MOV X1, X9
    BL printf
    LDR X0, =str555
    BL printf
    // --- End of print call ---
    LDR X9, =puntosPrintln_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosPrintln
    LDR X9, =puntosPrintln_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str556
    LDR X0, =str557
    MOV X1, X9
    BL printf
    LDR X0, =str558
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str559
    LDR X0, =str560
    MOV X1, X9
    BL printf
    LDR X0, =str561
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str562
    LDR X0, =str563
    MOV X1, X9
    BL printf
    LDR X0, =str564
    BL printf
    MOV X9, #10
    MOV X10, #5
    ADD X9, X9, X10
    LDR X0, =str565
    MOV X1, X9
    BL printf
    LDR X0, =str566
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str567
    LDR X0, =str568
    MOV X1, X9
    BL printf
    LDR X0, =str569
    BL printf
    MOV X9, #10
    MOV X10, #5
    CMP W9, W10
    CSET W9, GT
    LDR X0, =str570
    LDR X10, =str36
    LDR X11, =str37
    CMP X9, #0
    CSEL X1, X10, X11, NE
    BL printf
    LDR X0, =str571
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str572
    LDR X0, =str573
    MOV X1, X9
    BL printf
    LDR X0, =str574
    BL printf
    MOV X10, #1
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false134
    MOV X10, #0
    // Left was true, result is right operand
    MOV W9, W10
    B .L_logic_end133
.L_logic_false134:
    MOV W9, #0
.L_logic_end133:
    LDR X0, =str575
    LDR X10, =str36
    LDR X11, =str37
    CMP X9, #0
    CSEL X1, X10, X11, NE
    BL printf
    LDR X0, =str576
    BL printf
    // --- End of print call ---
    LDR X9, =puntosPrintln_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosPrintln
    LDR X9, =puntosPrintln_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str577
    LDR X0, =str578
    MOV X1, X9
    BL printf
    LDR X0, =str579
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str580
    LDR X0, =str581
    MOV X1, X9
    BL printf
    LDR X0, =str582
    BL printf
    // --- End of print call ---
    MOV X9, #0
    SUB SP, SP, #16
    STR X9, [SP]
    LDR X9, =puntosValorNulo_2
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str583
    LDR X0, =str584
    MOV X1, X9
    BL printf
    LDR X0, =str585
    BL printf
    // --- End of print call ---
    // Declaring enteroNulo without initializer
    // Declaring decimalNulo without initializer
    // Declaring textoNulo without initializer
    // Declaring booleanoNulo without initializer
    // --- Start of print call ---
    LDR X9, =str586
    LDR X0, =str587
    MOV X1, X9
    BL printf
    LDR X0, =str588
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str589
    LDR X0, =str590
    MOV X1, X9
    BL printf
    LDR X0, =str591
    BL printf
    LDR X9, =enteroNulo_2
    LDRSW X10, [X9]
    LDR X0, =str592
    MOV X1, X10
    BL printf
    LDR X0, =str593
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str594
    LDR X0, =str595
    MOV X1, X9
    BL printf
    LDR X0, =str596
    BL printf
    LDR X9, =decimalNulo_2
    LDR D8, [X9]
    LDR X0, =str597
    FMOV D0, D8
    BL printf
    LDR X0, =str598
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str599
    LDR X0, =str600
    MOV X1, X9
    BL printf
    LDR X0, =str601
    BL printf
    LDR X9, =textoNulo_2
    LDR X10, [X9]
    LDR X0, =str602
    MOV X1, X10
    BL printf
    LDR X0, =str603
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str604
    LDR X0, =str605
    MOV X1, X9
    BL printf
    LDR X0, =str606
    BL printf
    LDR X9, =booleanoNulo_2
    LDRB W10, [X9]
    LDR X0, =str607
    LDR X9, =str36
    LDR X11, =str37
    CMP X10, #0
    CSEL X1, X9, X11, NE
    BL printf
    LDR X0, =str608
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str609
    LDR X0, =str610
    MOV X1, X9
    BL printf
    LDR X0, =str611
    BL printf
    // --- End of print call ---
    LDR X12, =enteroNulo_2
    LDRSW X13, [X12]
    MOV X12, #0
    CMP W13, W12
    CSET W13, EQ
    // Short-circuit AND: check left operand
    CMP W13, #0
    B.EQ .L_logic_false142
    LDR X12, =decimalNulo_2
    LDR D8, [X12]
    LDR D9, F612
    FCMP D8, D9
    CSET W12, EQ
    // Left was true, result is right operand
    MOV W11, W12
    B .L_logic_end141
.L_logic_false142:
    MOV W11, #0
.L_logic_end141:
    // Short-circuit AND: check left operand
    CMP W11, #0
    B.EQ .L_logic_false140
    LDR X11, =textoNulo_2
    LDR X12, [X11]
    LDR X11, =str613
    MOV X0, X12
    MOV X1, X11
    BL strcmp
    CMP W0, #0
    CSET W13, EQ
    // Left was true, result is right operand
    MOV W10, W13
    B .L_logic_end139
.L_logic_false140:
    MOV W10, #0
.L_logic_end139:
    // Short-circuit AND: check left operand
    CMP W10, #0
    B.EQ .L_logic_false138
    LDR X10, =booleanoNulo_2
    LDRB W11, [X10]
    MOV X10, #0
    CMP W11, W10
    CSET W11, EQ
    // Left was true, result is right operand
    MOV W9, W11
    B .L_logic_end137
.L_logic_false138:
    MOV W9, #0
.L_logic_end137:
    // If statement condition check
    CMP W9, #0
    B.EQ .Lelse135
    // 'Then' block
    LDR X9, =puntosValorNulo_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosValorNulo
    LDR X9, =puntosValorNulo_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str614
    LDR X0, =str615
    MOV X1, X9
    BL printf
    LDR X0, =str616
    BL printf
    // --- End of print call ---
    B .Lendif136
.Lelse135:
    // 'Else' block
    // --- Start of print call ---
    LDR X9, =str617
    LDR X0, =str618
    MOV X1, X9
    BL printf
    LDR X0, =str619
    BL printf
    // --- End of print call ---
.Lendif136:
    // --- Start of print call ---
    LDR X9, =str620
    LDR X0, =str621
    MOV X1, X9
    BL printf
    LDR X0, =str622
    BL printf
    // --- End of print call ---
    LDR X9, =puntosValorNulo_2
    LDRSW X10, [X9]
    MOV X9, #1
    ADD X10, X10, X9
    // Storing value for assignment to puntosValorNulo
    LDR X9, =puntosValorNulo_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str623
    LDR X0, =str624
    MOV X1, X9
    BL printf
    LDR X0, =str625
    BL printf
    // --- End of print call ---
    LDR X9, =puntosDeclaracion_2
    LDRSW X10, [X9]
    LDR X9, =puntosAsignacion_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    LDR X9, =puntosOperacionesAritmeticas_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    LDR X9, =puntosOperacionesRelacionales_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    LDR X9, =puntosOperacionesLogicas_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    LDR X9, =puntosPrintln_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    LDR X9, =puntosValorNulo_2
    LDRSW X11, [X9]
    ADD X10, X10, X11
    // Storing value for assignment to puntos
    LDR X9, =puntos_2
    SUB SP, SP, #16
    STR X10, [SP]
    LDR W10, [SP]
    STR W10, [X9]
    ADD SP, SP, #16
    // --- Start of print call ---
    LDR X9, =str626
    LDR X0, =str627
    MOV X1, X9
    BL printf
    LDR X0, =str628
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str629
    LDR X0, =str630
    MOV X1, X9
    BL printf
    LDR X0, =str631
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str632
    LDR X0, =str633
    MOV X1, X9
    BL printf
    LDR X0, =str634
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str635
    LDR X0, =str636
    MOV X1, X9
    BL printf
    LDR X0, =str637
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str638
    LDR X0, =str639
    MOV X1, X9
    BL printf
    LDR X0, =str640
    BL printf
    LDR X9, =puntosDeclaracion_2
    LDRSW X10, [X9]
    LDR X0, =str641
    MOV X1, X10
    BL printf
    LDR X0, =str640
    BL printf
    LDR X9, =str642
    LDR X0, =str643
    MOV X1, X9
    BL printf
    LDR X0, =str644
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str645
    LDR X0, =str646
    MOV X1, X9
    BL printf
    LDR X0, =str647
    BL printf
    LDR X9, =puntosAsignacion_2
    LDRSW X10, [X9]
    LDR X0, =str648
    MOV X1, X10
    BL printf
    LDR X0, =str647
    BL printf
    LDR X9, =str649
    LDR X0, =str650
    MOV X1, X9
    BL printf
    LDR X0, =str651
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str652
    LDR X0, =str653
    MOV X1, X9
    BL printf
    LDR X0, =str654
    BL printf
    LDR X9, =puntosOperacionesAritmeticas_2
    LDRSW X10, [X9]
    LDR X0, =str655
    MOV X1, X10
    BL printf
    LDR X0, =str654
    BL printf
    LDR X9, =str656
    LDR X0, =str657
    MOV X1, X9
    BL printf
    LDR X0, =str658
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str659
    LDR X0, =str660
    MOV X1, X9
    BL printf
    LDR X0, =str661
    BL printf
    LDR X9, =puntosOperacionesRelacionales_2
    LDRSW X10, [X9]
    LDR X0, =str662
    MOV X1, X10
    BL printf
    LDR X0, =str661
    BL printf
    LDR X9, =str663
    LDR X0, =str664
    MOV X1, X9
    BL printf
    LDR X0, =str665
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str666
    LDR X0, =str667
    MOV X1, X9
    BL printf
    LDR X0, =str668
    BL printf
    LDR X9, =puntosOperacionesLogicas_2
    LDRSW X10, [X9]
    LDR X0, =str669
    MOV X1, X10
    BL printf
    LDR X0, =str668
    BL printf
    LDR X9, =str670
    LDR X0, =str671
    MOV X1, X9
    BL printf
    LDR X0, =str672
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str673
    LDR X0, =str674
    MOV X1, X9
    BL printf
    LDR X0, =str675
    BL printf
    LDR X9, =puntosPrintln_2
    LDRSW X10, [X9]
    LDR X0, =str676
    MOV X1, X10
    BL printf
    LDR X0, =str675
    BL printf
    LDR X9, =str677
    LDR X0, =str678
    MOV X1, X9
    BL printf
    LDR X0, =str679
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str680
    LDR X0, =str681
    MOV X1, X9
    BL printf
    LDR X0, =str682
    BL printf
    LDR X9, =puntosValorNulo_2
    LDRSW X10, [X9]
    LDR X0, =str683
    MOV X1, X10
    BL printf
    LDR X0, =str682
    BL printf
    LDR X9, =str684
    LDR X0, =str685
    MOV X1, X9
    BL printf
    LDR X0, =str686
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str687
    LDR X0, =str688
    MOV X1, X9
    BL printf
    LDR X0, =str689
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str690
    LDR X0, =str691
    MOV X1, X9
    BL printf
    LDR X0, =str692
    BL printf
    LDR X9, =puntos_2
    LDRSW X10, [X9]
    LDR X0, =str693
    MOV X1, X10
    BL printf
    LDR X0, =str692
    BL printf
    LDR X9, =str694
    LDR X0, =str695
    MOV X1, X9
    BL printf
    LDR X0, =str696
    BL printf
    // --- End of print call ---
    // --- Start of print call ---
    LDR X9, =str697
    LDR X0, =str698
    MOV X1, X9
    BL printf
    LDR X0, =str699
    BL printf
    // --- End of print call ---
.Lmain_epilogue:
    MOV W0, #0
    LDP X29, X30, [SP], #16
    RET

