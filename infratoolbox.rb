class Infratoolbox < Formula
  desc "InfraToolbox integrates Terraform, Terraform Docs, TFEnv, TFLint, and other tools into a single CLI for managing infrastructure efficiently."
  homepage "https://github.com/dainer88/infratoolbox"
  url "https://github.com/dainer88/infratoolbox/releases/download/v0.2.0/infratoolbox-v0.2.0-macos.tar.gz"
  version "v0.2.0"
  sha256 "c5a0e59187e90e723db4c69657cc1a0d582035430d37fc88a082646cc2ebfca0"

  depends_on "terraform-docs"
  depends_on "terragrunt"
  depends_on "tfenv"
  depends_on "tflint"
  depends_on "checkov"

  def install
    # Instalar el binario principal
    bin.install "infratoolbox"

    # Generar y copiar scripts de auto-completado
    bash_completion.install Utils.safe_popen_read(bin/"infratoolbox", "completion", "bash") => "infratoolbox"
    zsh_completion.install Utils.safe_popen_read(bin/"infratoolbox", "completion", "zsh") => "_infratoolbox"
    fish_completion.install Utils.safe_popen_read(bin/"infratoolbox", "completion", "fish") => "infratoolbox.fish"
  end

  def post_install
    # Asegurar permisos correctos para el binario
    system "chmod", "0755", bin/"infratoolbox"
  end

  test do
    # Verificar que el comando principal funciona
    assert_match "InfraToolbox", shell_output("#{bin}/infratoolbox --help")

    # Verificar que los scripts de auto-completado se generen correctamente
    assert_match "-F _infratoolbox", shell_output("bash -c 'source #{bash_completion}/infratoolbox && complete -p infratoolbox'")
  end
end
