package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_deleteWebhookCmd = &cobra.Command{
	Use:   "delete-webhook",
	Short: "For an existing CodeBuild build project that has its source code stored in a GitHub or Bitbucket repository, stops CodeBuild from rebuilding the source code every time a code change is pushed to the repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_deleteWebhookCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_deleteWebhookCmd).Standalone()

		codebuild_deleteWebhookCmd.Flags().String("project-name", "", "The name of the CodeBuild project.")
		codebuild_deleteWebhookCmd.MarkFlagRequired("project-name")
	})
	codebuildCmd.AddCommand(codebuild_deleteWebhookCmd)
}
