package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getServiceSyncBlockerSummaryCmd = &cobra.Command{
	Use:   "get-service-sync-blocker-summary",
	Short: "Get detailed data for the service sync blocker summary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getServiceSyncBlockerSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_getServiceSyncBlockerSummaryCmd).Standalone()

		proton_getServiceSyncBlockerSummaryCmd.Flags().String("service-instance-name", "", "The name of the service instance that you want to get the service sync blocker summary for.")
		proton_getServiceSyncBlockerSummaryCmd.Flags().String("service-name", "", "The name of the service that you want to get the service sync blocker summary for.")
		proton_getServiceSyncBlockerSummaryCmd.MarkFlagRequired("service-name")
	})
	protonCmd.AddCommand(proton_getServiceSyncBlockerSummaryCmd)
}
