package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var tfdocCmd = &cobra.Command{
    Use:   "doc",
    Short: "Run terraform doc",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running terraform doc with arguments:", args)
        runTerraformDoc(args)
    },
}

func runTerraformDoc(args []string) {
    cmdArgs := append([]string{"markdown . > README.md"}, args...)
    cmd := exec.Command("terraform-docs", cmdArgs...)
    cmd.Stdout = cmd.Stderr
    if output, err := cmd.CombinedOutput(); err != nil {
        fmt.Printf("Error running init: %v\n", err)
        fmt.Printf("Checkov output: %s\n", output)
    }
}

func init() {
    rootCmd.AddCommand(tfdocCmd)
}
