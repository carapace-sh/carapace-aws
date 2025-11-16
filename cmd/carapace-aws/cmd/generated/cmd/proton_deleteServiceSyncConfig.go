package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_deleteServiceSyncConfigCmd = &cobra.Command{
	Use:   "delete-service-sync-config",
	Short: "Delete the Proton Ops file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_deleteServiceSyncConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_deleteServiceSyncConfigCmd).Standalone()

		proton_deleteServiceSyncConfigCmd.Flags().String("service-name", "", "The name of the service that you want to delete the service sync configuration for.")
		proton_deleteServiceSyncConfigCmd.MarkFlagRequired("service-name")
	})
	protonCmd.AddCommand(proton_deleteServiceSyncConfigCmd)
}
