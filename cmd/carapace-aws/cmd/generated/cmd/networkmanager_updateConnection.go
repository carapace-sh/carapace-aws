package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_updateConnectionCmd = &cobra.Command{
	Use:   "update-connection",
	Short: "Updates the information for an existing connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_updateConnectionCmd).Standalone()

	networkmanager_updateConnectionCmd.Flags().String("connected-link-id", "", "The ID of the link for the second device in the connection.")
	networkmanager_updateConnectionCmd.Flags().String("connection-id", "", "The ID of the connection.")
	networkmanager_updateConnectionCmd.Flags().String("description", "", "A description of the connection.")
	networkmanager_updateConnectionCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_updateConnectionCmd.Flags().String("link-id", "", "The ID of the link for the first device in the connection.")
	networkmanager_updateConnectionCmd.MarkFlagRequired("connection-id")
	networkmanager_updateConnectionCmd.MarkFlagRequired("global-network-id")
	networkmanagerCmd.AddCommand(networkmanager_updateConnectionCmd)
}
