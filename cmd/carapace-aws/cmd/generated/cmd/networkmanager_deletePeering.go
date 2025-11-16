package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_deletePeeringCmd = &cobra.Command{
	Use:   "delete-peering",
	Short: "Deletes an existing peering connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_deletePeeringCmd).Standalone()

	networkmanager_deletePeeringCmd.Flags().String("peering-id", "", "The ID of the peering connection to delete.")
	networkmanager_deletePeeringCmd.MarkFlagRequired("peering-id")
	networkmanagerCmd.AddCommand(networkmanager_deletePeeringCmd)
}
