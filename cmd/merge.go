package cmd

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge terraform files",
	Long:  "Merge terraform files into one file: infratoolbox merge file1.tf file2.tf --output merged.tf",
	Args: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("help") {
			return nil
		}
		if len(args) < 2 {
			return fmt.Errorf("requires at least 2 arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Merging terraform files")
		output, _ := cmd.Flags().GetString("output")

		if err := mergeTerraformFiles(args, output); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func mergeTerraformFiles(files []string, output string) error {
	mergedFile := hclwrite.NewEmptyFile()
	mergedBody := mergedFile.Body()

	seenBlocks := make(map[string]bool)

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("error reading file %s: %w", file, err)
		}

		// Asegurarse de que el contenido termine con una línea en blanco
		if len(content) > 0 && content[len(content)-1] != '\n' {
			content = append(content, '\n')
		}

		hclFile, diags := hclwrite.ParseConfig(content, file, hcl.Pos{Line: 1, Column: 1})
		if diags.HasErrors() {
			return fmt.Errorf("error parsing file %s: %v", file, diags)
		}

		for _, block := range hclFile.Body().Blocks() {
			blockKey := fmt.Sprintf("%s.%s", block.Type(), block.Labels())

			if !seenBlocks[blockKey] {
				mergedBody.AppendBlock(block)
				seenBlocks[blockKey] = true
			}
		}
	}

	if output == "" {
		fmt.Println(string(mergedFile.Bytes()))
		return nil
	}

	err := os.WriteFile(output, mergedFile.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("error writing merged file: %w", err)
	}

	fmt.Println("Merge completed:", output)
	return nil
}

func init() {
	rootCmd.AddCommand(mergeCmd)
	mergeCmd.Flags().StringP("output", "o", "", "Output file if not provided the merged file will be printed to stdout")
}
