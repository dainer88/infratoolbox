#!/bin/bash

# Installation script for InfraToolbox

# Function to print messages
function echo_message {
    echo "=============================="
    echo "$1"
    echo "=============================="
}

# Detect the operating system
OS=$(uname -s)

# Function to install a tool if it's not already installed
function install_tool {
    TOOL_NAME=$1
    INSTALL_CMD=$2
    if command -v $TOOL_NAME &> /dev/null; then
        echo_message "$TOOL_NAME is already installed. Skipping installation."
    else
        echo_message "Installing $TOOL_NAME..."
        eval $INSTALL_CMD
    fi
}

# Install necessary tools based on the operating system
case "$OS" in
    Darwin) # macOS
        echo_message "Detected macOS. Using Homebrew for installation."
        if ! command -v brew &> /dev/null; then
            echo_message "Homebrew is not installed. Please install it first."
            exit 1
        fi
        install_tool terraform-docs "brew install terraform-docs"
        install_tool tfenv "brew install tfenv"
        install_tool tflint "brew install tflint"
        install_tool checkov "brew install checkov"
        install_tool terragrunt "brew install terragrunt"
        ;;
    Linux) # Linux
        echo_message "Detected Linux. Using apt or yum for installation."
        if command -v apt &> /dev/null; then
            install_tool terraform-docs "sudo apt update && sudo apt install -y terraform-docs"
            install_tool tfenv "sudo apt update && sudo apt install -y tfenv"
            install_tool tflint "sudo apt update && sudo apt install -y tflint"
            install_tool checkov "sudo apt update && sudo apt install -y python3-pip && pip3 install checkov"
            install_tool terragrunt "sudo apt update && sudo apt install -y terragrunt"
        elif command -v yum &> /dev/null; then
            install_tool terraform-docs "sudo yum install -y terraform-docs"
            install_tool tfenv "sudo yum install -y tfenv" 
            install_tool tflint "sudo yum install -y tflint"
            install_tool checkov "sudo yum install -y python3 && pip3 install checkov"
            install_tool terragrunt "sudo yum install -y terragrunt"
        else
            echo_message "Neither apt nor yum package manager found. Please install dependencies manually."
            exit 1
        fi
        ;;
    CYGWIN*|MINGW*|MSYS*|Windows_NT) # Windows
        echo_message "Detected Windows. Using Chocolatey for installation."
        if ! command -v choco &> /dev/null; then
            echo_message "Chocolatey is not installed. Please install it first: https://chocolatey.org/install"
            exit 1
        fi
        install_tool terraform-docs "choco install terraform-docs -y"
        install_tool tfenv "choco install tfenv -y"
        install_tool tflint "choco install tflint -y"
        install_tool checkov "choco install python -y && pip install checkov"
        install_tool terragrunt "choco install terragrunt -y"
        ;;
    *)
        echo_message "Unsupported operating system: $OS"
        exit 1
        ;;
esac

# Create installation directory
INSTALL_DIR="/usr/local/bin"
EXECUTABLE_NAME="infratoolbox"

# Compile the project
echo_message "Compiling InfraToolbox..."
go build -o $INSTALL_DIR/$EXECUTABLE_NAME

# Confirm installation
if command -v infratoolbox &> /dev/null; then
    echo_message "InfraToolbox has been successfully installed at $INSTALL_DIR/$EXECUTABLE_NAME."
else
    echo_message "Error installing InfraToolbox."
    exit 1
fi

# Configure auto-completion
echo "Setting up shell auto-completion..."

# Bash
if [ -d "/etc/bash_completion.d" ]; then
    "$INSTALL_DIR/$EXECUTABLE_NAME" completion bash > /etc/bash_completion.d/$EXECUTABLE_NAME
    echo "Bash completion installed at /etc/bash_completion.d/$EXECUTABLE_NAME"
fi

# Zsh
if [ -d "${fpath[1]}" ]; then
    "$INSTALL_DIR/$EXECUTABLE_NAME" completion zsh > "${fpath[1]}/_$EXECUTABLE_NAME"
    echo "Zsh completion installed at ${fpath[1]}/_$EXECUTABLE_NAME"
fi

# Fish
if [ -d "$HOME/.config/fish/completions" ]; then
    "$INSTALL_DIR/$EXECUTABLE_NAME" completion fish > "$HOME/.config/fish/completions/$EXECUTABLE_NAME.fish"
    echo "Fish completion installed at $HOME/.config/fish/completions/$EXECUTABLE_NAME.fish"
fi

# PowerShell
if [ -n "$PSModulePath" ]; then
    "$INSTALL_DIR/$EXECUTABLE_NAME" completion powershell > "$HOME/$EXECUTABLE_NAME.ps1"
    echo "PowerShell completion script generated at $HOME/$EXECUTABLE_NAME.ps1"
    echo "To enable it, add the following line to your PowerShell profile:"
    echo "  . $HOME/$EXECUTABLE_NAME.ps1"
fi

# Finish
echo_message "Installation completed. You can run InfraToolbox with the command 'infratoolbox'."
