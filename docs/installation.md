# Installation Guide

## Prerequisites

- Go 1.20 or higher.
- Git installed on your system.

## Installation on Different Environments

### macOS

1. Install Go using Homebrew:

   ```bash
   brew install go
   ```

2. Install InfraToolbox:

   ```bash
   go install github.com/yourusername/infratoolbox@latest
   ```

### Linux

1. Install Go:

   ```bash
   sudo apt update
   sudo apt install golang
   ```

2. Install InfraToolbox:

   ```bash
   go install github.com/yourusername/infratoolbox@latest
   ```

### Windows

1. Download and install Go from [golang.org](https://golang.org/dl/).
2. Install InfraToolbox:

   ```powershell
   go install github.com/yourusername/infratoolbox@latest
   ```

## Verify Installation

Run the following command to verify that InfraToolbox is installed correctly:

```bash
infratoolbox --version
```
