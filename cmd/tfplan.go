package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
    Use:   "plan",
    Short: "Run terraform plan",
    Long:  "Generate and show an execution plan for Terraform.",
    DisableFlagParsing: true,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running terraform plan with arguments:", args)
        runTerraformPlan(args)
    },
}

func runTerraformPlan(args []string) {
    cmdArgs := append([]string{"plan"}, args...)
    cmd := exec.Command("terraform", cmdArgs...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Printf("Error running plan: %v\n", err)
    }
}

func init() {
    rootCmd.AddCommand(planCmd)
}
