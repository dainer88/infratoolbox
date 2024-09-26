package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var securityCmd = &cobra.Command{
    Use:   "security",
    Short: "Run Checkov to perform security scan",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running Checkov...")
        checkovCmd := exec.Command("checkov", "-d", ".")
        checkovCmd.Stdout = checkovCmd.Stderr
        if output, err := checkovCmd.CombinedOutput(); err != nil {
            fmt.Printf("Error running Checkov: %v\n", err)
            fmt.Printf("Checkov output: %s\n", output)
        }
    },
}

func init() {
    rootCmd.AddCommand(securityCmd)
}
