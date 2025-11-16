package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_deleteProvisioningProfileCmd = &cobra.Command{
	Use:   "delete-provisioning-profile",
	Short: "Delete a provisioning profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_deleteProvisioningProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_deleteProvisioningProfileCmd).Standalone()

		iotManagedIntegrations_deleteProvisioningProfileCmd.Flags().String("identifier", "", "The name of the provisioning template.")
		iotManagedIntegrations_deleteProvisioningProfileCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_deleteProvisioningProfileCmd)
}
