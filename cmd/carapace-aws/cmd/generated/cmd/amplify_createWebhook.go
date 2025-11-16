package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_createWebhookCmd = &cobra.Command{
	Use:   "create-webhook",
	Short: "Creates a new webhook on an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_createWebhookCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_createWebhookCmd).Standalone()

		amplify_createWebhookCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
		amplify_createWebhookCmd.Flags().String("branch-name", "", "The name for a branch that is part of an Amplify app.")
		amplify_createWebhookCmd.Flags().String("description", "", "The description for a webhook.")
		amplify_createWebhookCmd.MarkFlagRequired("app-id")
		amplify_createWebhookCmd.MarkFlagRequired("branch-name")
	})
	amplifyCmd.AddCommand(amplify_createWebhookCmd)
}
