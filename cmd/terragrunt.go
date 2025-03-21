package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var terragruntCmd = &cobra.Command{
    Use:   "terragrunt [command]",
    Short: "Run Terragrunt commands",
    Long: `Run Terragrunt commands seamlessly within InfraToolbox.
		Examples:
		infratoolbox terragrunt init
		infratoolbox terragrunt plan
		infratoolbox terragrunt apply`,
    DisableFlagParsing: true, // Permite pasar flags directamente a Terragrunt
    Args:                cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running Terragrunt command:", args)

        // Ejecutar el comando de Terragrunt
        if err := runTerragruntCommand(args); err != nil {
            fmt.Printf("Error running Terragrunt command: %v\n", err)
            os.Exit(1)
        }
    },
}

func runTerragruntCommand(args []string) error {
    // Crear el comando Terragrunt
    cmd := exec.Command("terragrunt", args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // Ejecutar el comando
    return cmd.Run()
}

func init() {
    rootCmd.AddCommand(terragruntCmd)
}