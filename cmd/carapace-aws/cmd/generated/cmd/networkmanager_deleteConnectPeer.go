package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_deleteConnectPeerCmd = &cobra.Command{
	Use:   "delete-connect-peer",
	Short: "Deletes a Connect peer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_deleteConnectPeerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_deleteConnectPeerCmd).Standalone()

		networkmanager_deleteConnectPeerCmd.Flags().String("connect-peer-id", "", "The ID of the deleted Connect peer.")
		networkmanager_deleteConnectPeerCmd.MarkFlagRequired("connect-peer-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_deleteConnectPeerCmd)
}
