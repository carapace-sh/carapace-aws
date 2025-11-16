package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getProvisioningProfileCmd = &cobra.Command{
	Use:   "get-provisioning-profile",
	Short: "Get a provisioning profile by template name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getProvisioningProfileCmd).Standalone()

	iotManagedIntegrations_getProvisioningProfileCmd.Flags().String("identifier", "", "The provisioning template the device uses for the provisioning process.")
	iotManagedIntegrations_getProvisioningProfileCmd.MarkFlagRequired("identifier")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getProvisioningProfileCmd)
}
