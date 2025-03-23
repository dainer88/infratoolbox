class Infratoolbox < Formula
    desc "InfraToolbox integrates Terraform, Terraform Docs, TFEnv, TFLint, and other tools into a single CLI for managing infrastructure efficiently."
    homepage "https://github.com/dainer88/infratoolbox"
    url "https://github.com/dainer88/infratoolbox/releases/download/v0.2.0/infratoolbox-v0.2.0-macos.tar.gz"
    version "v0.2.0"
    sha256 "c5a0e59187e90e723db4c69657cc1a0d582035430d37fc88a082646cc2ebfca0"

    def install
        bin.install "infratoolbox"
    end

    def post_install
        system "chmod", "0755", "/opt/homebrew/bin/infratoolbox"
    end

    depends_on "terraform-docs"
    depends_on "terragrunt"
    depends_on "tfenv"
    depends_on "tflint"
    depends_on "checkov"
end
