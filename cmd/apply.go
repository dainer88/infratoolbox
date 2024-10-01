package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
    Use:   "apply",
    Short: "Run terraform apply",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running terraform apply with arguments:", args)
        runTerraformApply(args)
    },
}

func runTerraformApply(args []string) {
    cmdArgs := append([]string{"apply"}, args...)
    cmd := exec.Command("terraform", cmdArgs...)
    cmd.Stdout = cmd.Stderr
    if output, err := cmd.CombinedOutput(); err != nil {
        fmt.Printf("Error running Checkov: %v\n", err)
        fmt.Printf("Checkov output: %s\n", output)
    }
}

func init() {
    rootCmd.AddCommand(applyCmd)
}
