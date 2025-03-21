# Resolución de Problemas

## Error: "command not found"
Asegúrate de que el directorio `$GOPATH/bin` esté en tu variable de entorno `PATH`.

### Solución:
Agrega lo siguiente a tu archivo `.bashrc` o `.zshrc`:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## Error: "Terragrunt is not installed"
InfraToolbox requiere que Terragrunt esté instalado para ejecutar comandos de Terragrunt.

### Solución:
Instala Terragrunt siguiendo las instrucciones en [terragrunt.io](https://terragrunt.io/).
