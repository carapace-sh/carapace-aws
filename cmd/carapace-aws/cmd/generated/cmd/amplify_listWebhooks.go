package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_listWebhooksCmd = &cobra.Command{
	Use:   "list-webhooks",
	Short: "Returns a list of webhooks for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_listWebhooksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_listWebhooksCmd).Standalone()

		amplify_listWebhooksCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
		amplify_listWebhooksCmd.Flags().String("max-results", "", "The maximum number of records to list in a single response.")
		amplify_listWebhooksCmd.Flags().String("next-token", "", "A pagination token.")
		amplify_listWebhooksCmd.MarkFlagRequired("app-id")
	})
	amplifyCmd.AddCommand(amplify_listWebhooksCmd)
}
