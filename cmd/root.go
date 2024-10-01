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
    CompletionOptions: cobra.CompletionOptions{ DisableDefaultCmd: true },
    DisableFlagsInUseLine: true,
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}

func Execute() {
    rootCmd.SetHelpTemplate(`
{{if .Long}}{{.Long}}
{{end}}
Usage:
  {{.UseLine}} [command] [flags]

{{if .HasAvailableSubCommands}}
Available Commands:{{range .Commands}}{{if (and .IsAvailableCommand (not .Hidden))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}
{{if .HasAvailableLocalFlags}}
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}
{{end}}
{{if .HasAvailableInheritedFlags}}Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}
{{end}}
{{if .HasHelpSubCommands}}Additional help topics:
{{range .Commands}}{{if .IsHelpCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}
{{end}}{{if .HasAvailableSubCommands}}
Use "{{.CommandPath}} [command] --help" for more information about a command.
{{end}}
`)
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
