# InfraToolbox

InfraToolbox es una herramienta CLI para gestionar infraestructura como código, integrando soporte para Terraform, Terragrunt y módulos preconstruidos para servicios en la nube.

## Características
- Ejecuta comandos de Terraform y Terragrunt.
- Genera módulos preconstruidos para AWS, Azure y GCP.
- Fusiona múltiples archivos de configuración de Terraform.
- Fácil de usar y extensible.

## Instalación Rápida
### Usando Go
```bash
go install github.com/tuusuario/infratoolbox@latest
```

### Desde el código fuente
```bash
git clone https://github.com/tuusuario/infratoolbox.git
cd infratoolbox
go build -o infratoolbox
```

## Uso Básico
```bash
# Ejecutar un comando de Terraform
infratoolbox terraform init

# Generar un módulo preconstruido
infratoolbox template aws s3 my-s3-module
```

## Documentación
- [Guía de instalación](docs/installation.md)
- [Casos de uso avanzados](docs/advanced-usage.md)
- [Resolución de problemas](docs/troubleshooting.md)
- [Contribuir al proyecto](docs/contributing.md)

## Features

- **Unified CLI**: Run multiple infrastructure tools like Terraform, Terraform Docs, TFEnv, TFLint, and Checkov with one command-line tool.
- **Module Creation**: Automatically generate new Terraform module templates, ready for customization.
- **Security Scans**: Ensure the security of your infrastructure using Checkov.
- **Linting**: Enforce Terraform best practices by using TFLint.
- **Plan and Apply**: Streamline Terraform planning and applying commands, ensuring smoother workflow.

## Installation

### Using Homebrew

```bash
brew tap dainer88/infratoolbox https://github.com/dainer88/infratoolbox
brew install infratoolbox
```

### Manual installation

#### Macos

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/infratoolbox.git
   ```

2. Navigate to the directory:

    ```bash
    cd infratoolbox
    ```

3. Run the install.sh script to install dependencies and configure the environment:

    ```bash
    ./install.sh
    ```

## Auto-completion

InfraToolbox automatically sets up shell auto-completion during installation for the following shells:

- **Bash**: The completion script is installed in `/etc/bash_completion.d/`.
- **Zsh**: The completion script is installed in the Zsh functions directory.
- **Fish**: The completion script is installed in `~/.config/fish/completions/`.
- **PowerShell**: The completion script is generated in `$HOME/infratoolbox.ps1`.

### Enabling Auto-completion

- **Bash**: Ensure that Bash completion is enabled on your system.
- **Zsh**: Ensure that `autoload -U compinit; compinit` is added to your `~/.zshrc`.
- **Fish**: No additional steps are required.
- **PowerShell**: Add the following line to your PowerShell profile:
  ```powershell
  . $HOME/infratoolbox.ps1
  ```

## Usage

InfraToolbox is designed to manage infrastructure efficiently. You can run any of the following commands:

```bash
infratoolbox [command]
```

### Available Commands

| Command     | Description                                                     |
|-------------|-----------------------------------------------------------------|
| `apply`     | Run `terraform apply`                                           |
| `doc`       | Generate Terraform documentation using `terraform-docs`         |
| `env`       | Run `terraform env`                                             |
| `init`      | Run `terraform init`                                            |
| `license`   | Show the license under which InfraToolbox is distributed        |
| `lint`      | Run TFLint to lint Terraform code                               |
| `module`    | Create a new Terraform module template                          |
| `plan`      | Run `terraform plan`                                            |
| `scan`      | Run Checkov to scan infrastructure for security vulnerabilities |
| `validate`  | Display the current version of InfraToolbox                     |

## Compiling the Project

To compile InfraToolbox from source, follow these steps:

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/infratoolbox.git
    ```

2. Navigate to the project directory:

    ```bash
    cd infratoolbox
    ```

3. Compile the project:

    ```bash
    go mod init github.com/dainer88/infratoolbox
    go mod tidy
    go build
    ```

    This command will compile the project and generate an executable file.

4. Run the executable:

    After building, you can run InfraToolbox directly from the command line. For example:

    ```bash
    ./infratoolbox [command]
    ```

    Replace [command] with one of the available commands listed in the Usage section.

5. Optional: Install the executable:

    If you want to install the executable globally on your system, you can run:

    ```bash
    go install
    ```

    This will place the infratoolbox executable in your Go binaries directory, typically $GOPATH/bin or $HOME/go/bin.

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Submit a pull request with a clear description of the changes.

## Acknowledgements

InfraToolbox utilizes the following open-source tools:

- [Terraform](https://www.terraform.io)
- [Terraform Docs](https://terraform-docs.io)
- [TFEnv](https://github.com/tfutils/tfenv)
- [TFLint](https://github.com/terraform-linters/tflint)
- [Checkov](https://www.checkov.io)

## Roadmap

[![Roadmap](https://img.shields.io/badge/Roadmap-View%20here-blue)](./ROADMAP.md)

## License

[![License](https://img.shields.io/badge/License-Apache%202.0-default.svg)](./LICENSE)
