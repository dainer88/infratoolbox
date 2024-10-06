class Infratoolbox < Formule
    desc "InfraToolbox integrates Terraform, Terraform Docs, TFEnv, TFLint, and other tools into a single CLI for managing infrastructure efficiently."
    homepage "https://github.com/dainer88/infratoolbox"
    url "https://github.com/dainer88/infratoolbox/releases/download/refs/tags/v0.1.0/infratoolbox-v0.1.0.tar.gz"
    version "0.1.0"
    sha256 "bd1a27de5786a6a91a5bb569384f5d28041fa1d42a55332358b88b79e02a4d1c"

    def install
        bin.install "infratoolbox"
    end

    depends_on "terraform"
    depends_on "terraform-docs"
    depends_on "tfenv"
    depends_on "tflint"
    depends_on "checkov"
end
  