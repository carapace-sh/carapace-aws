package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getInsightSummariesCmd = &cobra.Command{
	Use:   "get-insight-summaries",
	Short: "Retrieves the summaries of all insights in the specified group matching the provided filter values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getInsightSummariesCmd).Standalone()

	xray_getInsightSummariesCmd.Flags().String("end-time", "", "The end of the time frame in which the insights ended.")
	xray_getInsightSummariesCmd.Flags().String("group-arn", "", "The Amazon Resource Name (ARN) of the group.")
	xray_getInsightSummariesCmd.Flags().String("group-name", "", "The name of the group.")
	xray_getInsightSummariesCmd.Flags().String("max-results", "", "The maximum number of results to display.")
	xray_getInsightSummariesCmd.Flags().String("next-token", "", "Pagination token.")
	xray_getInsightSummariesCmd.Flags().String("start-time", "", "The beginning of the time frame in which the insights started.")
	xray_getInsightSummariesCmd.Flags().String("states", "", "The list of insight states.")
	xray_getInsightSummariesCmd.MarkFlagRequired("end-time")
	xray_getInsightSummariesCmd.MarkFlagRequired("start-time")
	xrayCmd.AddCommand(xray_getInsightSummariesCmd)
}
