# Guía de Instalación

## Requisitos Previos
- Go 1.20 o superior.
- Git instalado en tu sistema.

## Instalación en Diferentes Entornos

### macOS
1. Instala Go usando Homebrew:
   ```bash
   brew install go
   ```
2. Instala InfraToolbox:
   ```bash
   go install github.com/tuusuario/infratoolbox@latest
   ```

### Linux
1. Instala Go:
   ```bash
   sudo apt update
   sudo apt install golang
   ```
2. Instala InfraToolbox:
   ```bash
   go install github.com/tuusuario/infratoolbox@latest
   ```

### Windows
1. Descarga e instala Go desde [golang.org](https://golang.org/dl/).
2. Instala InfraToolbox:
   ```powershell
   go install github.com/tuusuario/infratoolbox@latest
   ```

## Verificar la Instalación
Ejecuta el siguiente comando para verificar que InfraToolbox está instalado correctamente:
```bash
infratoolbox --version
```
