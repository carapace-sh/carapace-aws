package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeMetricAttributionCmd = &cobra.Command{
	Use:   "describe-metric-attribution",
	Short: "Describes a metric attribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeMetricAttributionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_describeMetricAttributionCmd).Standalone()

		personalize_describeMetricAttributionCmd.Flags().String("metric-attribution-arn", "", "The metric attribution's Amazon Resource Name (ARN).")
		personalize_describeMetricAttributionCmd.MarkFlagRequired("metric-attribution-arn")
	})
	personalizeCmd.AddCommand(personalize_describeMetricAttributionCmd)
}
