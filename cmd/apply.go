package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
    Use:   "apply",
    Short: "Run terraform apply or terragrunt apply",
    Run: func(cmd *cobra.Command, args []string) {
        // Ejemplo básico de ejecución de Terraform o Terragrunt
        fmt.Println("Running apply...")
        applyTerraform()
    },
}

func applyTerraform() {
    // Ejecutar terraform apply
    cmd := exec.Command("terraform", "apply", "-auto-approve")
    cmd.Stdout = cmd.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Printf("Error running terraform apply: %v\n", err)
    }
}

func init() {
    rootCmd.AddCommand(applyCmd)
}
