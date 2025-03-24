class Infratoolbox < Formula
  desc "InfraToolbox integrates Terraform, Terraform Docs, TFEnv, TFLint, and other tools into a single CLI for managing infrastructure efficiently."
  homepage "https://github.com/dainer88/infratoolbox"
  url "https://github.com/dainer88/infratoolbox/releases/download/v0.2.0/infratoolbox-v0.2.0-macos.tar.gz"
  version "v0.2.0"
  sha256 "f5dfb34ed9ec9bbf13493a2629095a11a9ec20bdf758078c20ff4ad702753bba"

  depends_on "terraform-docs"
  depends_on "terragrunt"
  depends_on "tfenv"
  depends_on "tflint"
  depends_on "checkov"

  def install
    bin.install "infratoolbox"

    (bash_completion/"infratoolbox").write <<~EOS
      _infratoolbox() {
        local cur prev words cword
        _init_completion || return

        COMPREPLY=( $(compgen -W "$(infratoolbox --help | awk '/^  [a-z]/ {print $1}')" -- "$cur") )
        return 0
      }
      complete -F _infratoolbox infratoolbox
    EOS

    (zsh_completion/"_infratoolbox").write <<~EOS
      #compdef infratoolbox

      _infratoolbox() {
        local -a commands
        commands=(
          ${(f)"$(infratoolbox --help | awk '/^  [a-z]/ {print $1}')"}
        )
        _describe 'command' commands
      }
    EOS

    (fish_completion/"infratoolbox.fish").write <<~EOS
      function __fish_infratoolbox_complete
        set -l commands (infratoolbox --help | awk '/^  [a-z]/ {print $1}')
        for cmd in $commands
          echo $cmd
        end
      end

      complete -c infratoolbox -f -a "(__fish_infratoolbox_complete)"
    EOS
  end

  def post_install
    system "chmod", "0755", bin/"infratoolbox"
  end

  test do
    assert_match "InfraToolbox", shell_output("#{bin}/infratoolbox --help")

    assert_predicate bash_completion/"infratoolbox", :exist?
    assert_predicate zsh_completion/"_infratoolbox", :exist?
    assert_predicate fish_completion/"infratoolbox.fish", :exist?
  end
end
