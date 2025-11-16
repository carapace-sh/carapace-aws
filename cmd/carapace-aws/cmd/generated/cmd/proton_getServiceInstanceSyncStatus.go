package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getServiceInstanceSyncStatusCmd = &cobra.Command{
	Use:   "get-service-instance-sync-status",
	Short: "Get the status of the synced service instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getServiceInstanceSyncStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_getServiceInstanceSyncStatusCmd).Standalone()

		proton_getServiceInstanceSyncStatusCmd.Flags().String("service-instance-name", "", "The name of the service instance that you want the sync status input for.")
		proton_getServiceInstanceSyncStatusCmd.Flags().String("service-name", "", "The name of the service that the service instance belongs to.")
		proton_getServiceInstanceSyncStatusCmd.MarkFlagRequired("service-instance-name")
		proton_getServiceInstanceSyncStatusCmd.MarkFlagRequired("service-name")
	})
	protonCmd.AddCommand(proton_getServiceInstanceSyncStatusCmd)
}
