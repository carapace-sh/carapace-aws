package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_putRumMetricsDestinationCmd = &cobra.Command{
	Use:   "put-rum-metrics-destination",
	Short: "Creates or updates a destination to receive extended metrics from CloudWatch RUM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_putRumMetricsDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rum_putRumMetricsDestinationCmd).Standalone()

		rum_putRumMetricsDestinationCmd.Flags().String("app-monitor-name", "", "The name of the CloudWatch RUM app monitor that will send the metrics.")
		rum_putRumMetricsDestinationCmd.Flags().String("destination", "", "Defines the destination to send the metrics to.")
		rum_putRumMetricsDestinationCmd.Flags().String("destination-arn", "", "Use this parameter only if `Destination` is `Evidently`.")
		rum_putRumMetricsDestinationCmd.Flags().String("iam-role-arn", "", "This parameter is required if `Destination` is `Evidently`.")
		rum_putRumMetricsDestinationCmd.MarkFlagRequired("app-monitor-name")
		rum_putRumMetricsDestinationCmd.MarkFlagRequired("destination")
	})
	rumCmd.AddCommand(rum_putRumMetricsDestinationCmd)
}
