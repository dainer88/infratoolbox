# InfraToolbox

InfraToolbox is a CLI tool for managing infrastructure as code, integrating support for Terraform, Terragrunt, and pre-built modules for cloud services.

## Features

- Run Terraform and Terragrunt commands seamlessly.
- Merge multiple Terraform configuration files into a single file.
- Lint Terraform code using TFLint.
- Scan infrastructure for security vulnerabilities using Checkov.
- Generate Terraform documentation using `terraform-docs`.
- Auto-completion support for Bash, Zsh, Fish, and PowerShell.
- Easy to use and extensible.

## Quick Installation

### Using Go

```bash
go install github.com/yourusername/infratoolbox@latest
```

### From Source Code

```bash
git clone https://github.com/yourusername/infratoolbox.git
cd infratoolbox
go build -o infratoolbox
```

## Basic Usage

```bash
# Run a Terraform command
infratoolbox terraform init

# Merge multiple Terraform files
infratoolbox merge file1.tf file2.tf --output merged.tf
```

## Documentation

- [Installation Guide](docs/installation.md)
- [Advanced Use Cases](docs/advanced-usage.md)
- [Troubleshooting](docs/troubleshooting.md)
- [Contributing to the Project](docs/contributing.md)

## Installation

### Using Homebrew

```bash
brew tap yourusername/infratoolbox https://github.com/yourusername/infratoolbox
brew install infratoolbox
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
| `merge`     | Merge multiple Terraform configuration files                    |
| `module`    | Create a new Terraform module template                          |
| `plan`      | Run `terraform plan`                                            |
| `scan`      | Run Checkov to scan infrastructure for security vulnerabilities |
| `terragrunt`| Run Terragrunt commands                                         |
| `validate`  | Validate Terraform configuration                                |

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
    go mod init github.com/yourusername/infratoolbox
    go mod tidy
    go build
    ```

    This command will compile the project and generate an executable file.

4. Run the executable:

    After building, you can run InfraToolbox directly from the command line. For example:

    ```bash
    ./infratoolbox [command]
    ```

    Replace `[command]` with one of the available commands listed in the Usage section.

5. Optional: Install the executable:

    If you want to install the executable globally on your system, you can run:

    ```bash
    go install
    ```

    This will place the `infratoolbox` executable in your Go binaries directory, typically `$GOPATH/bin` or `$HOME/go/bin`.

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
- [Terragrunt](https://terragrunt.gruntwork.io)

## Roadmap

[![Roadmap](https://img.shields.io/badge/Roadmap-View%20here-blue)](./ROADMAP.md)

## License

[![License](https://img.shields.io/badge/License-Apache%202.0-default.svg)](./LICENSE)
