package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getManagedThingCertificateCmd = &cobra.Command{
	Use:   "get-managed-thing-certificate",
	Short: "Retrieves the certificate PEM for a managed IoT thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getManagedThingCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_getManagedThingCertificateCmd).Standalone()

		iotManagedIntegrations_getManagedThingCertificateCmd.Flags().String("identifier", "", "The identifier of the managed thing.")
		iotManagedIntegrations_getManagedThingCertificateCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getManagedThingCertificateCmd)
}
