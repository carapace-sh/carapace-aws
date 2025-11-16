package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_deleteAnomalyMonitorCmd = &cobra.Command{
	Use:   "delete-anomaly-monitor",
	Short: "Deletes a cost anomaly monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_deleteAnomalyMonitorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_deleteAnomalyMonitorCmd).Standalone()

		ce_deleteAnomalyMonitorCmd.Flags().String("monitor-arn", "", "The unique identifier of the cost anomaly monitor that you want to delete.")
		ce_deleteAnomalyMonitorCmd.MarkFlagRequired("monitor-arn")
	})
	ceCmd.AddCommand(ce_deleteAnomalyMonitorCmd)
}
