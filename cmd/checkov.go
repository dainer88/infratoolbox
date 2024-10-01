package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var checkovCmd = &cobra.Command{
    Use:   "scan",
    Short: "Run scan to perform security scan",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running checkov    with arguments:", args)
        runCheckov(args)
    },
}

func runCheckov(args []string) {
    cmd := exec.Command("checkov", args...)
    cmd.Stdout = cmd.Stderr
    if output, err := cmd.CombinedOutput(); err != nil {
        fmt.Printf("Error running Checkov: %v\n", err)
        fmt.Printf("Checkov output: %s\n", output)
    }
}

func init() {
    rootCmd.AddCommand(checkovCmd)
}
