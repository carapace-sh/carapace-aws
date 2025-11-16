package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_listSourceCredentialsCmd = &cobra.Command{
	Use:   "list-source-credentials",
	Short: "Returns a list of `SourceCredentialsInfo` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_listSourceCredentialsCmd).Standalone()

	codebuildCmd.AddCommand(codebuild_listSourceCredentialsCmd)
}
