package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getDevicesCmd = &cobra.Command{
	Use:   "get-devices",
	Short: "Gets information about one or more of your devices in a global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getDevicesCmd).Standalone()

	networkmanager_getDevicesCmd.Flags().String("device-ids", "", "One or more device IDs.")
	networkmanager_getDevicesCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_getDevicesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	networkmanager_getDevicesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	networkmanager_getDevicesCmd.Flags().String("site-id", "", "The ID of the site.")
	networkmanager_getDevicesCmd.MarkFlagRequired("global-network-id")
	networkmanagerCmd.AddCommand(networkmanager_getDevicesCmd)
}
