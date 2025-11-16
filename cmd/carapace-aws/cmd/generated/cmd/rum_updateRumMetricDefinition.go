package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_updateRumMetricDefinitionCmd = &cobra.Command{
	Use:   "update-rum-metric-definition",
	Short: "Modifies one existing metric definition for CloudWatch RUM extended metrics.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_updateRumMetricDefinitionCmd).Standalone()

	rum_updateRumMetricDefinitionCmd.Flags().String("app-monitor-name", "", "The name of the CloudWatch RUM app monitor that sends these metrics.")
	rum_updateRumMetricDefinitionCmd.Flags().String("destination", "", "The destination to send the metrics to.")
	rum_updateRumMetricDefinitionCmd.Flags().String("destination-arn", "", "This parameter is required if `Destination` is `Evidently`.")
	rum_updateRumMetricDefinitionCmd.Flags().String("metric-definition", "", "A structure that contains the new definition that you want to use for this metric.")
	rum_updateRumMetricDefinitionCmd.Flags().String("metric-definition-id", "", "The ID of the metric definition to update.")
	rum_updateRumMetricDefinitionCmd.MarkFlagRequired("app-monitor-name")
	rum_updateRumMetricDefinitionCmd.MarkFlagRequired("destination")
	rum_updateRumMetricDefinitionCmd.MarkFlagRequired("metric-definition")
	rum_updateRumMetricDefinitionCmd.MarkFlagRequired("metric-definition-id")
	rumCmd.AddCommand(rum_updateRumMetricDefinitionCmd)
}
