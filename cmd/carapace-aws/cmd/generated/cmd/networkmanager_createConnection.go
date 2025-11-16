package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_createConnectionCmd = &cobra.Command{
	Use:   "create-connection",
	Short: "Creates a connection between two devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_createConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_createConnectionCmd).Standalone()

		networkmanager_createConnectionCmd.Flags().String("connected-device-id", "", "The ID of the second device in the connection.")
		networkmanager_createConnectionCmd.Flags().String("connected-link-id", "", "The ID of the link for the second device.")
		networkmanager_createConnectionCmd.Flags().String("description", "", "A description of the connection.")
		networkmanager_createConnectionCmd.Flags().String("device-id", "", "The ID of the first device in the connection.")
		networkmanager_createConnectionCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_createConnectionCmd.Flags().String("link-id", "", "The ID of the link for the first device.")
		networkmanager_createConnectionCmd.Flags().String("tags", "", "The tags to apply to the resource during creation.")
		networkmanager_createConnectionCmd.MarkFlagRequired("connected-device-id")
		networkmanager_createConnectionCmd.MarkFlagRequired("device-id")
		networkmanager_createConnectionCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_createConnectionCmd)
}
