package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listMetricAttributionsCmd = &cobra.Command{
	Use:   "list-metric-attributions",
	Short: "Lists metric attributions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listMetricAttributionsCmd).Standalone()

	personalize_listMetricAttributionsCmd.Flags().String("dataset-group-arn", "", "The metric attributions' dataset group Amazon Resource Name (ARN).")
	personalize_listMetricAttributionsCmd.Flags().String("max-results", "", "The maximum number of metric attributions to return in one page of results.")
	personalize_listMetricAttributionsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	personalizeCmd.AddCommand(personalize_listMetricAttributionsCmd)
}
