package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getResourceEvaluationSummaryCmd = &cobra.Command{
	Use:   "get-resource-evaluation-summary",
	Short: "Returns a summary of resource evaluation for the specified resource evaluation ID from the proactive rules that were run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getResourceEvaluationSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_getResourceEvaluationSummaryCmd).Standalone()

		config_getResourceEvaluationSummaryCmd.Flags().String("resource-evaluation-id", "", "The unique `ResourceEvaluationId` of Amazon Web Services resource execution for which you want to retrieve the evaluation summary.")
		config_getResourceEvaluationSummaryCmd.MarkFlagRequired("resource-evaluation-id")
	})
	configCmd.AddCommand(config_getResourceEvaluationSummaryCmd)
}
