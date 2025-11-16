package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_updateWebhookCmd = &cobra.Command{
	Use:   "update-webhook",
	Short: "Updates a webhook.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_updateWebhookCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_updateWebhookCmd).Standalone()

		amplify_updateWebhookCmd.Flags().String("branch-name", "", "The name for a branch that is part of an Amplify app.")
		amplify_updateWebhookCmd.Flags().String("description", "", "The description for a webhook.")
		amplify_updateWebhookCmd.Flags().String("webhook-id", "", "The unique ID for a webhook.")
		amplify_updateWebhookCmd.MarkFlagRequired("webhook-id")
	})
	amplifyCmd.AddCommand(amplify_updateWebhookCmd)
}
