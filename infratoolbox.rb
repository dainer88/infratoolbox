class Infratoolbox < Formule
    desc "InfraToolbox integrates Terraform, Terraform Docs, TFEnv, TFLint, and other tools into a single CLI for managing infrastructure efficiently."
    homepage "https://github.com/dainer88/infratoolbox"
    url "https://github.com/dainer88/infratoolbox/releases/download/refs/tags/v0.1.0/infratoolbox-v0.1.0.tar.gz"
    version "0.1.0"
    sha256 "91d5d18aa5ca9c4803935279d05aef834a6b31a6a22a1e5ed03294aa5f90911e"

    def install
        bin.install "infratoolbox"
    end

    depends_on "terraform"
    depends_on "terraform-docs"
    depends_on "tfenv"
    depends_on "tflint"
    depends_on "checkov"
end
  