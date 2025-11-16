package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getManagedThingMetaDataCmd = &cobra.Command{
	Use:   "get-managed-thing-meta-data",
	Short: "Get the metadata information for a managed thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getManagedThingMetaDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_getManagedThingMetaDataCmd).Standalone()

		iotManagedIntegrations_getManagedThingMetaDataCmd.Flags().String("identifier", "", "The managed thing id.")
		iotManagedIntegrations_getManagedThingMetaDataCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getManagedThingMetaDataCmd)
}
