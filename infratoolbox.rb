class Infratoolbox < Formula
    desc "InfraToolbox integrates Terraform, Terraform Docs, TFEnv, TFLint, and other tools into a single CLI for managing infrastructure efficiently."
    homepage "https://github.com/dainer88/infratoolbox"
    url "https://github.com/dainer88/infratoolbox/releases/download/v0.1.0/infratoolbox-v0.1.0-macos.tar.gz"
    version "v0.1.0"
    sha256 "bdc6b2066e9ca7dd5ce01aeb15b2b73d6cacc3d5f9c29615a0355543852f5a25"

    def install
        bin.install "infratoolbox"
    end

    def post_install
        system "chmod", "0755", "/opt/homebrew/bin/infratoolbox"
    end

    depends_on "terraform-docs"
    depends_on "tfenv"
    depends_on "tflint"
    depends_on "checkov"
end
