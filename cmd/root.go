package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "infratoolbox",
    Short: "InfraToolbox is a unified CLI for infrastructure management",
    Long:  "InfraToolbox integrates Terraform, Terragrunt, TFLint, and other tools into a single CLI for managing infrastructure efficiently.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to InfraToolbox!")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
