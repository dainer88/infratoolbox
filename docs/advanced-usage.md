# Advanced Use Cases

## Using Terragrunt

You can run Terragrunt commands directly from InfraToolbox:

```bash
infratoolbox terragrunt plan --terragrunt-config=terragrunt.hcl
```

## Customizing Pre-built Modules

After generating a pre-built module, you can customize the generated files to suit your needs:

```bash
infratoolbox module my-custom-s3-module
cd my-custom-s3-module
# Edit the main.tf, variables.tf, etc. files.
```

## CI/CD Integration

Example of a pipeline in GitHub Actions:

```yaml
name: Terraform Plan

on:
  push:
    branches:
      - main

jobs:
  terraform:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v2

      - name: Install InfraToolbox
        run: |
          go install github.com/yourusername/infratoolbox@latest

      - name: Run Terraform Plan
        run: |
          infratoolbox terraform plan
```
