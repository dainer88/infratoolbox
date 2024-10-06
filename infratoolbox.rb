class Infratoolbox < Formule
    desc "InfraToolbox integrates Terraform, Terraform Docs, TFEnv, TFLint, and other tools into a single CLI for managing infrastructure efficiently."
    homepage "https://github.com/dainer88/infratoolbox"
    url "https://github.com/dainer88/infratoolbox/releases/download/v0.1.0/infratoolbox-0.1.0.tar.gz"
    version "0.1.0"
    sha256 "SHA256_DEL_ARCHIVO_TAR"

    def install
        bin.install "infratoolbox"
    end

    depends_on "terraform"
    depends_on "terraform-docs"
    depends_on "tfenv"
    depends_on "tflint"
    depends_on "checkov"
end
  