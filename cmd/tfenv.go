package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
    Use:   "env",
    Short: "Run tfenv to manage terraform versions",
    Long:  "Run tfenv to manage terraform versions in the current shell.",
    DisableFlagParsing: true,
    Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && (args[0] == "-h" || args[0] == "--help") {
			runTfenv([]string{"-h"})
			return
		}
        fmt.Println("Running terraform apply with arguments:", args)
        runTfenv(args)
    },
}

func runTfenv(args []string) {
    cmd := exec.Command("tfenv", args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    
    if err := cmd.Run(); err != nil {
        fmt.Printf("Error running apply: %v\n", err)
    }
}

func init() {
    rootCmd.AddCommand(envCmd)
}
