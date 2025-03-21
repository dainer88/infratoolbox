# Changelog

All notable changes to this project will be documented in this file.

## [v0.2.0] - 2025-03-21

### Added

- **Terragrunt Integration**: Added support for running Terragrunt commands alongside Terraform.
- **Auto-completion**: Implemented shell auto-completion for Bash, Zsh, Fish, and PowerShell.
- **Merge Command Enhancements**: Improved the `merge` command to ensure files end with a blank line before merging.

### Changed

- **Documentation Updates**: Updated the `README.md` to reflect the latest features and commands.
- **Cross-platform Installation**: Updated `install.sh` to support dependency installation on macOS, Linux, and Windows.

## [v0.1.0] - 2024-10-06

### Added

- Initial release of **InfraToolbox** CLI.
- Integrated **Terraform** commands:
  - `apply`, `init`, `plan`, `validate`, `env`.
- Added `doc` command for generating documentation using **terraform-docs**.
- Included `lint` command for running **TFLint** to enforce Terraform best practices.
- Implemented `scan` command to perform security scans with **Checkov**.
- Added `module` command to create new Terraform module templates.
- Introduced `license` command to display the tool's license.
- Streamlined the compilation process and added support for global installation.
- **Homebrew** support for easy installation.

### Documentation

- Complete documentation for installation, usage, and contributing included in the `README.md`.
