package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_createWebhookCmd = &cobra.Command{
	Use:   "create-webhook",
	Short: "For an existing CodeBuild build project that has its source code stored in a GitHub or Bitbucket repository, enables CodeBuild to start rebuilding the source code every time a code change is pushed to the repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_createWebhookCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_createWebhookCmd).Standalone()

		codebuild_createWebhookCmd.Flags().String("branch-filter", "", "A regular expression used to determine which repository branches are built when a webhook is triggered.")
		codebuild_createWebhookCmd.Flags().String("build-type", "", "Specifies the type of build this webhook will trigger.")
		codebuild_createWebhookCmd.Flags().String("filter-groups", "", "An array of arrays of `WebhookFilter` objects used to determine which webhooks are triggered.")
		codebuild_createWebhookCmd.Flags().String("manual-creation", "", "If manualCreation is true, CodeBuild doesn't create a webhook in GitHub and instead returns `payloadUrl` and `secret` values for the webhook.")
		codebuild_createWebhookCmd.Flags().String("project-name", "", "The name of the CodeBuild project.")
		codebuild_createWebhookCmd.Flags().String("pull-request-build-policy", "", "A PullRequestBuildPolicy object that defines comment-based approval requirements for triggering builds on pull requests.")
		codebuild_createWebhookCmd.Flags().String("scope-configuration", "", "The scope configuration for global or organization webhooks.")
		codebuild_createWebhookCmd.MarkFlagRequired("project-name")
	})
	codebuildCmd.AddCommand(codebuild_createWebhookCmd)
}
