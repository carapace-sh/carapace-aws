package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_importSourceCredentialsCmd = &cobra.Command{
	Use:   "import-source-credentials",
	Short: "Imports the source repository credentials for an CodeBuild project that has its source code stored in a GitHub, GitHub Enterprise, GitLab, GitLab Self Managed, or Bitbucket repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_importSourceCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_importSourceCredentialsCmd).Standalone()

		codebuild_importSourceCredentialsCmd.Flags().String("auth-type", "", "The type of authentication used to connect to a GitHub, GitHub Enterprise, GitLab, GitLab Self Managed, or Bitbucket repository.")
		codebuild_importSourceCredentialsCmd.Flags().String("server-type", "", "The source provider used for this project.")
		codebuild_importSourceCredentialsCmd.Flags().String("should-overwrite", "", "Set to `false` to prevent overwriting the repository source credentials.")
		codebuild_importSourceCredentialsCmd.Flags().String("token", "", "For GitHub or GitHub Enterprise, this is the personal access token.")
		codebuild_importSourceCredentialsCmd.Flags().String("username", "", "The Bitbucket username when the `authType` is BASIC\\_AUTH.")
		codebuild_importSourceCredentialsCmd.MarkFlagRequired("auth-type")
		codebuild_importSourceCredentialsCmd.MarkFlagRequired("server-type")
		codebuild_importSourceCredentialsCmd.MarkFlagRequired("token")
	})
	codebuildCmd.AddCommand(codebuild_importSourceCredentialsCmd)
}
