package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
    Use:   "lint",
    Short: "Run TFLint to lint Terraform code",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running lint with arguments:", args)
        runTfLint(args)
    },
}

func runTfLint(args []string) {
    fmt.Println("Running TFLint...")
    cmd := exec.Command("tflint", args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Printf("Error running lint: %v\n", err)
    }
}

func init() {
    rootCmd.AddCommand(lintCmd)
}
