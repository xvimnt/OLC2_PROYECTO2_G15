# Instalación de QEMU para Emulación ARM64

## Problema
El error `/usr/bin/qemu-aarch64-static: not found` indica que QEMU no está instalado en tu sistema WSL.

## Solución

### Opción 1: Instalación Automática (Recomendada)
```powershell
# En PowerShell, ejecuta:
wsl -d Ubuntu sudo apt update
wsl -d Ubuntu sudo apt install -y qemu-user-static
```

### Opción 2: Instalación Manual
1. Abre una terminal de Ubuntu WSL:
   ```powershell
   wsl -d Ubuntu
   ```

2. Actualiza los paquetes:
   ```bash
   sudo apt update
   ```

3. Instala QEMU:
   ```bash
   sudo apt install -y qemu-user-static
   ```

4. Verifica la instalación:
   ```bash
   qemu-aarch64-static --version
   ```

5. Sale de WSL:
   ```bash
   exit
   ```

### Verificación
Después de la instalación, puedes verificar que QEMU esté disponible:
```powershell
wsl -d Ubuntu which qemu-aarch64-static
```

### Uso
Una vez instalado QEMU, puedes usar los targets del Makefile:
```powershell
make run-arm           # Ejecutar el binario ARM con QEMU
make test-v-flow       # Flujo completo de compilación y ejecución
```

## Troubleshooting

### Si WSL Ubuntu no está instalado:
```powershell
wsl --install -d Ubuntu
```

### Si necesitas reiniciar Ubuntu WSL:
```powershell
wsl --shutdown
wsl -d Ubuntu
```

### Si tienes problemas de permisos:
Asegúrate de que tu usuario tenga permisos sudo en Ubuntu WSL.
