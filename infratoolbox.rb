class Infratoolbox < Formula
    desc "InfraToolbox integrates Terraform, Terraform Docs, TFEnv, TFLint, and other tools into a single CLI for managing infrastructure efficiently."
    homepage "https://github.com/dainer88/infratoolbox"
    url "https://github.com/dainer88/infratoolbox/releases/download/v0.1.0/infratoolbox-v0.1.0.tar.gz"
    version "v0.1.0"
    sha256 "a1a0ce23d0b39ea85c2c459cc3165d5d5f159bd6dc056f1c0d89f17c1a8e4aaa"

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
