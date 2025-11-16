package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listProvisioningProfilesCmd = &cobra.Command{
	Use:   "list-provisioning-profiles",
	Short: "List the provisioning profiles within the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listProvisioningProfilesCmd).Standalone()

	iotManagedIntegrations_listProvisioningProfilesCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iotManagedIntegrations_listProvisioningProfilesCmd.Flags().String("next-token", "", "A token that can be used to retrieve the next set of results.")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listProvisioningProfilesCmd)
}
