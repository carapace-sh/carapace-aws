package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getManagedThingStateCmd = &cobra.Command{
	Use:   "get-managed-thing-state",
	Short: "Returns the managed thing state for the given device Id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getManagedThingStateCmd).Standalone()

	iotManagedIntegrations_getManagedThingStateCmd.Flags().String("managed-thing-id", "", "The id of the device.")
	iotManagedIntegrations_getManagedThingStateCmd.MarkFlagRequired("managed-thing-id")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getManagedThingStateCmd)
}
