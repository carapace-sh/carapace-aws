package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_deleteRumMetricsDestinationCmd = &cobra.Command{
	Use:   "delete-rum-metrics-destination",
	Short: "Deletes a destination for CloudWatch RUM extended metrics, so that the specified app monitor stops sending extended metrics to that destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_deleteRumMetricsDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rum_deleteRumMetricsDestinationCmd).Standalone()

		rum_deleteRumMetricsDestinationCmd.Flags().String("app-monitor-name", "", "The name of the app monitor that is sending metrics to the destination that you want to delete.")
		rum_deleteRumMetricsDestinationCmd.Flags().String("destination", "", "The type of destination to delete.")
		rum_deleteRumMetricsDestinationCmd.Flags().String("destination-arn", "", "This parameter is required if `Destination` is `Evidently`.")
		rum_deleteRumMetricsDestinationCmd.MarkFlagRequired("app-monitor-name")
		rum_deleteRumMetricsDestinationCmd.MarkFlagRequired("destination")
	})
	rumCmd.AddCommand(rum_deleteRumMetricsDestinationCmd)
}
