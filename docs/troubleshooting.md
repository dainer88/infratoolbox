# Troubleshooting

## Error: "command not found"

Ensure that the `$GOPATH/bin` directory is in your `PATH` environment variable.

### Solution:

Add the following to your `.bashrc` or `.zshrc` file:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## Error: "Terragrunt is not installed"

InfraToolbox requires Terragrunt to be installed to execute Terragrunt commands.

### Solution:

Install Terragrunt by following the instructions at [terragrunt.io](https://terragrunt.io/).
