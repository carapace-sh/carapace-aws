package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_createAnomalyMonitorCmd = &cobra.Command{
	Use:   "create-anomaly-monitor",
	Short: "Creates a new cost anomaly detection monitor with the requested type and monitor specification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_createAnomalyMonitorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ce_createAnomalyMonitorCmd).Standalone()

		ce_createAnomalyMonitorCmd.Flags().String("anomaly-monitor", "", "The cost anomaly detection monitor object that you want to create.")
		ce_createAnomalyMonitorCmd.Flags().String("resource-tags", "", "An optional list of tags to associate with the specified [`AnomalyMonitor`](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_AnomalyMonitor.html) .")
		ce_createAnomalyMonitorCmd.MarkFlagRequired("anomaly-monitor")
	})
	ceCmd.AddCommand(ce_createAnomalyMonitorCmd)
}
