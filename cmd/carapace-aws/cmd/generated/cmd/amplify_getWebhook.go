package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_getWebhookCmd = &cobra.Command{
	Use:   "get-webhook",
	Short: "Returns the webhook information that corresponds to a specified webhook ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_getWebhookCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_getWebhookCmd).Standalone()

		amplify_getWebhookCmd.Flags().String("webhook-id", "", "The unique ID for a webhook.")
		amplify_getWebhookCmd.MarkFlagRequired("webhook-id")
	})
	amplifyCmd.AddCommand(amplify_getWebhookCmd)
}
