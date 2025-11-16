package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_batchDeleteRumMetricDefinitionsCmd = &cobra.Command{
	Use:   "batch-delete-rum-metric-definitions",
	Short: "Removes the specified metrics from being sent to an extended metrics destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_batchDeleteRumMetricDefinitionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rum_batchDeleteRumMetricDefinitionsCmd).Standalone()

		rum_batchDeleteRumMetricDefinitionsCmd.Flags().String("app-monitor-name", "", "The name of the CloudWatch RUM app monitor that is sending these metrics.")
		rum_batchDeleteRumMetricDefinitionsCmd.Flags().String("destination", "", "Defines the destination where you want to stop sending the specified metrics.")
		rum_batchDeleteRumMetricDefinitionsCmd.Flags().String("destination-arn", "", "This parameter is required if `Destination` is `Evidently`.")
		rum_batchDeleteRumMetricDefinitionsCmd.Flags().String("metric-definition-ids", "", "An array of structures which define the metrics that you want to stop sending.")
		rum_batchDeleteRumMetricDefinitionsCmd.MarkFlagRequired("app-monitor-name")
		rum_batchDeleteRumMetricDefinitionsCmd.MarkFlagRequired("destination")
		rum_batchDeleteRumMetricDefinitionsCmd.MarkFlagRequired("metric-definition-ids")
	})
	rumCmd.AddCommand(rum_batchDeleteRumMetricDefinitionsCmd)
}
