package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_deleteMetricAttributionCmd = &cobra.Command{
	Use:   "delete-metric-attribution",
	Short: "Deletes a metric attribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_deleteMetricAttributionCmd).Standalone()

	personalize_deleteMetricAttributionCmd.Flags().String("metric-attribution-arn", "", "The metric attribution's Amazon Resource Name (ARN).")
	personalize_deleteMetricAttributionCmd.MarkFlagRequired("metric-attribution-arn")
	personalizeCmd.AddCommand(personalize_deleteMetricAttributionCmd)
}
