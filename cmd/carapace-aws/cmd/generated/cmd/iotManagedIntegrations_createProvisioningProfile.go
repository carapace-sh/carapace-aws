package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_createProvisioningProfileCmd = &cobra.Command{
	Use:   "create-provisioning-profile",
	Short: "Create a provisioning profile for a device to execute the provisioning flows using a provisioning template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_createProvisioningProfileCmd).Standalone()

	iotManagedIntegrations_createProvisioningProfileCmd.Flags().String("ca-certificate", "", "The id of the certificate authority (CA) certificate.")
	iotManagedIntegrations_createProvisioningProfileCmd.Flags().String("client-token", "", "An idempotency token.")
	iotManagedIntegrations_createProvisioningProfileCmd.Flags().String("name", "", "The name of the provisioning template.")
	iotManagedIntegrations_createProvisioningProfileCmd.Flags().String("provisioning-type", "", "The type of provisioning workflow the device uses for onboarding to IoT managed integrations.")
	iotManagedIntegrations_createProvisioningProfileCmd.Flags().String("tags", "", "A set of key/value pairs that are used to manage the provisioning profile.")
	iotManagedIntegrations_createProvisioningProfileCmd.MarkFlagRequired("provisioning-type")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_createProvisioningProfileCmd)
}
