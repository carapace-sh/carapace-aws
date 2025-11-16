package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_createMetricAttributionCmd = &cobra.Command{
	Use:   "create-metric-attribution",
	Short: "Creates a metric attribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_createMetricAttributionCmd).Standalone()

	personalize_createMetricAttributionCmd.Flags().String("dataset-group-arn", "", "The Amazon Resource Name (ARN) of the destination dataset group for the metric attribution.")
	personalize_createMetricAttributionCmd.Flags().String("metrics", "", "A list of metric attributes for the metric attribution.")
	personalize_createMetricAttributionCmd.Flags().String("metrics-output-config", "", "The output configuration details for the metric attribution.")
	personalize_createMetricAttributionCmd.Flags().String("name", "", "A name for the metric attribution.")
	personalize_createMetricAttributionCmd.MarkFlagRequired("dataset-group-arn")
	personalize_createMetricAttributionCmd.MarkFlagRequired("metrics")
	personalize_createMetricAttributionCmd.MarkFlagRequired("metrics-output-config")
	personalize_createMetricAttributionCmd.MarkFlagRequired("name")
	personalizeCmd.AddCommand(personalize_createMetricAttributionCmd)
}
