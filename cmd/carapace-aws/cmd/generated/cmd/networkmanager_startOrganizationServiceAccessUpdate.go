package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_startOrganizationServiceAccessUpdateCmd = &cobra.Command{
	Use:   "start-organization-service-access-update",
	Short: "Enables the Network Manager service for an Amazon Web Services Organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_startOrganizationServiceAccessUpdateCmd).Standalone()

	networkmanager_startOrganizationServiceAccessUpdateCmd.Flags().String("action", "", "The action to take for the update request.")
	networkmanager_startOrganizationServiceAccessUpdateCmd.MarkFlagRequired("action")
	networkmanagerCmd.AddCommand(networkmanager_startOrganizationServiceAccessUpdateCmd)
}
