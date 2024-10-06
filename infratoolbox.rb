class Infratoolbox < Formule
    desc "InfraToolbox integrates Terraform, Terraform Docs, TFEnv, TFLint, and other tools into a single CLI for managing infrastructure efficiently."
    homepage "https://github.com/dainer88/infratoolbox"
    url "https://github.com/yourusername/yourrepo/releases/download/refs/tags/v0.1.0/infratoolbox.tar.gz"
    version "0.1.0"
    sha256 "d8acae6185ed4ca4a58a7c8275edf6261fd3e3d527d439102c018d2a7bb2175e"

    def install
        bin.install "infratoolbox"
    end

    depends_on "terraform"
    depends_on "terraform-docs"
    depends_on "tfenv"
    depends_on "tflint"
    depends_on "checkov"
end
  