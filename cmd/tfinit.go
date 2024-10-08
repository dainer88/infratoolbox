package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
    Use:   "init",
    Short: "Run terraform init",
    DisableFlagParsing: true,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running terraform init with arguments:", args)
        runTerraformInit(args)
    },
}

func runTerraformInit(args []string) {
    cmdArgs := append([]string{"init"}, args...)
    cmd := exec.Command("terraform", cmdArgs...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Printf("Error running init: %v\n", err)
    }
}

func init() {
    rootCmd.AddCommand(initCmd)
}
