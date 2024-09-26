package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// planCmd representa el comando "plan"
var planCmd = &cobra.Command{
    Use:   "plan",
    Short: "Run terraform plan",
    Long:  "Generate and show an execution plan for Terraform.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running terraform plan...")
        runTerraformPlan()
    },
}

// runTerraformPlan ejecuta el comando "terraform plan"
func runTerraformPlan() {
    cmd := exec.Command("terraform", "plan")
    cmd.Stdout = cmd.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Printf("Error running terraform plan: %v\n", err)
    }
}

func init() {
    rootCmd.AddCommand(planCmd) // Agrega el comando "plan" al comando raíz
}
