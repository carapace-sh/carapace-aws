package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_deregisterWebhookWithThirdPartyCmd = &cobra.Command{
	Use:   "deregister-webhook-with-third-party",
	Short: "Removes the connection between the webhook that was created by CodePipeline and the external tool with events to be detected.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_deregisterWebhookWithThirdPartyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_deregisterWebhookWithThirdPartyCmd).Standalone()

		codepipeline_deregisterWebhookWithThirdPartyCmd.Flags().String("webhook-name", "", "The name of the webhook you want to deregister.")
	})
	codepipelineCmd.AddCommand(codepipeline_deregisterWebhookWithThirdPartyCmd)
}
