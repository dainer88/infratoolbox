package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var checkovCmd = &cobra.Command{
    Use:   "scan",
    Short: "Run scan to perform security scan",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running scan with arguments:", args)
        runCheckov(args)
    },
}

func runCheckov(args []string) {
    cmd := exec.Command("checkov", args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Printf("Error running scan: %v\n", err)
    }
}

func init() {
    rootCmd.AddCommand(checkovCmd)
}
