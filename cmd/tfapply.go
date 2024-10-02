package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
    Use:   "apply",
    Short: "Run terraform apply",
    Long:  "Run terraform apply to apply the changes required to reach the desired state of the configuration.",
    DisableFlagParsing: true,
    Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && (args[0] == "-h" || args[0] == "--help") {
			runTerraformApply([]string{"-h"})
			return
		}
        fmt.Println("Running terraform apply with arguments:", args)
        runTerraformApply(args)
    },
}

func runTerraformApply(args []string) {
    cmdArgs := append([]string{"apply"}, args...)
    cmd := exec.Command("terraform", cmdArgs...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    
    if err := cmd.Run(); err != nil {
        fmt.Printf("Error running apply: %v\n", err)
    }
}

func init() {
    rootCmd.AddCommand(applyCmd)
}
