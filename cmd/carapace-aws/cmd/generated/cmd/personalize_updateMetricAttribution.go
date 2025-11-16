package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_updateMetricAttributionCmd = &cobra.Command{
	Use:   "update-metric-attribution",
	Short: "Updates a metric attribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_updateMetricAttributionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_updateMetricAttributionCmd).Standalone()

		personalize_updateMetricAttributionCmd.Flags().String("add-metrics", "", "Add new metric attributes to the metric attribution.")
		personalize_updateMetricAttributionCmd.Flags().String("metric-attribution-arn", "", "The Amazon Resource Name (ARN) for the metric attribution to update.")
		personalize_updateMetricAttributionCmd.Flags().String("metrics-output-config", "", "An output config for the metric attribution.")
		personalize_updateMetricAttributionCmd.Flags().String("remove-metrics", "", "Remove metric attributes from the metric attribution.")
	})
	personalizeCmd.AddCommand(personalize_updateMetricAttributionCmd)
}
