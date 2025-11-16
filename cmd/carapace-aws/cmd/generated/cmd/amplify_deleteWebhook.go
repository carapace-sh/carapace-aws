package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_deleteWebhookCmd = &cobra.Command{
	Use:   "delete-webhook",
	Short: "Deletes a webhook.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_deleteWebhookCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_deleteWebhookCmd).Standalone()

		amplify_deleteWebhookCmd.Flags().String("webhook-id", "", "The unique ID for a webhook.")
		amplify_deleteWebhookCmd.MarkFlagRequired("webhook-id")
	})
	amplifyCmd.AddCommand(amplify_deleteWebhookCmd)
}
