package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_registerWebhookWithThirdPartyCmd = &cobra.Command{
	Use:   "register-webhook-with-third-party",
	Short: "Configures a connection between the webhook that was created and the external tool with events to be detected.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_registerWebhookWithThirdPartyCmd).Standalone()

	codepipeline_registerWebhookWithThirdPartyCmd.Flags().String("webhook-name", "", "The name of an existing webhook created with PutWebhook to register with a supported third party.")
	codepipelineCmd.AddCommand(codepipeline_registerWebhookWithThirdPartyCmd)
}
