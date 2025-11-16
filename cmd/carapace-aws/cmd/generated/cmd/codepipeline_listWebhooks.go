package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_listWebhooksCmd = &cobra.Command{
	Use:   "list-webhooks",
	Short: "Gets a listing of all the webhooks in this Amazon Web Services Region for this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_listWebhooksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_listWebhooksCmd).Standalone()

		codepipeline_listWebhooksCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		codepipeline_listWebhooksCmd.Flags().String("next-token", "", "The token that was returned from the previous ListWebhooks call, which can be used to return the next set of webhooks in the list.")
	})
	codepipelineCmd.AddCommand(codepipeline_listWebhooksCmd)
}
