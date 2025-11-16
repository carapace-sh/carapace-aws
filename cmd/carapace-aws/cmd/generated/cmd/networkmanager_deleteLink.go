package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_deleteLinkCmd = &cobra.Command{
	Use:   "delete-link",
	Short: "Deletes an existing link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_deleteLinkCmd).Standalone()

	networkmanager_deleteLinkCmd.Flags().String("global-network-id", "", "The ID of the global network.")
	networkmanager_deleteLinkCmd.Flags().String("link-id", "", "The ID of the link.")
	networkmanager_deleteLinkCmd.MarkFlagRequired("global-network-id")
	networkmanager_deleteLinkCmd.MarkFlagRequired("link-id")
	networkmanagerCmd.AddCommand(networkmanager_deleteLinkCmd)
}
