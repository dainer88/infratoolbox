# Casos de Uso Avanzados

## Uso de Terragrunt
Puedes ejecutar comandos de Terragrunt directamente desde InfraToolbox:
```bash
infratoolbox terragrunt plan --terragrunt-config=terragrunt.hcl
```

## Personalización de Módulos Preconstruidos
Después de generar un módulo preconstruido, puedes personalizar los archivos generados para adaptarlos a tus necesidades:
```bash
infratoolbox template aws s3 my-custom-s3-module
cd my-custom-s3-module
# Edita los archivos main.tf, variables.tf, etc.
```

## Integración con CI/CD
Ejemplo de un pipeline en GitHub Actions:
```yaml
name: Terraform Plan

on:
  push:
    branches:
      - main

jobs:
  terraform:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v2

      - name: Install InfraToolbox
        run: |
          go install github.com/tuusuario/infratoolbox@latest

      - name: Run Terraform Plan
        run: |
          infratoolbox terraform plan
```
