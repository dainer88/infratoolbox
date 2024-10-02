package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var tfdocCmd = &cobra.Command{
    Use:   "doc",
    Short: "Run terraform doc",
    Run: func(cmd *cobra.Command, args []string) {
        runTerraformDoc(args)
    },
}

func runTerraformDoc(args []string) {
    cmdArgs := append([]string{"markdown", "."}, args...)
    cmd := exec.Command("terraform-docs", cmdArgs...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Printf("Error running doc: %v\n", err)
    }
}

func init() {
    rootCmd.AddCommand(tfdocCmd)
}
