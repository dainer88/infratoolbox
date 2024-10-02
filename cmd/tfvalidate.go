package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
    Use:   "validate",
    Short: "Run terraform validate",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running terraform validate with arguments:", args)
        runTerraformVerify(args)
    },
}

func runTerraformVerify(args []string) {
    cmdArgs := append([]string{"validate"}, args...)
    cmd := exec.Command("terraform", cmdArgs...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    
    if err := cmd.Run(); err != nil {
        fmt.Printf("Error running validate: %v\n", err)
    }
}

func init() {
    rootCmd.AddCommand(verifyCmd)
}
