package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
    Use:   "init",
    Short: "Run terraform init",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running terraform init with arguments:", args)
        runTerraformInit(args)
    },
}

func runTerraformInit(args []string) {
    cmdArgs := append([]string{"init"}, args...)
    cmd := exec.Command("terraform", cmdArgs...)
    cmd.Stdout = cmd.Stderr
    if output, err := cmd.CombinedOutput(); err != nil {
        fmt.Printf("Error running init: %v\n", err)
        fmt.Printf("Checkov output: %s\n", output)
    }
}

func init() {
    rootCmd.AddCommand(initCmd)
}
