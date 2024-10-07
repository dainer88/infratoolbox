class Infratoolbox < Formule
    desc "InfraToolbox integrates Terraform, Terraform Docs, TFEnv, TFLint, and other tools into a single CLI for managing infrastructure efficiently."
    homepage "https://github.com/dainer88/infratoolbox"
    url "https://github.com/dainer88/infratoolbox/releases/download/refs/tags/v0.1.0/infratoolbox-v0.1.0.tar.gz"
    version "0.1.0"
    sha256 "235e67da3c0eae876fe1f0296e6a5f92c4abcdf5b826d7be7e598cf5739bb738"

    def install
        bin.install "infratoolbox"
    end

    depends_on "terraform"
    depends_on "terraform-docs"
    depends_on "tfenv"
    depends_on "tflint"
    depends_on "checkov"
end
  