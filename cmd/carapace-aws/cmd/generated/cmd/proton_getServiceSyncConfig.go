package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getServiceSyncConfigCmd = &cobra.Command{
	Use:   "get-service-sync-config",
	Short: "Get detailed information for the service sync configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getServiceSyncConfigCmd).Standalone()

	proton_getServiceSyncConfigCmd.Flags().String("service-name", "", "The name of the service that you want to get the service sync configuration for.")
	proton_getServiceSyncConfigCmd.MarkFlagRequired("service-name")
	protonCmd.AddCommand(proton_getServiceSyncConfigCmd)
}
