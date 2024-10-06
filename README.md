# InfraToolbox

InfraToolbox integrates Terraform, TFEnv, TFLint, Checkov, and other tools into a single CLI for managing infrastructure efficiently. It simplifies the execution of multiple infrastructure-as-code tools, providing a unified interface for everyday operations.

## Features

- **Unified CLI**: Run multiple infrastructure tools like Terraform, Terragrunt, TFLint, and Checkov with one command-line tool.
- **Module Creation**: Automatically generate new Terraform module templates, ready for customization.
- **Security Scans**: Ensure the security of your infrastructure using Checkov.
- **Linting**: Enforce Terraform best practices by using TFLint.
- **Plan and Apply**: Streamline Terraform planning and applying commands, ensuring smoother workflow.

## Installation

### Using Homebrew

```bash
brew tap dainer88/infratoolbox-tap
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
    ````

3. Run the install.sh script to install dependencies and configure the environment:

    ```bash
    ./install.sh
    ````

4. Add InfraToolbox to your system's PATH:

    ```bash
    export PATH=$PATH:/path/to/infratoolbox/bin
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

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](./LICENSE) file for details.
