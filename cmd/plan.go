package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
    Use:   "plan",
    Short: "Run terraform plan",
    Long:  "Generate and show an execution plan for Terraform.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running terraform plan with arguments:", args)
        runTerraformPlan(args)
    },
}

func runTerraformPlan(args []string) {
    cmdArgs := append([]string{"plan"}, args...)
    cmd := exec.Command("terraform", cmdArgs...)
    cmd.Stdout = cmd.Stderr
    if output, err := cmd.CombinedOutput(); err != nil {
        fmt.Printf("Error running Checkov: %v\n", err)
        fmt.Printf("Checkov output: %s\n", output)
    }
}

func init() {
    rootCmd.AddCommand(planCmd)
}
