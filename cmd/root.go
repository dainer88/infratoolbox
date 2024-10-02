package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
    version = "0.1.0"
)

var rootCmd = &cobra.Command{
    Use:   "infratoolbox",
    Short: "InfraToolbox is a unified CLI for infrastructure management",
    Long:  "InfraToolbox integrates Terraform, Terraform Docs, TFLint, Checkov, and other tools into a single CLI for managing infrastructure efficiently.",
    CompletionOptions: cobra.CompletionOptions{ DisableDefaultCmd: true },
    Run: func(cmd *cobra.Command, args []string) {
        if cmd.Flags().Changed("version") {
            fmt.Printf("InfraToolbox version: %s\n", version)
        } else {
            cmd.Help()
        }
    },
}

func Execute() {
    rootCmd.Flags().BoolP("version", "v", false, "Print the version number")
    rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
