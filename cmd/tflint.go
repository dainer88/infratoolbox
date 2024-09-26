package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
    Use:   "lint",
    Short: "Run TFLint to lint Terraform code",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running TFLint...")
        tflintCmd := exec.Command("tflint")
        tflintCmd.Stdout = tflintCmd.Stderr
        if output, err := tflintCmd.CombinedOutput(); err != nil {
            fmt.Printf("Error running Checkov: %v\n", err)
            fmt.Printf("Checkov output: %s\n", output)
        }
    },
}

func init() {
    rootCmd.AddCommand(lintCmd)
}
