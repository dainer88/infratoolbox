# Changelog

All notable changes to this project will be documented in this file.

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
