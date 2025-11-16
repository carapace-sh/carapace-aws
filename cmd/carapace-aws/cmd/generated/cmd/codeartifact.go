package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifactCmd = &cobra.Command{
	Use:   "codeartifact",
	Short: "CodeArtifact is a fully managed artifact repository compatible with language-native package managers and build tools such as npm, Apache Maven, pip, and dotnet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifactCmd).Standalone()

	rootCmd.AddCommand(codeartifactCmd)
}
