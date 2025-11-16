package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_deleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "Deletes the specified connection in your global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_deleteConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_deleteConnectionCmd).Standalone()

		networkmanager_deleteConnectionCmd.Flags().String("connection-id", "", "The ID of the connection.")
		networkmanager_deleteConnectionCmd.Flags().String("global-network-id", "", "The ID of the global network.")
		networkmanager_deleteConnectionCmd.MarkFlagRequired("connection-id")
		networkmanager_deleteConnectionCmd.MarkFlagRequired("global-network-id")
	})
	networkmanagerCmd.AddCommand(networkmanager_deleteConnectionCmd)
}
