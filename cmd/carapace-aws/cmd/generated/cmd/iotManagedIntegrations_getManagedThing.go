package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_getManagedThingCmd = &cobra.Command{
	Use:   "get-managed-thing",
	Short: "Get details of a managed thing including its attributes and capabilities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_getManagedThingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_getManagedThingCmd).Standalone()

		iotManagedIntegrations_getManagedThingCmd.Flags().String("identifier", "", "The id of the managed thing.")
		iotManagedIntegrations_getManagedThingCmd.MarkFlagRequired("identifier")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_getManagedThingCmd)
}
