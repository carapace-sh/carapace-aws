package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listClusterAlertsCmd = &cobra.Command{
	Use:   "list-cluster-alerts",
	Short: "List the alerts for a cluster with optional filtering based on alert state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listClusterAlertsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_listClusterAlertsCmd).Standalone()

		medialive_listClusterAlertsCmd.Flags().String("cluster-id", "", "The unique ID of the cluster")
		medialive_listClusterAlertsCmd.Flags().String("max-results", "", "The maximum number of items to return")
		medialive_listClusterAlertsCmd.Flags().String("next-token", "", "The next pagination token")
		medialive_listClusterAlertsCmd.Flags().String("state-filter", "", "Specifies the set of alerts to return based on their state.")
		medialive_listClusterAlertsCmd.MarkFlagRequired("cluster-id")
	})
	medialiveCmd.AddCommand(medialive_listClusterAlertsCmd)
}
