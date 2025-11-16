package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_updateAnomalyMonitorCmd = &cobra.Command{
	Use:   "update-anomaly-monitor",
	Short: "Updates an existing cost anomaly monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_updateAnomalyMonitorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_updateAnomalyMonitorCmd).Standalone()

		ce_updateAnomalyMonitorCmd.Flags().String("monitor-arn", "", "Cost anomaly monitor Amazon Resource Names (ARNs).")
		ce_updateAnomalyMonitorCmd.Flags().String("monitor-name", "", "The new name for the cost anomaly monitor.")
		ce_updateAnomalyMonitorCmd.MarkFlagRequired("monitor-arn")
	})
	ceCmd.AddCommand(ce_updateAnomalyMonitorCmd)
}
