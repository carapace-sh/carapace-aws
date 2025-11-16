package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_deleteSourceCredentialsCmd = &cobra.Command{
	Use:   "delete-source-credentials",
	Short: "Deletes a set of GitHub, GitHub Enterprise, or Bitbucket source credentials.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_deleteSourceCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_deleteSourceCredentialsCmd).Standalone()

		codebuild_deleteSourceCredentialsCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the token.")
		codebuild_deleteSourceCredentialsCmd.MarkFlagRequired("arn")
	})
	codebuildCmd.AddCommand(codebuild_deleteSourceCredentialsCmd)
}
