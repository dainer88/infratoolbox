class Infratoolbox < Formule
    desc "InfraToolbox integrates Terraform, Terraform Docs, TFEnv, TFLint, and other tools into a single CLI for managing infrastructure efficiently."
    homepage "https://github.com/dainer88/infratoolbox"
    url "https://github.com/dainer88/infratoolbox/releases/download/refs/tags/v0.1.0/infratoolbox-v0.1.0.tar.gz"
    version "0.1.0"
    sha256 "e4be7c6ce79edd9e34a27400f74d8a699104d18a1dd50ee77be96ae008777550"

    def install
        bin.install "infratoolbox"
    end

    depends_on "terraform"
    depends_on "terraform-docs"
    depends_on "tfenv"
    depends_on "tflint"
    depends_on "checkov"
end
  