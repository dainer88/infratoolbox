package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var newModuleCmd = &cobra.Command{
    Use:   "module [module-name]",
    Short: "Create a new Terraform module",
    Long:  `Creates a new directory with the basic files for a Terraform module (main.tf, variables.tf, outputs.tf, README.md).`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        moduleName := args[0]
        createTerraformModule(moduleName)
    },
}

func createTerraformModule(moduleName string) {
    if err := os.Mkdir(moduleName, 0755); err != nil {
        fmt.Printf("Error creating module directory: %v\n", err)
        return
    }

    files := map[string]string{
        "main.tf":      `// main.tf - Terraform configuration`,
        "variables.tf": `// variables.tf - Define your variables here`,
        "outputs.tf":   `// outputs.tf - Define your outputs here`,
        "README.md":    `# %s\n\nModule description here.`, 
    }

    for fileName, content := range files {
        filePath := filepath.Join(moduleName, fileName)
        fileContent := fmt.Sprintf(content, moduleName)
        if err := os.WriteFile(filePath, []byte(fileContent), 0644); err != nil {
            fmt.Printf("Error creating file %s: %v\n", fileName, err)
        } else {
            fmt.Printf("Created file: %s\n", filePath)
        }
    }

    fmt.Printf("Terraform module '%s' created successfully.\n", moduleName)
}

func init() {
    rootCmd.AddCommand(newModuleCmd)
}
