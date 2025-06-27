# Script de PowerShell para instalar QEMU en WSL Ubuntu
Write-Host "=== Instalando QEMU en WSL Ubuntu ===" -ForegroundColor Yellow

# Verificar si WSL Ubuntu está disponible
try {
    $wslList = wsl --list --quiet
    if ($wslList -notcontains "Ubuntu") {
        Write-Host "Error: Ubuntu WSL no está instalado." -ForegroundColor Red
        Write-Host "Instala Ubuntu WSL con: wsl --install -d Ubuntu" -ForegroundColor Yellow
        exit 1
    }
} catch {
    Write-Host "Error: WSL no está disponible en este sistema." -ForegroundColor Red
    exit 1
}

# Verificar si QEMU ya está instalado
Write-Host "Verificando si QEMU ya está instalado..." -ForegroundColor Cyan
try {
    $qemuCheck = wsl -d Ubuntu bash -c "command -v qemu-aarch64-static"
    if ($qemuCheck) {
        Write-Host "QEMU ya está instalado:" -ForegroundColor Green
        wsl -d Ubuntu qemu-aarch64-static --version | Select-Object -First 1
        exit 0
    }
} catch {
    # Continuar con la instalación si hay error verificando
}

Write-Host "QEMU no está instalado. Procediendo con la instalación..." -ForegroundColor Yellow

# Instalar QEMU
Write-Host "Actualizando lista de paquetes..." -ForegroundColor Cyan
try {
    wsl -d Ubuntu bash -c "sudo apt update -y"
    if ($LASTEXITCODE -ne 0) {
        throw "Error actualizando paquetes"
    }
} catch {
    Write-Host "Error: No se pudo actualizar la lista de paquetes." -ForegroundColor Red
    Write-Host "Intenta manualmente: wsl -d Ubuntu sudo apt update" -ForegroundColor Yellow
    exit 1
}

Write-Host "Instalando QEMU..." -ForegroundColor Cyan
try {
    wsl -d Ubuntu bash -c "sudo apt install -y qemu-user-static"
    if ($LASTEXITCODE -ne 0) {
        throw "Error instalando QEMU"
    }
} catch {
    Write-Host "Error: No se pudo instalar QEMU." -ForegroundColor Red
    Write-Host "Intenta manualmente: wsl -d Ubuntu sudo apt install -y qemu-user-static" -ForegroundColor Yellow
    exit 1
}

# Verificar instalación
Write-Host "Verificando instalación..." -ForegroundColor Cyan
try {
    $qemuPath = wsl -d Ubuntu bash -c "command -v qemu-aarch64-static"
    if ($qemuPath) {
        Write-Host "¡QEMU instalado exitosamente!" -ForegroundColor Green
        wsl -d Ubuntu qemu-aarch64-static --version | Select-Object -First 1
        Write-Host "Ahora puedes usar: make run-arm" -ForegroundColor Green
    } else {
        throw "QEMU no se encontró después de la instalación"
    }
} catch {
    Write-Host "Error: La instalación falló o QEMU no está accesible." -ForegroundColor Red
    exit 1
}
