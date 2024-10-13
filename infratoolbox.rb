class Infratoolbox < Formula
    desc "InfraToolbox integrates Terraform, Terraform Docs, TFEnv, TFLint, and other tools into a single CLI for managing infrastructure efficiently."
    homepage "https://github.com/dainer88/infratoolbox"
    url "https://github.com/dainer88/infratoolbox/releases/download/v0.1.0/infratoolbox-v0.1.0.tar.gz"
    version "0.1.0"
    sha256 "a6ce7dc51a530b04f09cff2f0ea13f8a383ec5bb77a28f9ccedf7536f6e43445"

    def install
        bin.install "infratoolbox"
    end

    depends_on "terraform-docs"
    depends_on "tfenv"
    depends_on "tflint"
    depends_on "checkov"
end
