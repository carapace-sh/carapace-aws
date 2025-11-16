package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_listOrganizationServiceAccessStatusCmd = &cobra.Command{
	Use:   "list-organization-service-access-status",
	Short: "Gets the status of the Service Linked Role (SLR) deployment for the accounts in a given Amazon Web Services Organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_listOrganizationServiceAccessStatusCmd).Standalone()

	networkmanager_listOrganizationServiceAccessStatusCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	networkmanager_listOrganizationServiceAccessStatusCmd.Flags().String("next-token", "", "The token for the next page of results.")
	networkmanagerCmd.AddCommand(networkmanager_listOrganizationServiceAccessStatusCmd)
}
