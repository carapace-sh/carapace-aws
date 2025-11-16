package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_updateWebhookCmd = &cobra.Command{
	Use:   "update-webhook",
	Short: "Updates the webhook associated with an CodeBuild build project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_updateWebhookCmd).Standalone()

	codebuild_updateWebhookCmd.Flags().String("branch-filter", "", "A regular expression used to determine which repository branches are built when a webhook is triggered.")
	codebuild_updateWebhookCmd.Flags().String("build-type", "", "Specifies the type of build this webhook will trigger.")
	codebuild_updateWebhookCmd.Flags().String("filter-groups", "", "An array of arrays of `WebhookFilter` objects used to determine if a webhook event can trigger a build.")
	codebuild_updateWebhookCmd.Flags().Bool("no-rotate-secret", false, "A boolean value that specifies whether the associated GitHub repository's secret token should be updated.")
	codebuild_updateWebhookCmd.Flags().String("project-name", "", "The name of the CodeBuild project.")
	codebuild_updateWebhookCmd.Flags().String("pull-request-build-policy", "", "A PullRequestBuildPolicy object that defines comment-based approval requirements for triggering builds on pull requests.")
	codebuild_updateWebhookCmd.Flags().Bool("rotate-secret", false, "A boolean value that specifies whether the associated GitHub repository's secret token should be updated.")
	codebuild_updateWebhookCmd.Flag("no-rotate-secret").Hidden = true
	codebuild_updateWebhookCmd.MarkFlagRequired("project-name")
	codebuildCmd.AddCommand(codebuild_updateWebhookCmd)
}
