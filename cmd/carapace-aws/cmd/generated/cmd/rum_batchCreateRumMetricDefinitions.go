package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_batchCreateRumMetricDefinitionsCmd = &cobra.Command{
	Use:   "batch-create-rum-metric-definitions",
	Short: "Specifies the extended metrics and custom metrics that you want a CloudWatch RUM app monitor to send to a destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_batchCreateRumMetricDefinitionsCmd).Standalone()

	rum_batchCreateRumMetricDefinitionsCmd.Flags().String("app-monitor-name", "", "The name of the CloudWatch RUM app monitor that is to send the metrics.")
	rum_batchCreateRumMetricDefinitionsCmd.Flags().String("destination", "", "The destination to send the metrics to.")
	rum_batchCreateRumMetricDefinitionsCmd.Flags().String("destination-arn", "", "This parameter is required if `Destination` is `Evidently`.")
	rum_batchCreateRumMetricDefinitionsCmd.Flags().String("metric-definitions", "", "An array of structures which define the metrics that you want to send.")
	rum_batchCreateRumMetricDefinitionsCmd.MarkFlagRequired("app-monitor-name")
	rum_batchCreateRumMetricDefinitionsCmd.MarkFlagRequired("destination")
	rum_batchCreateRumMetricDefinitionsCmd.MarkFlagRequired("metric-definitions")
	rumCmd.AddCommand(rum_batchCreateRumMetricDefinitionsCmd)
}
