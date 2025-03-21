#!/bin/bash

# Installation script for InfraToolbox

# Function to print messages
function echo_message {
    echo "=============================="
    echo "$1"
    echo "=============================="
}

# Check if Homebrew is installed
if ! command -v brew &> /dev/null; then
    echo_message "Homebrew is not installed. Please install it first."
    exit 1
fi

# Function to install a tool if it's not already installed
function install_tool {
    TOOL_NAME=$1
    if command -v $TOOL_NAME &> /dev/null; then
        echo_message "$TOOL_NAME is already installed. Skipping installation."
    else
        echo_message "Installing $TOOL_NAME..."
        brew install $TOOL_NAME
    fi
}

# Install necessary tools using Homebrew
install_tool terraform-docs
install_tool tfenv
install_tool tflint
install_tool checkov
install_tool terragrunt

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

# Configurar auto-completado
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
