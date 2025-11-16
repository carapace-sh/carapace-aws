package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_deleteWebhookCmd = &cobra.Command{
	Use:   "delete-webhook",
	Short: "Deletes a previously created webhook by name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_deleteWebhookCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_deleteWebhookCmd).Standalone()

		codepipeline_deleteWebhookCmd.Flags().String("name", "", "The name of the webhook you want to delete.")
		codepipeline_deleteWebhookCmd.MarkFlagRequired("name")
	})
	codepipelineCmd.AddCommand(codepipeline_deleteWebhookCmd)
}
