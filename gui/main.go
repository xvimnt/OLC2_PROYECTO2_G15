package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Tema claro personalizado
type myTheme struct{}

// Función auxiliar para obtener el mínimo de dos enteros
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (m *myTheme) Font(s fyne.TextStyle) fyne.Resource { return theme.DefaultTheme().Font(s) }
func (m *myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		return color.White // Fondo general blanco
	case theme.ColorNameButton:
		return color.RGBA{R: 230, G: 230, B: 230, A: 255} // Botones gris claro
	case theme.ColorNameDisabled:
		return color.RGBA{R: 200, G: 200, B: 200, A: 255}
	case theme.ColorNameInputBackground:
		return color.White // Fondo de entradas blanco
	case theme.ColorNameForeground:
		return color.Black // Texto negro
	case theme.ColorNamePrimary:
		return color.Black // Acentos en negro
	case theme.ColorNamePlaceHolder:
		return color.RGBA{R: 120, G: 120, B: 120, A: 255}
	case theme.ColorNameHover:
		return color.RGBA{R: 210, G: 210, B: 210, A: 255}
	}
	return theme.DefaultTheme().Color(n, v)
}
func (m *myTheme) Icon(n fyne.ThemeIconName) fyne.Resource { return theme.DefaultTheme().Icon(n) }
func (m *myTheme) Size(n fyne.ThemeSizeName) float32 {
	if n == theme.SizeNameText {
		return 20
	}
	return theme.DefaultTheme().Size(n)
}

// Fondo con barra de título personalizada (claro)
func fondoConTitulo(titulo string, contenido fyne.CanvasObject) fyne.CanvasObject {
	bg := canvas.NewRectangle(color.White)
	tituloLabel := widget.NewLabelWithStyle(titulo, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	tituloLabel.Alignment = fyne.TextAlignCenter
	tituloLabel.Wrapping = fyne.TextTruncate
	tituloBar := container.NewVBox(
		canvas.NewRectangle(color.White),
		tituloLabel,
	)
	return container.NewMax(
		bg,
		container.NewBorder(tituloBar, nil, nil, nil, contenido),
	)
}

// --- Funciones de lógica (sin cambios) ---

func obtenerSVGAST(gramatica, codigo string) (string, error) {
	payload := map[string]string{
		"grammar":    gramatica,
		"lexgrammar": gramatica,
		"input":      codigo,
		"start":      "program",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://lab.antlr.org/parse/", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	type ANTLRResponse struct {
		Result struct {
			SVGTree string `json:"svgtree"`
		} `json:"result"`
	}

	var antlrResp ANTLRResponse
	err = json.Unmarshal(body, &antlrResp)
	if err != nil {
		return "", err
	}

	return antlrResp.Result.SVGTree, nil
}

func leerGramatica() (string, error) {
	data, err := os.ReadFile("../grammar/VLangCherry.g4")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func formatearTabla(tabla string) string {
	lines := strings.Split(tabla, "\n")
	var col1, col2 int
	for _, line := range lines {
		cols := strings.Fields(line)
		if len(cols) > 0 && len(cols[0]) > col1 {
			col1 = len(cols[0])
		}
		if len(cols) > 1 && len(cols[1]) > col2 {
			col2 = len(cols[1])
		}
	}
	var result strings.Builder
	for _, line := range lines {
		cols := strings.Fields(line)
		if len(cols) == 0 {
			result.WriteString("\n")
			continue
		}
		result.WriteString(fmt.Sprintf("%-*s  ", col1, cols[0]))
		if len(cols) > 1 {
			result.WriteString(fmt.Sprintf("%-*s  ", col2, cols[1]))
		}
		if len(cols) > 2 {
			valor := strings.Join(cols[2:], " ")
			result.WriteString(valor)
		}
		result.WriteString("\n")
	}
	return result.String()
}

// --- Ventanas personalizadas ---

func mostrarTablaSimbolos(tabla string, parent fyne.Window) {
	lines := strings.Split(strings.TrimSpace(tabla), "\n")
	// Ignora líneas vacías al inicio
	for len(lines) > 0 && strings.TrimSpace(lines[0]) == "" {
		lines = lines[1:]
	}

	if len(lines) < 2 {
		dialog.ShowInformation("Tabla de Símbolos", "No hay símbolos para mostrar.", parent)
		return
	}

	// Procesar encabezados y filas usando tabuladores
	headers := strings.Split(lines[0], "\t")
	rows := [][]string{}
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "" {
			continue
		}
		cols := strings.Split(line, "\t")
		for len(cols) < len(headers) {
			cols = append(cols, "")
		}
		rows = append(rows, cols)
	}

	table := widget.NewTable(
		func() (int, int) { return len(rows) + 1, len(headers) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			if id.Row == 0 {
				label.SetText(headers[id.Col])
				label.TextStyle = fyne.TextStyle{Bold: true}
			} else {
				label.SetText(rows[id.Row-1][id.Col])
				label.TextStyle = fyne.TextStyle{Bold: true}
			}
		},
	)

	// Ajusta el tamaño mínimo de columnas y filas
	for i := 0; i < len(headers); i++ {
		table.SetColumnWidth(i, 180)
	}
	for i := 0; i < len(rows)+1; i++ {
		table.SetRowHeight(i, 30)
	}

	w := fyne.CurrentApp().NewWindow("Tabla de Símbolos")
	btnCerrar := widget.NewButton("Cerrar", func() { w.Close() })
	btnCerrar.Importance = widget.HighImportance
	bottomBar := container.NewCenter(btnCerrar)
	w.SetContent(fondoConTitulo("Tabla de Símbolos",
		container.NewBorder(nil, bottomBar, nil, nil, table),
	))
	w.Resize(fyne.NewSize(700, 600))
	w.Show()
}

func mostrarTablaErrores(tabla string, parent fyne.Window) {
	lines := strings.Split(strings.TrimSpace(tabla), "\n")

	// Ignora líneas vacías al inicio
	for len(lines) > 0 && strings.TrimSpace(lines[0]) == "" {
		lines = lines[1:]
	}

	if len(lines) < 2 {
		dialog.ShowInformation("Tabla de Errores", "No hay errores para mostrar.", parent)
		return
	}

	headers := strings.Split(lines[0], "\t")
	rows := [][]string{}
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "" {
			continue
		}
		cols := strings.Split(line, "\t")
		for len(cols) < len(headers) {
			cols = append(cols, "")
		}
		rows = append(rows, cols)
	}

	table := widget.NewTable(
		func() (int, int) { return len(rows) + 1, len(headers) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			if id.Row == 0 {
				label.SetText(headers[id.Col])
				label.TextStyle = fyne.TextStyle{Bold: true}
			} else {
				label.SetText(rows[id.Row-1][id.Col])
				label.TextStyle = fyne.TextStyle{Bold: true}
			}
		},
	)

	// Ajusta el tamaño mínimo de columnas y filas
	for i := 0; i < len(headers); i++ {
		table.SetColumnWidth(i, 180)
	}
	for i := 0; i < len(rows)+1; i++ {
		table.SetRowHeight(i, 30)
	}

	w := fyne.CurrentApp().NewWindow("Tabla de Errores")
	btnCerrar := widget.NewButton("Cerrar", func() { w.Close() })
	btnCerrar.Importance = widget.HighImportance
	bottomBar := container.NewCenter(btnCerrar)
	w.SetContent(fondoConTitulo("Tabla de Errores",
		container.NewBorder(nil, bottomBar, nil, nil, table),
	))
	w.Resize(fyne.NewSize(700, 600))
	w.Show()
}

func mostrarIntegrantes(parent fyne.Window) {
	dialog.ShowInformation(
		"Integrantes Grupo 15",
		`Gabriel Orlando Chinchilla Vásquez
Javier Alejandro Monterroso Lopez
Luis Andres Calvo Arreaga`,
		parent,
	)
}

// --- MAIN Fyne personalizado ---

func main() {
	myApp := app.NewWithID("com.vlangcherry.ide")
	myApp.Settings().SetTheme(&myTheme{})
	myWindow := myApp.NewWindow("VLang Cherry IDE")

	editor := widget.NewMultiLineEntry()
	editor.SetPlaceHolder("Escribe tu código VLang Cherry aquí...")
	editorScroller := container.NewVScroll(editor)
	editorScroller.SetMinSize(fyne.NewSize(800, 300))
	editorScrollerBG := canvas.NewRectangle(color.White)

	consola := widget.NewMultiLineEntry()
	consola.SetPlaceHolder("Consola de salida...")
	consola.Wrapping = fyne.TextWrapWord
	consolaScroller := container.NewVScroll(consola)
	consolaScroller.SetMinSize(fyne.NewSize(800, 250))
	consolaScrollerBG := canvas.NewRectangle(color.White)

	btnAbrir := widget.NewButtonWithIcon("Abrir", theme.FolderOpenIcon(), nil)
	btnGuardar := widget.NewButtonWithIcon("Guardar", theme.DocumentSaveIcon(), nil)
	btnEjecutar := widget.NewButtonWithIcon("Ejecutar", theme.MediaPlayIcon(), nil)
	btnTraducir := widget.NewButtonWithIcon("Traducir", theme.DocumentIcon(), nil)
	btnErrores := widget.NewButtonWithIcon("Errores", theme.ErrorIcon(), nil)
	btnSimbolos := widget.NewButtonWithIcon("Tabla de Símbolos", theme.InfoIcon(), nil)
	btnAST := widget.NewButtonWithIcon("AST", theme.VisibilityIcon(), nil)
	btnIntegrantes := widget.NewButtonWithIcon("Integrantes", theme.AccountIcon(), nil)
	btnIntegrantes.OnTapped = func() {
		mostrarIntegrantes(myWindow)
	}
	// Asigna las funciones originales a los botones
	btnAbrir.OnTapped = func() {
		openDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader == nil {
				return
			}
			defer reader.Close()
			data, err := io.ReadAll(reader)
			if err == nil {
				editor.SetText(string(data))
			}
		}, myWindow)
		openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".v"}))
		openDialog.Show()
	}
	btnGuardar.OnTapped = func() {
		saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil || writer == nil {
				return
			}
			defer writer.Close()
			writer.Write([]byte(editor.Text))
		}, myWindow)
		saveDialog.SetFilter(storage.NewExtensionFileFilter([]string{".v"}))
		saveDialog.Show()
	}
	btnEjecutar.OnTapped = func() {
		consola.SetText("Ejecutando código...\n")

		// Crear archivo temporal con el código
		tmpFile, err := ioutil.TempFile("", "*.v")
		if err != nil {
			consola.SetText("Error creando archivo temporal: " + err.Error())
			return
		}
		defer os.Remove(tmpFile.Name())
		tmpFile.WriteString(editor.Text)
		tmpFile.Close()

		// Traducir a ARM64 Assembly
		cmd := exec.Command(".\\OLC2_PROYECTO2_G15.exe", "translate", tmpFile.Name())
		cmd.Dir = ".." // Cambiar al directorio padre donde está el ejecutable
		output, err := cmd.CombinedOutput()
		if err != nil {
			consola.SetText("Error en traducción: " + err.Error())
			return
		}

		// Compilar Assembly a binario ARM64
		cmd = exec.Command("make", "build-arm")
		cmd.Dir = ".." // Ejecutar en el directorio padre
		output, err = cmd.CombinedOutput()
		if err != nil {
			consola.SetText("Error en compilación ARM64:\n" + string(output))
			return
		}

		// Ejecutar el binario ARM64 con QEMU
		cmd = exec.Command("make", "run-arm")
		cmd.Dir = ".." // Ejecutar en el directorio padre
		output, err = cmd.CombinedOutput()
		if err != nil {
			consola.SetText("Error ejecutando con QEMU:\n" + string(output))
			return
		}
		
		// Filtrar el mensaje de QEMU y solo mostrar la salida del programa
		outputStr := string(output)
		if strings.Contains(outputStr, "Running ARM executable with QEMU...") {
			// Buscar la línea después del mensaje de QEMU
			lines := strings.Split(outputStr, "\n")
			var filteredLines []string
			skipNext := false
			for _, line := range lines {
				if strings.Contains(line, "Running ARM executable with QEMU...") {
					skipNext = true
					continue
				}
				if !skipNext {
					filteredLines = append(filteredLines, line)
				} else {
					// Después del mensaje de QEMU, incluir todas las líneas
					filteredLines = append(filteredLines, line)
					skipNext = false
				}
			}
			outputStr = strings.Join(filteredLines, "\n")
		}
		
		consola.SetText(outputStr)
	}
	
	btnTraducir.OnTapped = func() {
		// Crear archivo temporal con el código
		tmpFile, err := ioutil.TempFile("", "*.v")
		if err != nil {
			consola.SetText("Error creando archivo temporal: " + err.Error())
			return
		}
		defer os.Remove(tmpFile.Name())
		tmpFile.WriteString(editor.Text)
		tmpFile.Close()

		// Traducir a ARM64 Assembly
		cmd := exec.Command(".\\OLC2_PROYECTO2_G15.exe", "translate", tmpFile.Name())
		cmd.Dir = ".." // Cambiar al directorio padre donde está el ejecutable
		output, err := cmd.CombinedOutput()
		if err != nil {
			consola.SetText("Error en traducción: " + err.Error())
			return
		}
		
		// Filtrar las líneas de debug y solo mostrar el código Assembly
		outputStr := string(output)
		lines := strings.Split(outputStr, "\n")
		var assemblyLines []string
		inAssembly := false
		
		for _, line := range lines {
			// Iniciar captura después de "--- Generated assembly code ---"
			if strings.Contains(line, "--- Generated assembly code ---") {
				inAssembly = true
				continue
			}
			// Terminar captura antes de "--- End of assembly code ---"
			if strings.Contains(line, "--- End of assembly code ---") {
				inAssembly = false
				continue
			}
			// Capturar solo las líneas del assembly
			if inAssembly {
				assemblyLines = append(assemblyLines, line)
			}
		}
		
		// Mostrar solo el código Assembly
		consola.SetText(strings.Join(assemblyLines, "\n"))
	}
	
	btnErrores.OnTapped = func() {
		tmpFile, err := ioutil.TempFile("", "*.v")
		if err != nil {
			consola.SetText("Error creando archivo temporal")
			return
		}
		defer os.Remove(tmpFile.Name())
		tmpFile.WriteString(editor.Text)
		tmpFile.Close()

		cmd := exec.Command("../OLC2_PROYECTO2_G15", "run", tmpFile.Name())
		cmd.Env = append(os.Environ(), "VLANG_ERRORS=1")
		output, _ := cmd.CombinedOutput()

		outStr := string(output)
		idx := strings.Index(outStr, "=== Tabla de Errores ===")
		if idx != -1 {
			tabla := outStr[idx+len("=== Tabla de Errores ==="):]

			mostrarTablaErrores(tabla, myWindow)
		} else {
			dialog.ShowInformation("Tabla de Errores", "No se encontró la tabla de errores.", myWindow)
		}
	}
	btnSimbolos.OnTapped = func() {
		tmpFile, err := ioutil.TempFile("", "*.v")
		if err != nil {
			consola.SetText("Error creando archivo temporal")
			return
		}
		defer os.Remove(tmpFile.Name())
		tmpFile.WriteString(editor.Text)
		tmpFile.Close()

		cmd := exec.Command("../OLC2_PROYECTO2_G15", "run", tmpFile.Name())
		cmd.Env = append(os.Environ(), "VLANG_SYMBOLS=1")
		output, err := cmd.CombinedOutput()
		if err != nil {
			consola.SetText("Error al ejecutar:\n" + string(output))
			return
		}

		outStr := string(output)
		idx := bytes.Index([]byte(outStr), []byte("=== Tabla de Símbolos ==="))
		if idx != -1 {
			tabla := outStr[idx+len("=== Tabla de Símbolos ==="):]

			mostrarTablaSimbolos(tabla, myWindow)
		} else {
			dialog.ShowInformation("Tabla de Símbolos", "No se encontró la tabla de símbolos.", myWindow)
		}
	}
	btnAST.OnTapped = func() {
		consola.SetText("Generando AST, espera un momento...")

		gramatica, err := leerGramatica()
		if err != nil {
			consola.SetText("Error leyendo gramática: " + err.Error())
			return
		}
		codigo := editor.Text

		svg, err := obtenerSVGAST(gramatica, codigo)
		if err != nil {
			consola.SetText("Error generando AST: " + err.Error())
			return
		}

		err = os.WriteFile("arbol_cst.svg", []byte(svg), 0644)
		if err != nil {
			consola.SetText("Error guardando SVG: " + err.Error())
			return
		}

		cmd := exec.Command("rsvg-convert", "-b", "white", "arbol_cst.svg", "-o", "arbol_cst.png")
		err = cmd.Run()

		// Verifica si la conversión fue exitosa y si el archivo PNG existe
		if err == nil {
			if _, statErr := os.Stat("arbol_cst.png"); statErr == nil {
				// Abre el PNG normalmente
				exec.Command("xdg-open", "arbol_cst.png").Start()
				consola.SetText("AST generado y abierto en visor externo.")
			} else {
				// Si no existe el PNG, abre el SVG en Firefox
				exec.Command("firefox", "arbol_cst.svg").Start()
				consola.SetText("No se pudo generar el PNG, abriendo SVG en Firefox.")
			}
		} else {
			// Si hubo error en la conversión, abre el SVG en Firefox
			exec.Command("firefox", "arbol_cst.svg").Start()
			consola.SetText("No se pudo convertir el SVG a PNG, abriendo SVG en Firefox.")
		}
	}

	barra := container.NewHBox(
		btnAbrir, btnGuardar, btnEjecutar, btnTraducir, btnErrores, btnSimbolos, btnAST, btnIntegrantes,
	)
	barraBG := canvas.NewRectangle(color.White)

	content := container.NewVBox(
		container.NewMax(barraBG, barra),
		container.NewMax(editorScrollerBG, editorScroller),
		widget.NewLabel("Consola:"),
		container.NewMax(consolaScrollerBG, consolaScroller),
	)

	myWindow.SetContent(fondoConTitulo("VLang Cherry IDE", content))
	myWindow.Resize(fyne.NewSize(1000, 800))
	myWindow.ShowAndRun()
}
